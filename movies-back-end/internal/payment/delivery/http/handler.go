package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/common/entity"
	"movies-service/internal/payment"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
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

		err := r.paymentService.VerifyPayment(c, entity.STRIPE, paymentID, username, typeCode, uint(refID))
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

func (r paymentHandler) FetchPaymentsByUserAndRefID() gin.HandlerFunc {
	return func(c *gin.Context) {
		typeCode := c.Query("type")
		refIdStr := c.Param("id")
		refID, _ := strconv.Atoi(refIdStr)

		result, err := r.paymentService.GetPaymentsByUserAndTypeCodeAndRefID(c, typeCode, uint(refID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (r paymentHandler) FetchPaymentsByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("q")

		pageable, _ := pagination.ReadPageRequest(c)

		if err := util.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		results, err := r.paymentService.GetPaymentsByUser(c, keyword, pageable)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, results)
	}
}
