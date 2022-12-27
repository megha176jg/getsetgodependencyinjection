package salesforce

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	httpclientmocks "bitbucket.org/junglee_games/getsetgo/httpclient/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestDigilockerSuite(t *testing.T) {
	suite.Run(t, new(salesforceSuite))
}

type salesforceSuite struct {
	suite.Suite
	httpClient httpclientmocks.HTTPClient
	srv        Salesforce
}

func (suite *salesforceSuite) SetupTest() {
	suite.httpClient = *httpclientmocks.NewHTTPClient(suite.T())
	suite.srv = New(SalesforceSDKConfig{}, &suite.httpClient)
}

func (suite *salesforceSuite) TestRequestAccessToken() {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *AccessTokenResponse
		doErr   error
		wantErr bool
		request AccessTokenRequest
	}{
		{
			name:    "when RequestAccessToken returns error",
			want:    nil,
			doErr:   errors.New("some error"),
			wantErr: true,
			request: AccessTokenRequest{},
		},
		{
			name: "when RequestAccessToken works fine",
			want: &AccessTokenResponse{
				AccessToken: "00D0T0000008aZc!AQgAQEvKdrdchjv22bKPlSkwsN8h3R1USKIBCd.wzp6HizknyCSUKYGSHWVHLovTw7b_ty0ots_w3SRvSiHcWBPzQQbdSkSi",
				InstanceURL: "https://jungleerummy--uat.sandbox.my.salesforce.com",
				ID:          "https://test.salesforce.com/id/00D0T0000008aZcUAI/0050T000000bXflQAE",
				TokenType:   "Bearer",
				IssuedAt:    "1668490886675",
				Signature:   "7je2GT4/pCLmMhI0jx3oakvIYp0KWPljOsUP/yDf9AU=",
			},
			request: AccessTokenRequest{
				GrantType:    "password",
				Assertion:    "tempAssert",
				ClientId:     "123",
				ClientSecret: "abc",
				Username:     "user",
				Password:     "123",
			},
			doErr:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(&tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)
			got, err := suite.srv.RequestAccessToken(tt.args.ctx, tt.request, 1)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}

func (suite *salesforceSuite) TestCreateTask() {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateTaskResponse
		doErr   error
		wantErr bool
		request CreateTaskRequest
	}{
		{
			name:    "when RequestAccessToken returns error",
			want:    nil,
			doErr:   errors.New("some error"),
			wantErr: true,
			request: CreateTaskRequest{},
		},
		{
			name: "when RequestAccessToken works fine",
			want: &CreateTaskResponse{
				ID:      "123",
				Success: true,
			},
			request: CreateTaskRequest{
				Status:        "Open",
				Subject:       "Dummy Subject",
				Priority:      "Priority",
				UserID:        "1963886",
				SocialNetwork: "JUNGLEERUMMY",
			},
			doErr:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(&tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)
			got, err := suite.srv.CreateTask(tt.args.ctx, tt.request, 1)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}
