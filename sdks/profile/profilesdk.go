package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/newrelic"

	"github.com/pkg/errors"
)

type ProfileSDK struct {
	config     ProfileConfig
	nr         newrelic.Agent
	httpClient httpclient.HTTPClient
}

func New(config ProfileConfig, nr newrelic.Agent, httpClient httpclient.HTTPClient) *ProfileSDK {
	return &ProfileSDK{
		config:     config,
		nr:         nr,
		httpClient: httpClient,
	}
}

func (psdk *ProfileSDK) GetUserByID(userId int) (*ProfileResponse, error) {
	url := fmt.Sprintf(psdk.config.GetProfileEndpoint()+"?id=%s", userId)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}

	req.Header.Add("Authorization", psdk.config.GetJWRAuthToken())
	res, err := psdk.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingProfile, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.Wrap(ErrStatusCodeOtherThan200, fmt.Sprintf("returned status code %d", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadingBody, err.Error())
	}
	var profile ProfileResponse
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshlingResponse, err.Error())
	}
	return &profile, nil
}
