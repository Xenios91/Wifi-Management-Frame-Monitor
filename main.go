package main

import (
	"log"
	"net/http"
	management_frame "wifi-management-frame-monitor/management_frame"
	monitor "wifi-management-frame-monitor/monitor"
)

var Monitor_Queue = monitor.New()

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		mf := management_frame.New("")
		Monitor_Queue.AddToQueue(mf)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/addToQueue", handleRequest)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
