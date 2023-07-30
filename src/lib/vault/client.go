package vault

import "github.com/hashicorp/vault/api"

type Client struct {
	config api.Config
}

type ClientOption = func(*Client)

func NewVaultClient(address string, opts ...ClientOption) *Client {
	config := &api.Config{
		Address: address,
	}

	client := &Client{
		config: *config,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}
	return client
}

func Address(address string) ClientOption {
	return func(c *Client) {
		c.config.Address = address
	}
}
