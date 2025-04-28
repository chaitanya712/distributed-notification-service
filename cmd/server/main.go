package main

import (
	"log"
	"net"
	"net/http"

	"github.com/chaitanya712/distributed-notification-service/graphql"
	"github.com/chaitanya712/distributed-notification-service/grpc"
	"github.com/chaitanya712/distributed-notification-service/internal/queue"
	"github.com/chaitanya712/distributed-notification-service/metrics"
	pb "github.com/chaitanya712/distributed-notification-service/proto"
	gql "github.com/graph-gophers/graphql-go"
	graphqlhandler "github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
)

func main() {
	metrics.Init()
	queue.StartQueue()

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterNotificationServiceServer(grpcServer, &grpc.NotificationServer{})
		log.Println("gRPC server running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start GraphQL server
	schema := gql.MustParseSchema(graphql.Schema, &graphql.Resolver{})
	http.Handle("/graphql", &graphqlhandler.Handler{Schema: schema})
	log.Println("GraphQL server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
