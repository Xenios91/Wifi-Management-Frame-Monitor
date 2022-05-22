package monitor

import (
	"strings"
	"sync"
	management_frame "wifi-management-frame-monitor/management_frame"
	"wifi-management-frame-monitor/notification"
)

var (
	once          sync.Once
	Monitor_Queue *MonitorQueue
	nq            *notification.NotificationQueue = notification.NewNotificationQueue()
)

type essid string

type MonitorQueue struct {
	queue       chan *management_frame.ManagementFrame
	activityLog map[essid]accessPointData
}

func New() *MonitorQueue {
	once.Do(func() {
		queue := make(chan *management_frame.ManagementFrame)
		Monitor_Queue = &MonitorQueue{queue: queue, activityLog: make(map[essid]accessPointData)}
		go Monitor_Queue.startService()
	})
	return Monitor_Queue
}

func (mq *MonitorQueue) AddToQueue(mf *management_frame.ManagementFrame) {
	mq.queue <- mf
}

func (mq *MonitorQueue) auditManagementFrame(mf *management_frame.ManagementFrame) {
	switch strings.ToLower(mf.FrameType) {
	case "deauth":
		if isDeauthAttack := checkForDeauth(mf); isDeauthAttack {
			nq.AddNotification("deauth attack detected", mf)
		}

	case "beacon":
		if isRogue := checkForRogue(mf); isRogue {
			nq.AddNotification("rogue ap detected", mf)
		}

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
