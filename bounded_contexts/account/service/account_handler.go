package service

import (
	"com/txfer/bounded_contexts/account/dtos"
	"com/txfer/bounded_contexts/account/entities"
	"com/txfer/bounded_contexts/account/repositories"
	"com/txfer/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountHandler struct {
	AccountRepo *repositories.AccountRepo
}

func NewAccountHandler(DB *gorm.DB) AccountHandler {
	return AccountHandler{repositories.NewAccountRepo(DB)}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	// Parse input
	var accountReq dtos.CreateAccountRequest
	if err := helpers.BindData(c, &accountReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create account
	newAccount, err := entities.NewAccount(accountReq.InitialBalance, accountReq.AccountId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save
	if err := h.AccountRepo.CreateAccount(newAccount); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	// Parse input
	accountIdStr := c.Param("account_id")
	accountId, err := entities.ParseAccountId(accountIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get account by id
	acc, err := h.AccountRepo.GetAccountById(int64(accountId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return
	c.JSON(http.StatusOK, dtos.AccountResp(acc))
}
