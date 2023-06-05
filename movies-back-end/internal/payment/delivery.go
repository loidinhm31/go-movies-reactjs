package payment

import "github.com/gin-gonic/gin"

type Handler interface {
	VerifyStripePayment() gin.HandlerFunc
	FetchPaymentsByUser() gin.HandlerFunc
	FetchPaymentsByUserAndRefID() gin.HandlerFunc
}
