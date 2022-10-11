package jwr

import "context"

const (
	GetUserProfilePath = "/user"
)

type JWR interface {
	GetUserProfile(ctx context.Context, userID int, apiTimeOut int) (*UserProfile, error)
}

func New(config JWRSDKConfig) (JWR, error) {

	return &JWRImpl{
		BaseURL:           config.BaseURL,
		Token:             config.Token,
		DefaultAPITimeout: config.APITimeout,
	}, nil
}
