package deposit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/newrelic"

	"github.com/pkg/errors"
)

type DepositSDK struct {
	nr         newrelic.Agent
	httpClient httpclient.HTTPClient
	config     DepositConfig
}

func New(config DepositConfig, nr newrelic.Agent, httpClient httpclient.HTTPClient) *DepositSDK {
	return &DepositSDK{config: config, nr: nr, httpClient: httpClient}
}

func (ds *DepositSDK) GetFirstDepositFromHouzat(mobile string) (*DepositResponse, error) {
	tr := ds.nr.StartTransaction(GET_FIRST_DEPOSIT_FROM_HOUZAT_CALL)
	defer tr.End()

	url := fmt.Sprintf(ds.config.GetDepositEndpoint()+"/payment-service/pd/getDepositSummary?userId=&mobileNo=%s", mobile)
	method := "GET"
	log.Print(url)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("apiKey", ds.config.GetDepositAPIKey())
	req.Header.Add("merchantName", "playfantasy")
	req.Header.Add("role", "ADMIN")
	req.Header.Add("token", ds.config.GetDepositAuthToken())
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
