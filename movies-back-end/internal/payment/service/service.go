package service

import (
	"context"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"log"
	"movies-service/config"
	"movies-service/internal/collection"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/movie"
	"movies-service/internal/payment"
	"movies-service/pkg/util"
	"time"
)

type paymentService struct {
	config               *config.Config
	managementCtrl       control.Service
	movieRepository      movie.Repository
	paymentRepository    payment.Repository
	collectionRepository collection.Repository
}

func NewPaymentService(config *config.Config, managementCtrl control.Service, movieRepository movie.Repository, paymentRepository payment.Repository, collectionRepository collection.Repository) payment.Service {
	return &paymentService{
		config:               config,
		managementCtrl:       managementCtrl,
		movieRepository:      movieRepository,
		paymentRepository:    paymentRepository,
		collectionRepository: collectionRepository,
	}
}

func (ps paymentService) VerifyPayment(ctx context.Context, provider model.PaymentProvider, providerPaymentID string, username string, typeCode string, refID uint) error {
	// Check user
	log.Printf("checking user...")
	isValidUser, _ := ps.managementCtrl.CheckUser(username)
	if !isValidUser {
		return errors.ErrInvalidClient
	}

	// Check movie
	log.Printf("checking movie...")
	theMovie, err := ps.movieRepository.FindMovieByID(ctx, refID)
	if err != nil {
		return err
	}

	if theMovie.ID == 0 || !theMovie.Price.Valid {
		return errors.ErrInvalidInput
	}

	// Check existed payment
	log.Printf("checking existed payment...")
	thePayment, err := ps.paymentRepository.FindByProviderPaymentID(ctx, provider, providerPaymentID)
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
	insertedPayment, err := ps.paymentRepository.InsertPayment(ctx, &model.Payment{
		RefID:             refID,
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
	theCollection := &model.Collection{
		Username: username,

		TypeCode:  typeCode,
		PaymentID: insertedPayment.ID,
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

	return nil
}
