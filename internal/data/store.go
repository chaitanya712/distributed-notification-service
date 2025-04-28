package data

import "sync"

var StoreLock = sync.Mutex{}

func AddNotification(userID string, n Notification) {
	StoreLock.Lock()
	defer StoreLock.Unlock()
	NotificationStore[userID] = append(NotificationStore[userID], n)
}
