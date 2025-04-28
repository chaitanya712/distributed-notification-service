package worker

import (
	"log"
	"math/rand"
	"time"

	"github.com/chaitanya712/distributed-notification-service/internal/data"
)

func ProcessNotifications(ch <-chan data.Notification) {
	for n := range ch {
		retry := 0
		success := false
		for retry < 3 && !success {
			if rand.Float32() < 0.9 {
				data.AddNotification(n.UserID, n)
				log.Printf("Notification sent to %s for Post %s", n.UserID, n.PostID)
				success = true
			} else {
				time.Sleep(time.Duration((1<<retry)*100) * time.Millisecond)
				retry++
			}
		}
		if !success {
			log.Printf("Failed to deliver notification to %s for Post %s", n.UserID, n.PostID)
		}
	}
}
