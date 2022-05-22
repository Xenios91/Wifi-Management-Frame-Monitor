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

type Notification interface {
	SetNotificationType(*string)
	SetAssociatedEssid(*string)
	SetTimeStamp(time.Time)
}

type NotificationItem struct {
	NotificationType string
	AssociatedEssid  string
	TimeStamp        time.Time
}

func (notificationItem *NotificationItem) SetNotificationType(notificationType *string) {
	notificationItem.NotificationType = *notificationType
}

func (notificationItem *NotificationItem) SetAssociatedEssid(notificationEssid *string) {
	notificationItem.AssociatedEssid = *notificationEssid
}

func (notificationItem *NotificationItem) SetTimeStamp(notificationTimeStamp time.Time) {
	notificationItem.TimeStamp = notificationTimeStamp
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
	notification := &NotificationItem{NotificationType: notificationType, AssociatedEssid: mf.Essid, TimeStamp: time.Now()}
	if _, ok := nq.notifications[mf.Essid]; !ok {
		nq.notifications[mf.Essid] = make(map[string]Notification)
	}
	nq.notifications[mf.Essid][notificationType] = notification
}

func (nq *NotificationQueue) GetNotifications() *string {
	bytes, err := json.Marshal(nq.notifications)
	if err != nil {
		log.Println(err)
	}
	jsonBytes := string(bytes)
	return &jsonBytes
}
