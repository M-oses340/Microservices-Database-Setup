package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/M-oses340/Microservices-Database-Setup/ecomm-grpc/pb"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		grpcAddr = envflag.String(
			"GRPC_SVC_ADDR",
			"0.0.0.0:9091",
			"address where ecomm-grpc service is listening",
		)
	)
	envflag.Parse()

	log.Printf("connecting to ecomm-grpc at %s", *grpcAddr)

	// ---- gRPC connection ----
	conn, err := grpc.Dial(
		*grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial gRPC service: %v", err)
	}
	defer conn.Close()

	// ---- gRPC client ----
	client := pb.NewEcommClient(conn)

	// ---- Context with graceful shutdown ----
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	log.Println("ecomm-notification service started")

	// ---- Example: block until shutdown ----
	<-ctx.Done()

	log.Println("ecomm-notification service stopped")
	_ = client // keeps client available for future notification logic
}
