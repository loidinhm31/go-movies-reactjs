package service

import (
	"context"
	"database/sql"
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

func (ps paymentService) VerifyPayment(ctx context.Context, provider model.PaymentProvider, providerPaymentID string, username string, movieID uint) error {
	// Check user
	log.Printf("checking user...")
	if !ps.managementCtrl.CheckUser(username) {
		return errors.ErrInvalidClient
	}

	// Check movie
	log.Printf("checking movie...")
	theMovie, err := ps.movieRepository.FindMovieById(ctx, movieID)
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

	log.Printf("verifying from provider...")
	// Set Stripe secret key
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

	log.Printf("verified from provider...")
	// Add payment
	insertedPayment, err := ps.paymentRepository.InsertPayment(ctx, &model.Payment{
		Provider: string(provider),
		ProviderPaymentID: sql.NullString{
			String: providerPaymentID,
			Valid:  true,
		},
		Amount:        float64(paymentIntent.Amount),
		Received:      float64(paymentIntent.Amount - paymentIntent.ApplicationFeeAmount),
		Currency:      string(paymentIntent.Currency),
		PaymentMethod: string(paymentIntent.PaymentMethod.Type),
		Status:        string(paymentIntent.Status),
		CreatedAt:     time.Now(),
		CreatedBy:     "system",
	})
	if err != nil {
		return err
	}

	// Add collection
	err = ps.collectionRepository.InsertCollection(ctx, &model.Collection{
		Username:  username,
		MovieID:   movieID,
		Payment:   insertedPayment,
		CreatedAt: time.Now(),
		CreatedBy: username,
	})
	if err != nil {
		return err
	}

	return nil
}
