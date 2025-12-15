package server

import (
	"context"
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
func (s *Server) Run(ctx context.Context) error {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil

		}

	}
}
