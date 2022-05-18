package monitor

import "time"

type accessPoint struct {
	accessPointName string
	bssid           []string
	deauthCount     uint
	deauthStartTime time.Time
}
