package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	management_frame "wifi-management-frame-monitor/management_frame"
	monitor "wifi-management-frame-monitor/monitor"
)

var Monitor_Queue = monitor.New()

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var mf management_frame.ManagementFrame
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(body, &mf)
		if err != nil {
			log.Fatalln(err)
		}

		Monitor_Queue.AddToQueue(&mf)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/addToQueue", handleRequest)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
