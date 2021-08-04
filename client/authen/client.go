package authen

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type AuthenClient struct {
	cfg    *Configuration
	Client *http.Client
}

func NewAuthenClient(cfg *Configuration) (*AuthenClient, error) {
	c := &AuthenClient{}
	c.cfg = cfg
	config := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       []string{"email"},
		TokenURL:     cfg.TokenURL,
	}

	_, err := config.TokenSource(context.Background()).Token()
	if err != nil {
		return nil, err
	}
	c.Client = config.Client(context.TODO())
	return c, nil
}
