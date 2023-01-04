package hooks

import (
	"context"

	nrf "github.com/newrelic/go-agent/v3/newrelic"
)

func PushMetrics(ctx context.Context, msg string) error {
	nTxn := nrf.FromContext(ctx)

	nTxn.RecordLog(nrf.LogData{
		Message: msg,
	})
	return nil
}
