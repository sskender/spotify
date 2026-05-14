package spotify

import (
	"context"
	"net/http"

	"github.com/sskender/spotify/pkg/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

type persistentTokenSource struct {
	store  auth.TokenStore
	source oauth2.TokenSource
}

func (s *persistentTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.source.Token()
	if err != nil {
		return nil, err
	}

	refreshToken := token.RefreshToken
	if len(refreshToken) > 0 && refreshToken != s.store.GetRefreshToken() {
		if err = s.store.RotateRefreshToken(refreshToken); err != nil {
			return nil, err
		}
	}

	return token, nil
}

func NewHTTPClient(ctx context.Context, store auth.TokenStore) *http.Client {
	config := &oauth2.Config{
		ClientID:     store.GetClientID(),
		ClientSecret: store.GetClientSecret(),
		Endpoint:     spotify.Endpoint,
	}

	initialToken := &oauth2.Token{
		RefreshToken: store.GetRefreshToken(),
		TokenType:    "Bearer",
	}

	tokenSrc := config.TokenSource(ctx, initialToken)

	persistentTokenSrc := &persistentTokenSource{
		source: tokenSrc,
		store:  store,
	}

	client := oauth2.NewClient(ctx, persistentTokenSrc)

	return client
}
