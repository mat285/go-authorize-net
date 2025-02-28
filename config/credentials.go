package config

type MerchantCredentials struct {
	ID             string `json:"id" yaml:"id"`
	TransactionKey string `json:"transactionKey" yaml:"transactionKey"`
}
