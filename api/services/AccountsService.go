package services

import (
	"context"
	"rinhadev/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountsService struct {
	db *pgxpool.Pool
}

func NewAccountsService(db *pgxpool.Pool) AccountsService {
	return AccountsService{
		db: db,
	}
}

func (s AccountsService) GetStatement(accountId int64) (*models.Statement, error) {
	rows, err := s.db.Query(context.Background(), query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	i := 0
	account := models.Account{}
	transactions := make([]models.Transaction, 0)
	for rows.Next() {
		if i == 0 {
			err = rows.Scan(
				nil,
				&account.BalanceAmount,
				nil,
				nil,
				nil,
				&account.CreditLimit,
			)
		} else {
			transaction := models.Transaction{}
			err = rows.Scan(
				nil,
				&transaction.Amount,
				&transaction.TransactionType,
				&transaction.Description,
				&transaction.CreatedAt,
				nil,
			)
			transactions = append(transactions, transaction)
		}
		i++

		if err != nil {
			return nil, err
		}
	}

	if i == 0 {
		return nil, nil
	}

	statement := models.Statement{
		Account:      account,
		Transactions: transactions,
	}

	return &statement, nil
}

const query = `
select 0 as i, balance, 'c' as transaction_type, '' as description, now() createdat, alimit
from accounts a 
where a.account_id = $1
union all
select 1 as i, amount, "type", description, createdat, 0 
from transactions t
where t.account_id = $1
order by i, createdat desc
limit 11
`
