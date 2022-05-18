package monitor

import (
	"time"
	"wifi-management-frame-monitor/management_frame"
)

func checkForRogue(management_frame *management_frame.ManagementFrame) bool {
	apOfInterest := Monitor_Queue.activityLog[management_frame.APName]
	for i := 0; i < len(apOfInterest.bssid); i++ {
		if apOfInterest.bssid[i] == management_frame.BSSID {
			return false
		}
	}
	return true
}

func checkForDeauth(management_frame *management_frame.ManagementFrame) bool {
	monitor_item := Monitor_Queue.activityLog[management_frame.APName]
	currentDeauth := monitor_item.deauthCount
	if currentDeauth >= 50 {
		monitor_item.deauthCount = 0
		Monitor_Queue.activityLog[management_frame.APName] = monitor_item
		return true
	}

	if monitor_item.deauthCount == 0 {
		monitor_item.deauthStartTime = time.Now()
	}

	if monitor_item.deauthStartTime.Add(time.Second * 10).Before(time.Now()) {
		monitor_item.deauthCount++
	}

	Monitor_Queue.activityLog[management_frame.APName] = monitor_item
	return false
}
