package transactions

import (
	"context"

	"github.com/mat285/go-authorize-net/config"
	"github.com/mat285/go-authorize-net/httpclient"
)

var (
	_ Interface = new(Client)
)

type Client struct {
	*httpclient.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		Client: httpclient.New(cfg),
	}
}

func (c *Client) CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*CreateTransactionResponse, error) {
	req := &Request{CreateTransactionRequest: &request}
	var out CreateTransactionResponse
	err := c.Client.PostJSON(ctx, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
