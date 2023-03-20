package vault

import (
	"errors"

	"github.com/hashicorp/vault/api"
)

type Client struct {
	*api.Client
}

func NewClient(config *api.Config) (*Client, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

func (c *Client) Authenticate(appRoleID, appSecretID string) (string, error) {
	path := "auth/approle/login"

	secret, err := c.Logical().Write(path, map[string]interface{}{
		"role_id":   appRoleID,
		"secret_id": appSecretID,
	})
	if err != nil {
		return "", err
	}
	if secret == nil || secret.Auth == nil {
		return "", errors.New("authentication failed")
	}
	return secret.Auth.ClientToken, nil
}
