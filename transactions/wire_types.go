package transactions

import "github.com/mat285/go-authorize-net/authentication"

type CreateTransactionRequest struct {
	MerchantAuthentication authentication.MerchantAuthentication `json:"merchantAuthentication"`

	RefID              string             `json:"refId"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

type CreateTransactionResponse struct {
	TransactionResponse TransactionResponse `json:"transactionResponse"`

	RefID    string                            `json:"refId"`
	Messages CreateTransactionResponseMessages `json:"messages"`
}

type Request struct {
	CreateTransactionRequest *CreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

func (r *Request) SetMerchantAuth(auth authentication.MerchantAuthentication) {
	if r == nil || r.CreateTransactionRequest == nil {
		return
	}
	r.CreateTransactionRequest.MerchantAuthentication = auth
}
