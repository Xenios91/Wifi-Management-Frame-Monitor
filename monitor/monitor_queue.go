package monitor

import (
	"fmt"
	"sync"
	management_frame "wifi-management-frame-monitor/management_frame"
)

var (
	once          sync.Once
	Monitor_Queue MonitorQueue
)

type MonitorQueue struct {
	queue       chan *management_frame.ManagementFrame
	activityLog map[string]accessPoint
}

func New() *MonitorQueue {
	once.Do(func() {
		queue := make(chan *management_frame.ManagementFrame)
		Monitor_Queue = MonitorQueue{queue: queue, activityLog: make(map[string]accessPoint)}
		go Monitor_Queue.startService()
	})
	return &Monitor_Queue
}

func (mq *MonitorQueue) AddToQueue(mf *management_frame.ManagementFrame) {
	mq.queue <- mf
}

func (mq *MonitorQueue) auditManagementFrame(mf *management_frame.ManagementFrame) {
	switch mf.FrameType {
	case "deauth":
		isDeauthAttack := checkForDeauth(mf)
		fmt.Println(isDeauthAttack)
		//send alert
	case "beacon":
		isRogue := checkForRogue(mf)
		fmt.Println(isRogue)
		//send alert
	default:
		break
	}
}

func (mq *MonitorQueue) startService() {
	for {
		var mf *management_frame.ManagementFrame = <-Monitor_Queue.queue
		mq.auditManagementFrame(mf)
	}
}
