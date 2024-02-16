package api

import "rinhadev/api/models"

type AccountsService struct {
}

func (AccountsService) GetStatement(accountId int64) (models.Statement, error) {
}
