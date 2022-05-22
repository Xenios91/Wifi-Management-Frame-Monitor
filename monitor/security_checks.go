package monitor

import (
	"errors"
	"log"
	"strconv"
	"time"
	"wifi-management-frame-monitor/management_frame"
)

func checkForEvilTwin(management_frame *management_frame.ManagementFrame) bool {
	apOfInterest := Monitor_Queue.activityLog[essid(management_frame.Essid)]
	if _, ok := Monitor_Queue.activityLog[essid(apOfInterest.essid)]; !ok {
		return false
	}

	for i := 0; i < len(apOfInterest.bssid); i++ {
		if apOfInterest.bssid[i] == management_frame.Bssid {
			return false
		}
	}
	return true
}

func getTimeStamp(timeStamp *string) (time.Time, error) {
	var ts time.Time = time.Time{}
	timesStampUnix, err := strconv.Atoi(*timeStamp)
	if err != nil {
		log.Println(err)
		return ts, errors.New("invalid unix time")
	} else {
		ts = time.Unix(int64(timesStampUnix), 0)
	}
	return ts, nil
}

func checkForDeauth(management_frame *management_frame.ManagementFrame) bool {
	var ok bool
	var monitor_item accessPointData
	if monitor_item, ok = Monitor_Queue.activityLog[essid(management_frame.Essid)]; !ok {
		return false
	}

	currentDeauth := monitor_item.deauthCount
	if currentDeauth >= 50 {
		monitor_item.deauthCount = 0
		Monitor_Queue.activityLog[essid(management_frame.Essid)] = monitor_item
		return true
	}

	timesStampUnix, err := getTimeStamp(&management_frame.TimeStamp)
	if err != nil {
		log.Println(err)
	} else if monitor_item.deauthStartTime.IsZero() {
		monitor_item.deauthStartTime = timesStampUnix
		monitor_item.deauthCount++
		Monitor_Queue.activityLog[essid(management_frame.Essid)] = monitor_item
	}

	if monitor_item.deauthStartTime.Add(time.Second * 10).After(timesStampUnix) {
		monitor_item.deauthCount++
		Monitor_Queue.activityLog[essid(management_frame.Essid)] = monitor_item
	} else {
		monitor_item.deauthCount = 0
		monitor_item.deauthStartTime = timesStampUnix
		Monitor_Queue.activityLog[essid(management_frame.Essid)] = monitor_item
	}

	return false
}
