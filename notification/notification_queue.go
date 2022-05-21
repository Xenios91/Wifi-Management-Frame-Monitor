package notification

import (
	"encoding/json"
	"log"
	"sync"
	"time"
	"wifi-management-frame-monitor/management_frame"
)

type NotificationQueue struct {
	notifications map[string]map[string]Notification
}

type Notification struct {
	NotificationType string
	AssociatedEssid  string
	Time             time.Time
}

var (
	once               sync.Once
	Notification_Queue *NotificationQueue
)

func NewNotificationQueue() *NotificationQueue {
	once.Do(func() {
		notifications := make(map[string]map[string]Notification)
		Notification_Queue = &NotificationQueue{notifications}
	})
	return Notification_Queue
}

func (nq *NotificationQueue) AddNotification(notificationType string, mf *management_frame.ManagementFrame) {
	notification := &Notification{NotificationType: notificationType, AssociatedEssid: mf.Essid, Time: time.Now()}
	if _, ok := nq.notifications[mf.Essid]; !ok {
		nq.notifications[mf.Essid] = make(map[string]Notification)
	}
	nq.notifications[mf.Essid][notificationType] = *notification
}

func (nq *NotificationQueue) GetNotifications() *string {
	bytes, err := json.Marshal(nq.notifications)
	if err != nil {
		log.Println(err)
	}
	jsonBytes := string(bytes)
	return &jsonBytes
}
