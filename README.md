# Distributed Notification Service

## Setup

```bash
go mod tidy
go run cmd/server/main.go
```

## gRPC
Service is exposed on `localhost:50051`

### `PublishPost`
Send a post and notify followers using:
```proto
rpc PublishPost(Post) returns (NotificationResponse)
```

### Test with grpcurl
```bash
grpcurl -plaintext -d '{"id": "p99", "user_id": "u1", "body": "My new test post"}' localhost:50051 notification.NotificationService/PublishPost
```

## GraphQL
Available at `http://localhost:8080/graphql`

Example Query:
```graphql
{
  getNotifications(userId: "u2") {
    id
    message
    timestamp
  }
}
```
