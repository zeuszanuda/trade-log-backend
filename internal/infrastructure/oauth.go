package infrastructure

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	TokenURL     string
	Scopes       []string
}

func NewOAuth2Config(cfg OAuthConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       cfg.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.AuthURL,
			TokenURL: cfg.TokenURL,
		},
	}
}

func ExchangeCode(ctx context.Context, conf *oauth2.Config, code string) (*oauth2.Token, error) {
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("oauth exchange error: %w", err)
	}
	return token, nil
}
