package idfy

import (
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
	logger "bitbucket.org/junglee_games/getsetgo/logger_old"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type idfyConfig struct {
	AccountId string
	ApiKey    string
	EndPoint  string
}

func (config idfyConfig) GetIdfyAccountId() string {
	return config.AccountId
}
func (config idfyConfig) GetIdfyApiKey() string {
	return config.ApiKey
}
func (config idfyConfig) GetIdfyEndpoint() string {
	return config.EndPoint
}
func main() {
	cfg := idfyConfig{
		AccountId: "provided",
		ApiKey:    "provided",
		EndPoint:  "localhost:8080",
	}
	logger := *logger.Logger
	nr, err := setupNewRelic(logger, "", "")
	if err != nil {
		panic(err)
	}
	httpclient := http.Client{}
	idfyClient := New(cfg, nr, &httpclient)
	idreq := IdfyRequest{}
	idfyClient.ExtractAadhar("aadhar", idreq)
}

func setupNewRelic(logger zap.Logger, name, key string) (newrelic.Agent, error) {
	a, err := newrelic.New(name, key)
	if err != nil {
		return newrelic.Agent{}, errors.Wrap(err, "main: error occured while starting newrelic")
	}
	return *a, nil
}
