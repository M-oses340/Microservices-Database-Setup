package server

import (
	"github.com/M-oses340/Microservices-Database-Setup/ecomm-grpc/pb"
)

type Server struct {
	client pb.EcommClient
}

func NewServer(client pb.EcommClient) *Server {
	return &Server{client: client}
}
