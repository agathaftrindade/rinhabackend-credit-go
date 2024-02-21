package presenters

type TransactionCreateRequest struct {
	Amount      int64  `json:"valor" binding:"required"`
	Type        string `json:"tipo" binding:"required"`
	Description string `json:"descricao" binding:"required"`
}
