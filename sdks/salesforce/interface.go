package salesforce

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
)

const (
	RequestAccessTokenPath = "/services/oauth2/token"
	CreateTaskPath         = "/services/data/v56.0/sobjects/Task/"
)

type Salesforce interface {
	RequestAccessToken(ctx context.Context, accessTokenRequest AccessTokenRequest, apiTimeout int) (*AccessTokenResponse, error)
	CreateTask(ctx context.Context, createTaskRequest CreateTaskRequest, apiTimeout int) (*CreateTaskResponse, error)
}

func New(config SalesforceSDKConfig, client httpclient.HTTPClient) Salesforce {
	return &SalesforceImpl{
		BaseURL:           config.BaseURL,
		Token:             config.Token,
		DefaultAPITimeout: config.APITimeout,
		httpClient:        client,
	}
}
