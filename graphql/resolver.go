package graphql

import (
	"context"

	"github.com/chaitanya712/distributed-notification-service/internal/data"
)

type Resolver struct{}

func (r *Resolver) GetNotifications(ctx context.Context, userId string) ([]*data.Notification, error) {
	ns := data.NotificationStore[userId]
	if len(ns) > 20 {
		ns = ns[len(ns)-20:]
	}
	var result []*data.Notification
	for _, n := range ns {
		nCopy := n
		result = append(result, &nCopy)
	}
	return result, nil
}
