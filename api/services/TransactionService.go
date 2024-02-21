package services

import (
	"context"
	"rinhadev/api/models"
	"rinhadev/api/presenters"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionsService struct {
	db *pgxpool.Pool
}

func NewTransactionsService(db *pgxpool.Pool) TransactionsService {
	return TransactionsService{
		db: db,
	}
}

func (t TransactionsService) CreateTransaction(accountId int64, request presenters.TransactionCreateRequest) (*models.Account, error) {
	tx, err := t.db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), insert_transaction_sql, accountId, request.Amount, request.Type, request.Description)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), get_account_balance_sql, accountId)

	var account models.Account
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&account.CreditLimit, &account.BalanceAmount)
		if err != nil {
			return nil, err
		}
	}
	tx.Commit(context.Background())

	return &account, nil
}

const insert_transaction_sql = `
with upd as (
	update accounts 
	set balance = case when $3 = 'c' then balance + $2 else balance - $2 end
	where account_id = $1
)
insert into transactions (account_id, amount, "type", description)
values ($1, $2, $3::transaction_type, $4);
`

const get_account_balance_sql = `
select alimit, balance 
from accounts
where account_id = $1
`
