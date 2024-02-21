package presenters

import (
	"rinhadev/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func NewStatementShowResponse(statement models.Statement) gin.H {

	transactions := make([]gin.H, len(statement.Transactions))
	for i, t := range statement.Transactions {
		transactions[i] = gin.H{
			"valor":        t.Amount,
			"tipo":         t.TransactionType,
			"descricao":    t.Description,
			"realizada_em": t.CreatedAt.String(),
		}
	}

	return gin.H{
		"saldo": map[string]any{
			"total":        statement.Account.BalanceAmount,
			"data_extrato": time.Now().String(),
			"limite":       statement.Account.CreditLimit,
		},
		"ultimas_transacoes": transactions,
	}
}
