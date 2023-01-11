package main

import (
	"log"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
	"bitbucket.org/junglee_games/getsetgo/sdks/deposit"
)

type Conf struct{}

func (c *Conf) GetDepositEndpoint() string {
	return "http://payments-hzt-qa-1.howzatfantasy.com"
}
func (c *Conf) GetDepositAuthToken() string {
	return "test"
}
func (c *Conf) GetDepositAPIKey() string {
	return "qJ3ycOUJ88UnZE5HkUQ9Tg=="
}
func main() {

	depositSdk := deposit.New(&Conf{}, newrelic.Agent{}, httpclient.NewHttpClient(60))
	res, err := depositSdk.GetFirstDepositFromHouzat("5455445566")
	log.Print(res, err)
}
