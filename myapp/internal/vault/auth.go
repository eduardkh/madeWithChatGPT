package vault

import (
	"fmt"
)

func (c *Client) AuthenticateWithUserPass(username, password string) (string, error) {
	// Replace "userpass" with the correct authentication method's path if different
	path := fmt.Sprintf("secret/data/myapp/%s", username)

	secret, err := c.Client.Logical().Write(path, map[string]interface{}{
		"password": password,
	})
	if err != nil {
		return "", err
	}
	if secret == nil || secret.Auth == nil {
		return "", fmt.Errorf("authentication failed")
	}
	return secret.Auth.ClientToken, nil
}

func (c *Client) Authorize(token, path string) error {
	client, err := c.Client.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone client: %v", err)
	}
	client.SetToken(token)

	secret, err := client.Logical().Read(path)
	if err != nil {
		return err
	}
	if secret == nil {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
