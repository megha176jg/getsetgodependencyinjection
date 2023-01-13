package monitoring

import (
	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
	"github.com/pkg/errors"
)

type Factory struct {
	agents map[string]Agent
	cfg    config
}

func NewMonitoringFactory(cfg config) *Factory {
	return &Factory{cfg: cfg, agents: make(map[string]Agent)}
}

func setupNewRelic(name, key string) (Agent, error) {
	a, err := newrelic.New(name, key)
	if err != nil {
		return nil, errors.Wrap(err, "main: while starting newrelic")
	}
	return a, nil
}

func (f *Factory) GetMonitoringAgent(name string) (Agent, error) {
	switch name {
	case NEWRELIC:
		if agent, exists := f.agents[NEWRELIC]; exists {
			return agent, nil
		}
		agent, err := setupNewRelic(f.cfg.GetServiceName(), f.cfg.GetNewRelicKey())
		if err != nil {
			return nil, err
		}
		f.agents[NEWRELIC] = agent
		return agent, nil
	default:
		return nil, ErrInvalidMonitoringAgent
	}
}
