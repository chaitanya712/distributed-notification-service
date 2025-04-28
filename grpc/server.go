package grpc

import (
	"context"
	"fmt"

	"github.com/chaitanya712/distributed-notification-service/internal/data"
	"github.com/chaitanya712/distributed-notification-service/internal/queue"
	pb "github.com/chaitanya712/distributed-notification-service/proto"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationServer) PublishPost(ctx context.Context, post *pb.Post) (*pb.NotificationResponse, error) {
	followers := data.Followers[post.UserId]
	total := 0
	for _, follower := range followers {
		n := data.Notification{
			ID:        fmt.Sprintf("n-%s-%s", follower, post.Id),
			UserID:    follower,
			PostID:    post.Id,
			Message:   fmt.Sprintf("New post from user %s", post.UserId),
			Timestamp: post.Timestamp.AsTime(),
		}
		queue.EnqueueNotification(n)
		total++
	}
	return &pb.NotificationResponse{
		Status:             "Queued",
		TotalNotifications: int32(total),
	}, nil
}
