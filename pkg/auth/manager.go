package auth

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type token struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type tokenManager struct {
	lock       sync.Mutex
	ctx        context.Context
	session    *secretsmanager.Client
	secretName string
	token      token
}

func NewTokenManager(ctx context.Context) (TokenStore, error) {
	secretName := os.Getenv("SPOTIFY_API_CREDENTIALS")
	if secretName == "" {
		return nil, errors.New("missing environment variable: SPOTIFY_API_CREDENTIALS")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		return nil, err
	}

	var token token
	if err = json.Unmarshal([]byte(*result.SecretString), &token); err != nil {
		return nil, err
	}

	return &tokenManager{
		ctx:        ctx,
		session:    svc,
		secretName: secretName,
		token:      token,
	}, nil
}

func (m *tokenManager) GetClientID() string {
	return m.token.ClientID
}

func (m *tokenManager) GetClientSecret() string {
	return m.token.ClientSecret
}

func (m *tokenManager) GetRefreshToken() string {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.token.RefreshToken
}

func (m *tokenManager) RotateRefreshToken(newToken string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.token.RefreshToken == newToken {
		return nil
	}

	m.token.RefreshToken = newToken

	secretString, err := json.Marshal(m.token)
	if err != nil {
		return err
	}

	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(m.secretName),
		SecretString: aws.String(string(secretString)),
	}

	_, err = m.session.PutSecretValue(m.ctx, input)
	if err != nil {
		return err
	}

	return nil
}
