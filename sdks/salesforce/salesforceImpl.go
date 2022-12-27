package salesforce

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type SalesforceImpl struct {
	BaseURL           string
	Token             string
	DefaultAPITimeout int
	httpClient        httpclient.HTTPClient
}

func (salesforceImpl *SalesforceImpl) RequestAccessToken(ctx context.Context, accessTokenRequest AccessTokenRequest, apiTimeout int) (*AccessTokenResponse, error) {
	validate := validator.New()
	if err := validate.Struct(accessTokenRequest); err != nil {
		return nil, err
	}

	accessTokenRequestBytes, err := json.Marshal(accessTokenRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, salesforceImpl.BaseURL+RequestAccessTokenPath, bytes.NewReader(accessTokenRequestBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set(contentType, applicationJson)
	request.Header.Set(authorization, salesforceImpl.Token)

	resp, err := salesforceImpl.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrapf(err, "while making api call to Request Access Token")
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading response from Request Access Token")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + body.String())
	}

	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body.Bytes(), &accessTokenResponse)
	if err != nil {
		return nil, err
	}
	return &accessTokenResponse, nil
}

func (salesforceImpl *SalesforceImpl) CreateTask(ctx context.Context, createTaskRequest CreateTaskRequest, apiTimeout int) (*CreateTaskResponse, error) {
	validate := validator.New()
	if err := validate.Struct(createTaskRequest); err != nil {
		return nil, err
	}

	createTaskRequestBytes, err := json.Marshal(createTaskRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, salesforceImpl.BaseURL+CreateTaskPath, bytes.NewReader(createTaskRequestBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set(contentType, applicationJson)
	request.Header.Set("access_token", salesforceImpl.Token)

	resp, err := salesforceImpl.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrapf(err, "while making api call to Create Task")
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading response from Create Task")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + body.String())
	}

	var createTaskResponse CreateTaskResponse
	err = json.Unmarshal(body.Bytes(), &createTaskResponse)
	if err != nil {
		return nil, err
	}

	for _, errors := range createTaskResponse.Errors {
		if errors == "xyx" {
			return nil, fmt.Errorf(errors, ErrInAvailableService)
		}
	}

	return &createTaskResponse, nil
}
