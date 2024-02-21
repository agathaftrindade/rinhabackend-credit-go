package presenters

import (
	"rinhadev/api/models"

	"github.com/gin-gonic/gin"
)

func NewTransactionCreateResponse(account models.Account) gin.H {
	return gin.H{
		"limite": account.CreditLimit,
		"saldo":  account.BalanceAmount,
	}
}
