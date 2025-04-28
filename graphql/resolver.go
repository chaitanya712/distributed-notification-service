package graphql

import (
	"context"

	"github.com/chaitanya712/distributed-notification-service/internal/data"
	"github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

type GetNotificationsArgs struct {
	UserID string
}

type NotificationResolver struct {
	Notification *data.Notification
}

func (r *NotificationResolver) Id() graphql.ID {
	return graphql.ID(r.Notification.ID)
}

func (r *NotificationResolver) UserId() graphql.ID {
	return graphql.ID(r.Notification.UserID)
}

func (r *NotificationResolver) PostID() graphql.ID {
	return graphql.ID(r.Notification.PostID)
}

func (r *NotificationResolver) Message() string {
	return r.Notification.Message
}

func (r *Resolver) GetNotifications(ctx context.Context, args GetNotificationsArgs) ([]*NotificationResolver, error) {
	ns := data.NotificationStore[args.UserID]
	if len(ns) > 20 {
		ns = ns[len(ns)-20:]
	}
	var result []*NotificationResolver
	for _, n := range ns {

		result = append(result, &NotificationResolver{&n})
	}
	return result, nil
}
