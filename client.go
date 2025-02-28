package goauthorizenet

import (
	"github.com/mat285/go-authorize-net/config"
	"github.com/mat285/go-authorize-net/transactions"
)

type Client struct {
	Config             config.Config
	transactionsClient transactions.Interface
}

func New(cfg config.Config) *Client {
	return &Client{
		Config:             cfg,
		transactionsClient: transactions.New(cfg),
	}
}

func (c *Client) Transactions() transactions.Interface {
	return c.transactionsClient
}
