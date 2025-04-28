package queue

import (
	"github.com/chaitanya712/distributed-notification-service/internal/data"
	"github.com/chaitanya712/distributed-notification-service/internal/worker"
)

var NotificationChannel = make(chan data.Notification, 100)

func EnqueueNotification(n data.Notification) {
	NotificationChannel <- n
}

func StartQueue() {
	for i := 0; i < 5; i++ {
		go worker.ProcessNotifications(NotificationChannel)
	}
}
