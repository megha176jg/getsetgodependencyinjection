package digilocker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/newrelic"

	"github.com/pkg/errors"
)

type DigilockerImpl struct {
	appId       string
	appKey      string
	hvEndpoint  string
	nr          newrelic.Agent
	httpClient  httpclient.HTTPClient
	redirectURL string
}

//New creates a new digilocker client
func New(appId, appKey string, hvEndpoint string, nr newrelic.Agent, client httpclient.HTTPClient, redirectURL string) *DigilockerImpl {
	dl := DigilockerImpl{
		appId:       appId,
		appKey:      appKey,
		hvEndpoint:  hvEndpoint,
		nr:          nr,
		httpClient:  client,
		redirectURL: redirectURL,
	}
	return &dl
}

func (dl *DigilockerImpl) GetRedirectURL() string {
	return dl.redirectURL
}

func (dl *DigilockerImpl) addHeaders(req *http.Request) {
	req.Header.Add("appId", dl.appId)
	req.Header.Add("appKey", dl.appKey)
	req.Header.Add("Content-Type", "application/json")

}
func (dl *DigilockerImpl) StartKYC(transactionId, referenceId, redirectURL string) (*KYCStartDetails, error) {
	tr := dl.nr.StartTransaction(HV_START_KYC_CALL)
	defer tr.End()
	url := dl.hvEndpoint + "/api/digilocker/start"
	method := "POST"

	reqObj, _ := json.Marshal(&KYCStartRequest{ReferenceId: referenceId, RedirectURL: redirectURL})
	payload := strings.NewReader(string(reqObj))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}
	dl.addHeaders(req)
	req.Header.Add("transactionId", transactionId)

	res, err := dl.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingHyperverge, err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadResponseBody, err.Error())
	}
	var result KYCStartResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalJson, err.Error())
	}
	switch result.StatusCode {
	case "200":
		return &result.Result, nil
	case "400":
		return nil, errors.Wrap(ErrReqValidate, result.Error.Message)
	case "500":
		return nil, errors.Wrap(ErrHVServer, result.Error.Message)
	}
	return &result.Result, nil
}

func (dl *DigilockerImpl) CheckAccountstatus(mobile, aadhaar string) (*AccountStatusDetails, error) {
	tr := dl.nr.StartTransaction(HV_ACCOUNT_STATUS_CALL)
	defer tr.End()
	url := dl.hvEndpoint + "/api/digilocker/accountStatus"
	method := "POST"

	reqObj, _ := json.Marshal(&AccountStatusRequest{Mobile: mobile, Aadhaar: aadhaar})
	payload := strings.NewReader(string(reqObj))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}
	dl.addHeaders(req)

	res, err := dl.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingHyperverge, err.Error())

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadResponseBody, err.Error())
	}
	var result AccountStatusResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalJson, err.Error())
	}
	switch result.StatusCode {
	case "200":
		return &result.Result, nil
	case "400":
		return nil, errors.Wrap(ErrReqValidate, result.Error.Message)
	case "500":
		return nil, errors.Wrap(ErrHVServer, result.Error.Message)
	}
	return &result.Result, nil
}

func (dl *DigilockerImpl) GetAddharDetails(transactionId, referenceId string) (*AadhaarDetails, error) {
	tr := dl.nr.StartTransaction(HV_ADDHAR_DETAILS_CALL)
	defer tr.End()
	url := dl.hvEndpoint + "/api/digilocker/eAadhaarDetails"
	method := "POST"

	reqObj, _ := json.Marshal(&EAadhaarDetailsRequest{ReferenceId: referenceId, AadhaarFile: "yes"})
	payload := strings.NewReader(string(reqObj))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}
	dl.addHeaders(req)
	req.Header.Add("transactionId", transactionId)

	res, err := dl.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingHyperverge, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadResponseBody, err.Error())
	}
	var result AadhaarDetailsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalJson, err.Error())
	}
	switch result.StatusCode {
	case "200":
		return &result.Result, nil
	case "504":
		return nil, errors.Wrap(ErrHVServer, result.Error.Message)
	case "500":
		return nil, errors.Wrap(ErrHVServer, result.Error.Message)
	}
	return &result.Result, nil
}

func (dl *DigilockerImpl) Healthcheck() (*HealthcheckResult, error) {
	tr := dl.nr.StartTransaction(HV_HEALTHCHECK_CALL)
	defer tr.End()
	url := dl.hvEndpoint + "/api/health/digilocker/accountStatus"

	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}
	dl.addHeaders(req)
	res, err := dl.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingHyperverge, err.Error())

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.Wrap(ErrHVServer, fmt.Sprintf("status code %d", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadResponseBody, err.Error())
	}

	var result HypervergeHealthcheckResponse
	log.Print(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalJson, err.Error())
	}
	return &result.Result, nil
}
