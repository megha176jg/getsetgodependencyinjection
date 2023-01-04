package hooks

import (
	"context"
	"fmt"

	nrf "github.com/newrelic/go-agent/v3/newrelic"
)

func PushMetrics(ctx context.Context, msg string) error {
	nTxn := nrf.FromContext(ctx)

	if nTxn == nil {
		return fmt.Errorf("unable to get newrelic txn from context")
	}
	nTxn.RecordLog(nrf.LogData{
		Message: msg,
	})

	return nil
}
