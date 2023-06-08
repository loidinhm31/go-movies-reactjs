package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/payment"
)

// MapPaymentRoutes Map payment routes
func MapPaymentRoutes(paymentGroup *gin.RouterGroup, h payment.Handler) {
	paymentGroup.GET("/stripe/:id/verification", h.VerifyStripePayment())
	paymentGroup.GET("/refs/:id", h.FetchPaymentsByUserAndRefID())
}

func MapAuthPaymentRoutes(paymentGroup *gin.RouterGroup, h payment.Handler) {
	paymentGroup.POST("/", h.FetchPaymentsByUser())
	paymentGroup.GET("/", h.FetchPaymentsByUser())

}
