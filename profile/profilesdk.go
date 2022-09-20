package profile

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

type ProfileSDK struct {
	endpoint   string
	nr         newrelic.Agent
	httpClient httpclient.HTTPClient
	authToken  string
}

func New(endpoint string, nr newrelic.Agent, httpClient httpclient.HTTPClient, authToken string) *ProfileSDK {
	return &ProfileSDK{
		endpoint:   endpoint,
		nr:         nr,
		httpClient: httpClient,
		authToken:  authToken,
	}
}

func (psdk *ProfileSDK) GetUserByID(userId int) (*ProfileResponse, error) {
	return &ProfileResponse{RegistrationTime: time.Now().Unix(), UserId: userId, Mobile: "7040309988", FirstDepositAmount: 100, FirstDepositDate: time.Now().Unix()}, nil
	url := psdk.endpoint + "?id=123"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(ErrCreatingRequest, err.Error())
	}

	req.Header.Add("Authorization", psdk.authToken)
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
