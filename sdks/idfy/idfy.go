package idfy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
)

type IdfyImpl struct {
	config     IdfyConfig
	nr         newrelic.Agent
	httpClient httpclient.HTTPClient
}

// New creates a new Idfy client
func New(config IdfyConfig, nr newrelic.Agent, client httpclient.HTTPClient) *IdfyImpl {
	idfy := IdfyImpl{
		config:     config,
		nr:         nr,
		httpClient: client,
	}
	return &idfy
}

func (idfyImpl *IdfyImpl) extract(documentType string, idfyrequest IdfyRequest) ([]byte, error) {
	url := idfyImpl.config.GetIdfyEndpoint() + documentType
	reqObj, _ := json.Marshal(idfyrequest)
	payload := strings.NewReader(string(reqObj))
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}
	idfyImpl.addHeaders(req)

	res, err := idfyImpl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (idfyImpl *IdfyImpl) ExtractPan(documentType string, idfyrequest IdfyRequest) (*IdfyPanResponse, error) {
	byteResp, err := idfyImpl.extract(documentType, idfyrequest)
	if err != nil {
		return nil, err
	}
	var idfyPanResp IdfyPanResponse
	err = json.Unmarshal(byteResp, &idfyPanResp)
	return &idfyPanResp, err
}

func (idfyImpl *IdfyImpl) ExtractAadhar(documentType string, idfyrequest IdfyRequest) (*IdfyAadharResponse, error) {
	byteResp, err := idfyImpl.extract(documentType, idfyrequest)
	if err != nil {
		return nil, err
	}
	var idfyAadharResponse IdfyAadharResponse
	err = json.Unmarshal(byteResp, &idfyAadharResponse)
	return &idfyAadharResponse, err
}

func (idfyImpl *IdfyImpl) ExtractDl(documentType string, idfyrequest IdfyRequest) (*IdfyDlResponse, error) {
	byteResp, err := idfyImpl.extract(documentType, idfyrequest)
	if err != nil {
		return nil, err
	}
	var idfyDlResponse IdfyDlResponse
	err = json.Unmarshal(byteResp, &idfyDlResponse)
	return &idfyDlResponse, err
}

func (idfyImpl *IdfyImpl) ExtractVoter(documentType string, idfyrequest IdfyRequest) (*IdfyVoterIdResponse, error) {
	byteResp, err := idfyImpl.extract(documentType, idfyrequest)
	if err != nil {
		return nil, err
	}
	var idfyVoterIdResponse IdfyVoterIdResponse
	err = json.Unmarshal(byteResp, &idfyVoterIdResponse)
	return &idfyVoterIdResponse, err
}

func (idfyImpl *IdfyImpl) ExtractPassport(documentType string, idfyrequest IdfyRequest) (*IdfyPassportResponse, error) {
	byteResp, err := idfyImpl.extract(documentType, idfyrequest)
	if err != nil {
		return nil, err
	}
	var idfyPassportResponse IdfyPassportResponse
	err = json.Unmarshal(byteResp, &idfyPassportResponse)
	return &idfyPassportResponse, err
}

func (idfyImpl *IdfyImpl) addHeaders(req *http.Request) {
	req.Header.Add("account-id", idfyImpl.config.GetIdfyAccountId())
	req.Header.Add("api-key", idfyImpl.config.GetIdfyApiKey())
	req.Header.Add("Content-Type", "application/json")
}
