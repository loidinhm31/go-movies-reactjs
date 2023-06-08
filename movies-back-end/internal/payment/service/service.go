package service

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"log"
	"movies-service/config"
	"movies-service/internal/collection"
	"movies-service/internal/common/constant"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/common/mapper"
	"movies-service/internal/common/model"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/mail"
	"movies-service/internal/middlewares"
	"movies-service/internal/movie"
	"movies-service/internal/payment"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"strings"
	"time"
)

type paymentService struct {
	config               *config.Config
	userRepository       user.UserRepository
	movieRepository      movie.Repository
	episodeRepository    episode.Repository
	paymentRepository    payment.Repository
	collectionRepository collection.Repository
	mailService          mail.Service
}

func NewPaymentService(config *config.Config, userRepository user.UserRepository,
	movieRepository movie.Repository, episodeRepository episode.Repository,
	paymentRepository payment.Repository, collectionRepository collection.Repository,
	mailService mail.Service) payment.Service {
	return &paymentService{
		config:               config,
		userRepository:       userRepository,
		movieRepository:      movieRepository,
		episodeRepository:    episodeRepository,
		paymentRepository:    paymentRepository,
		collectionRepository: collectionRepository,
		mailService:          mailService,
	}
}

func (ps *paymentService) VerifyPayment(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string, username string, typeCode string, refID uint) error {
	if typeCode == "" {
		return errors.ErrInvalidInput
	}

	// Check user
	log.Printf("checking user...")
	theUser, err := ps.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	if theUser.Role.RoleCode == "BANNED" {
		return errors.ErrInvalidClient
	}

	var theMovie *entity.Movie
	var theEpisode *entity.Episode
	if typeCode == "MOVIE" {
		// Check movie
		log.Printf("checking movie...")
		theMovie, err = ps.movieRepository.FindMovieByID(ctx, refID)
		if err != nil {
			return err
		}

		if theMovie.ID == 0 || !theMovie.Price.Valid {
			return errors.ErrInvalidInput
		}
	} else if typeCode == "TV" {
		// Check episode
		log.Printf("checking episode...")
		theEpisode, err = ps.episodeRepository.FindEpisodeByID(ctx, refID)
		if err != nil {
			return err
		}

		if theMovie.ID == 0 || !theMovie.Price.Valid {
			return errors.ErrInvalidInput
		}
	}

	// Check existed payment
	log.Printf("checking existed payment...")
	thePayment, err := ps.paymentRepository.FindPaymentByProviderPaymentID(ctx, provider, providerPaymentID)
	if err != nil {
		return err
	}

	if thePayment.ProviderPaymentID.Valid {
		return errors.ErrObjectExisted
	}

	// Set Stripe secret key
	log.Printf("verifying from provider...")
	stripe.Key = ps.config.Stripe.SecretKey

	params := &stripe.PaymentIntentParams{}
	params.AddExpand("payment_method")

	// Retrieve the PaymentIntent
	paymentIntent, err := paymentintent.Get(providerPaymentID, params)
	if err != nil {
		log.Printf("Error retrieving PaymentIntent: %v", err)
		return err
	}

	if paymentIntent.Status != "succeeded" {
		return errors.ErrPaymentNotFound
	}

	// Add payment
	log.Printf("verified from provider...")
	_, err = ps.paymentRepository.InsertPayment(ctx, &entity.Payment{
		RefID:             refID,
		UserID:            theUser.ID,
		TypeCode:          typeCode,
		Provider:          string(provider),
		ProviderPaymentID: util.StringToSQLNullString(providerPaymentID),
		Amount:            float64(paymentIntent.Amount),
		Received:          float64(paymentIntent.Amount - paymentIntent.ApplicationFeeAmount),
		Currency:          string(paymentIntent.Currency),
		PaymentMethod:     string(paymentIntent.PaymentMethod.Type),
		Status:            string(paymentIntent.Status),
		CreatedAt:         time.Now(),
		CreatedBy:         "system",
	})
	if err != nil {
		return err
	}

	// Add collection
	theCollection := &entity.Collection{
		UserID:    theUser.ID,
		TypeCode:  typeCode,
		CreatedAt: time.Now(),
		CreatedBy: username,
	}

	if typeCode == "MOVIE" {
		theCollection.MovieID = util.IntToSQLNullInt(int64(refID))
	} else if typeCode == "TV" {
		theCollection.EpisodeID = util.IntToSQLNullInt(int64(refID))
	}

	err = ps.collectionRepository.InsertCollection(ctx, theCollection)
	if err != nil {
		return err
	}

	// Send email to user
	go func() {
		var htmlMessage string
		if typeCode == "MOVIE" {
			htmlMessage = fmt.Sprintf(`
		<strong>Payment Confirmation</strong><br>
		Dear %s %s, <br>
		This is confirmation for your payment %f %s on %s at %s.`,
				theUser.FirstName, theUser.LastName,
				thePayment.Received, strings.ToUpper(thePayment.Currency),
				theMovie.Title,
				thePayment.CreatedAt.Format(constant.Layout))
		} else if typeCode == "TV" {
			htmlMessage = fmt.Sprintf(`
		<strong>Payment Confirmation</strong><br>
		Dear %s %s, <br>
		This is confirmation for your payment %f %s on %s - %s at %s.`,
				theUser.FirstName, theUser.LastName,
				thePayment.Received, strings.ToUpper(thePayment.Currency),
				theEpisode.Season.Name, theEpisode.Name,
				thePayment.CreatedAt.Format(constant.Layout))
		}

		err := ps.mailService.SendMessage(ctx, &model.MailData{
			To:           theUser.Email,
			From:         ps.config.Mail.From,
			Subject:      constant.PaymentConfirmationSubject,
			Content:      htmlMessage,
			TemplateMail: "basic.html",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (ps *paymentService) GetPaymentsByUserAndTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*dto.PaymentDto, error) {
	if typeCode == "" {
		return nil, errors.ErrInvalidInput
	}

	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, err := ps.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if theUser.Role.RoleCode == "BANNED" {
		return nil, errors.ErrInvalidClient
	}

	result, err := ps.paymentRepository.FindPaymentByUserIDAndTypeCodeAndRefID(ctx, theUser.ID, typeCode, refID)
	if err != nil {
		return nil, err
	}
	return &dto.PaymentDto{
		ID:       result.ID,
		UserID:   result.UserID,
		RefID:    result.RefID,
		TypeCode: result.TypeCode,
		Status:   result.Status,
	}, nil
}

func (ps *paymentService) GetPaymentsByUser(ctx context.Context, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CustomPaymentDto], error) {
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, err := ps.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	page := &pagination.Page[*entity.CustomPayment]{}

	if theUser.Role.RoleCode == "BANNED" {
		return nil, errors.ErrInvalidClient
	}

	results, err := ps.paymentRepository.FindPaymentsByUserID(ctx, theUser.ID, keyword, pageRequest, page)
	if err != nil {
		return nil, err
	}

	customPaymentDtos := mapper.MapToCustomPaymentDtoSlice(results.Content)

	return &pagination.Page[*dto.CustomPaymentDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: results.TotalElements,
		TotalPages:    results.TotalPages,
		Content:       customPaymentDtos,
	}, nil
}
