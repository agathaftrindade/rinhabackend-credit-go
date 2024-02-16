package models

type Account struct {
	BalanceAmount int64
	CreditLimit   int64
}

type Transaction struct {
	Amount          int64
	TransactionType rune
	Description     string
	CreatedAt       string
}

type Statement struct {
	Account      Account
	Transactions []Transaction
}
