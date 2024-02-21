package models

import "time"

type Account struct {
	BalanceAmount int64
	CreditLimit   int64
}

type Transaction struct {
	Amount          int64
	TransactionType string
	Description     string
	CreatedAt       time.Time
}

type Statement struct {
	Account      Account
	Transactions []Transaction
}
