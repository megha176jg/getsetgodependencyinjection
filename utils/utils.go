package utils

import "time"

func ContainString(l []string, s string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}
	return false
}

type RequestContext struct {
	SessionID     string
	RequestID     string
	ClientAppID   string
	UserID        string
	TransactionID string
	Method        string
	URI           string
	APIStartTime  time.Time
	IP            string
}
