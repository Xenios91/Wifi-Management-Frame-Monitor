package notification

import (
	"encoding/json"
	"log"
	"sync"
)

type NotificationQueue struct {
	notifications []string
}

var (
	once               sync.Once
	Notification_Queue *NotificationQueue
)

func New() *NotificationQueue {
	once.Do(func() {
		notifications := make([]string, 0)
		Notification_Queue = &NotificationQueue{notifications}
	})
	return Notification_Queue
}

func (nq *NotificationQueue) addNotification(notification string) {
	nq.notifications = append(nq.notifications, notification)
}

func (nq *NotificationQueue) sendNotifications() *string {
	bytes, err := json.Marshal(nq.notifications)
	if err != nil {
		log.Fatalln(err)
	}
	jsonBytes := string(bytes)
	return &jsonBytes
}
