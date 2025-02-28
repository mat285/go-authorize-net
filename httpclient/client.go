package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mat285/go-authorize-net/authentication"
	"github.com/mat285/go-authorize-net/config"
)

const (
	APIBasePath = "/xml/v1/request.api"

	SandboxAPIAddress    = "https://apitest.authorize.net" + APIBasePath
	ProductionAPIAddress = "https://api.authorize.net" + APIBasePath
)

type Client struct {
	Config config.Config
	Client *http.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		Config: cfg,
		Client: http.DefaultClient,
	}
}

func (c *Client) PostJSON(ctx context.Context, request authentication.Authenticatable, response interface{}) error {
	request.SetMerchantAuth(c.Config.MerchantAuth())
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiURL(), bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status returned %d", resp.StatusCode)
	}
	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// There's strange characters at the start of their json response for some reason, just trim til we find the start of json
	output = bytes.TrimLeftFunc(output, func(r rune) bool { return r > rune(1<<7) })

	return json.Unmarshal(output, response)
}

func (c *Client) apiURL() string {
	switch c.Config.Env {
	case config.EnvironmentProd:
		return ProductionAPIAddress
	default:
		return SandboxAPIAddress
	}
}
