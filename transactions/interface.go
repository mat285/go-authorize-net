package transactions

import "context"

type Interface interface {
	CreateTransaction(context.Context, CreateTransactionRequest) (*CreateTransactionResponse, error)
}
