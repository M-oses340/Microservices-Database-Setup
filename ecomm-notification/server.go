package server

import (
	"context"
	"sync"
	"time"

	"github.com/M-oses340/Microservices-Database-Setup/ecomm-grpc/pb"
)

type Server struct {
	client pb.EcommClient
}

func NewServer(client pb.EcommClient) *Server {
	return &Server{
		client: client,
	}
}
func (s *Server) Run(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return

		}

	}
}
func (s *Server) processNotificationEvents(ctx context.Context) error {
	res, err := s.client.ListNotificationEvents(ctx, &pb.ListNotificationEventsReq{})
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, ev := range res.Events {
		wg.Add(1)
		go func(ev *pb.NotificationEvent) {

		}(ev)

	}
	return nil
}
