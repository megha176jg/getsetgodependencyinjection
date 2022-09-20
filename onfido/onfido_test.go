package onfido

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	httpclientmocks "bitbucket.org/junglee_games/getsetgo/httpclient/mocks"
	newrelicmocks "bitbucket.org/junglee_games/getsetgo/newrelic/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestOnfidoSuite(t *testing.T) {
	suite.Run(t, new(onfidoSuite))
}

type onfidoSuite struct {
	suite.Suite
	httpClient httpclientmocks.HTTPClient
	nr         newrelicmocks.Agent
	srv        Onfido
}

func (suite *onfidoSuite) SetupTest() {
	suite.httpClient = *httpclientmocks.NewHTTPClient(suite.T())
	suite.nr = *newrelicmocks.NewAgent(suite.T())
	suite.nr.On("StartTransaction", mock.Anything).Return(nil)
	suite.srv = New("endpoint", "authToken", &suite.httpClient, &suite.nr)
}

func (suite *onfidoSuite) TestOnfidoSDK_CreateApplicant() {

	type args struct {
		firstName string
		lastName  string
		location  bool
	}
	tests := []struct {
		name    string
		args    args
		doErr   error
		want    *CreateApplicantResponse
		wantErr bool
	}{
		{
			name:    "when api gives error",
			doErr:   ErrCallingOnfido,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when api works fine",
			doErr:   nil,
			want:    &CreateApplicantResponse{ID: "testid"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)
			got, err := suite.srv.CreateApplicant(tt.args.firstName, tt.args.lastName, tt.args.location)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}

func (suite *onfidoSuite) TestOnfidoSDK_UploadDocument() {

	type args struct {
		applicantId string
		fileType    string
		filePath    string
		side        string
	}
	tests := []struct {
		name    string
		args    args
		doErr   error
		want    *UploadDocumentResponse
		wantErr bool
	}{
		{
			name:    "when api gives error",
			doErr:   ErrCallingOnfido,
			want:    nil,
			wantErr: true,
		},
		// {
		// 	name:    "when api works fine",
		// 	doErr:   nil,
		// 	want:    &UploadDocumentResponse{},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)

			got, err := suite.srv.UploadDocument(tt.args.applicantId, tt.args.fileType, tt.args.filePath, tt.args.side)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}

func (suite *onfidoSuite) TestOnfidoSDK_CreateCheck() {

	type args struct {
		applicantId string
		reportNames []string
	}

	tests := []struct {
		name    string
		args    args
		want    *CreateCheckResponse
		doErr   error
		wantErr bool
	}{
		{
			name:    "when api gives error",
			doErr:   ErrCallingOnfido,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when api works fine",
			doErr:   nil,
			want:    &CreateCheckResponse{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)

			got, err := suite.srv.CreateCheck(tt.args.applicantId, tt.args.reportNames)
			log.Print(err)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}

}

func (suite *onfidoSuite) TestOnfidoSDK_RetriveReport() {
	type fields struct {
		Endpoint  string
		AuthToken string
	}
	type args struct {
		reportId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ReportResponse
		doErr   error
		wantErr bool
	}{
		{
			name:    "when api gives error",
			doErr:   ErrCallingOnfido,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when api works fine",
			doErr:   nil,
			want:    &ReportResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte
			if tt.want != nil {
				str, _ = json.Marshal(tt.want)
			} else {
				str = nil
			}
			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)

			got, err := suite.srv.RetriveReport(tt.args.reportId)
			suite.Equal(tt.want, got)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}

func (suite *onfidoSuite) TestOnfidoSDK_DownloadDocument() {

	type args struct {
		documentId string
		destPath   string
	}
	tests := []struct {
		name    string
		args    args
		doErr   error
		wantErr bool
	}{
		{
			name:    "when api gives error",
			doErr:   ErrCallingOnfido,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var str []byte

			suite.httpClient.ExpectedCalls = []*mock.Call{}
			suite.httpClient.Calls = []mock.Call{}
			suite.httpClient.On("Do", mock.Anything).Return(&http.Response{Status: "Okay", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(string(str)))}, tt.doErr)

			err := suite.srv.DownloadDocument(tt.args.documentId, tt.args.destPath)
			suite.Equal(tt.wantErr, err != nil)
		})
	}
}
