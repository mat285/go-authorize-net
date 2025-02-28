package authentication

type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

type Authenticatable interface {
	SetMerchantAuth(MerchantAuthentication)
}
