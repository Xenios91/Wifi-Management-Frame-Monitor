package monitor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type accessPointData struct {
	bssid           []string
	essid           string
	deauthCount     uint
	deauthStartTime time.Time
}

type AccessPoint struct {
	Bssid []string
	Essid string
}

func LoadAccessPoints() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	var accessPoints []AccessPoint
	err = json.Unmarshal([]byte(file), &accessPoints)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(accessPoints); i++ {
		accessPoint := accessPoints[i]

		if accessPointDataRetrieved, ok := Monitor_Queue.activityLog[essid(accessPoint.Essid)]; ok {
			accessPointDataRetrieved.bssid = append(accessPointDataRetrieved.bssid, accessPoint.Bssid...)
		} else {
			Monitor_Queue.activityLog[essid(accessPoint.Essid)] = accessPointData{bssid: accessPoint.Bssid, essid: accessPoint.Essid, deauthCount: 0}
		}
	}
}
