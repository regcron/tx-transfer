package service

import (
	"com/txfer/bounded_contexts/transfer/dtos"
	"com/txfer/bounded_contexts/transfer/repositories"
	"com/txfer/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransferHandler struct {
	TransferRepo *repositories.TransferRepo
}

func NewTransferHandler(DB *gorm.DB) TransferHandler {
	return TransferHandler{repositories.NewTransferRepo(DB)}
}

func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	// Parse input
	var transferReq dtos.TransferRequestDto
	if err := helpers.BindData(c, &transferReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create transfer
	txn, err := h.TransferRepo.CreateTransfer(transferReq.SrcAccId, transferReq.DestAccId, transferReq.Amount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return
	c.JSON(http.StatusOK, dtos.TransferResp(txn))
}
