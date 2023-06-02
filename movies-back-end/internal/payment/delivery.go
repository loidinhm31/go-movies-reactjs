package payment

import "github.com/gin-gonic/gin"

type Handler interface {
	VerifyStripePayment() gin.HandlerFunc
}
