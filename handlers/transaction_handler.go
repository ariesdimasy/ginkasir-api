package handlers

import (
	"ginkasir/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (th *TransactionHandler) HandleCheckout(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "checkout success",
	})

}
