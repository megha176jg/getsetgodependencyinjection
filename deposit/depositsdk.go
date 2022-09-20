package deposit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/newrelic"

	"github.com/pkg/errors"
)

type DepositSDK struct {
	endpoint   string
	nr         newrelic.Agent
	authToken  string
	httpClient httpclient.HTTPClient
}

func New(endpoint string, nr newrelic.Agent, httpClient httpclient.HTTPClient, authToken string) *DepositSDK {
	return &DepositSDK{endpoint: endpoint, nr: nr, httpClient: httpClient, authToken: authToken}
}

func (ds *DepositSDK) GetFirstDepositFromHouzat(mobile string) (*DepositResponse, error) {
	return &DepositResponse{Amount: 6000, DepositedOn: time.Now().Unix()}, nil
	url := fmt.Sprintf(ds.endpoint+"?mobile=%s", mobile)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}

	req.Header.Add("Authorization", ds.authToken)
	res, err := ds.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingHouzatPaymentService, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.Wrap(ErrStatusCodeOtherThan200, fmt.Sprintf("Returned status code %d", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadingBody, err.Error())
	}
	var deposit DepositResponse
	err = json.Unmarshal(body, &deposit)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshlingResponse, err.Error())
	}
	return &deposit, nil

}
