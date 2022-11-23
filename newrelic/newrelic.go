package newrelic

import (
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
)

type Agent interface {
	StartTransaction(key string) *newrelic.Transaction
	RecordCustomMetric(key string)
	RecordCustomEvent(key string, params map[string]interface{})
}

type NewrelicImpl struct {
	a *newrelic.Application
}

func New(name, key string) (*NewrelicImpl, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(name),
		newrelic.ConfigLicense(key),
	)
	if err != nil {
		return &NewrelicImpl{}, errors.Wrap(err, "initiating newrelic session")
	}
	// Wait for the application to connect.
	if err = app.WaitForConnection(30 * time.Second); nil != err {
		return nil, err
	}
	if err != nil {
		return &NewrelicImpl{}, errors.Wrap(err, "getting newrelic session")
	}
	return &NewrelicImpl{app}, nil
}

func (ag *NewrelicImpl) StartTransaction(key string) *newrelic.Transaction {
	return ag.a.StartTransaction(key)
}

func (ag *NewrelicImpl) RecordCustomMetric(key string) {
	go ag.a.RecordCustomMetric(key, 1)
}

func (ag *NewrelicImpl) RecordCustomEvent(key string, params map[string]interface{}) {
	go ag.a.RecordCustomEvent(key, params)
}
