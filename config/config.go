package config

import "github.com/mat285/go-authorize-net/authentication"

type Config struct {
	Credentials MerchantCredentials `json:"credentials" yaml:"credentials"`

	Env string `json:"env" yaml:"env"`
}

func (c Config) MerchantAuth() authentication.MerchantAuthentication {
	return authentication.MerchantAuthentication{
		Name:           c.Credentials.ID,
		TransactionKey: c.Credentials.TransactionKey,
	}
}

const (
	EnvironmentSandbox = "sandbox"
	EnvironmentProd    = "prod"
)
