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
	time.NewTicker(30 * time.Second)
	for {
		
	}
}
