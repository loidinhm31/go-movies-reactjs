package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/model"
	"movies-service/internal/payment"
	"net/http"
	"strconv"
)

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(paymentService payment.Service) payment.Handler {
	return &paymentHandler{
		paymentService: paymentService,
	}
}

func (r paymentHandler) VerifyStripePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentID := c.Param("id")
		username := c.Query("username")
		typeCode := c.Query("type")

		refIdStr := c.Query("refId")
		refID, _ := strconv.Atoi(refIdStr)

		err := r.paymentService.VerifyPayment(c, model.STRIPE, paymentID, username, typeCode, uint(refID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
