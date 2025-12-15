package server

import (
	"context"
	"fmt"
	"strings"

	"sync"
	"time"

	"github.com/M-oses340/Microservices-Database-Setup/ecomm-grpc/pb"
	"golang.org/x/sync/semaphore"
	"gopkg.in/gomail.v2"
)

type AdminInfo struct {
	Email    string
	Password string
}

type Server struct {
	client    pb.EcommClient
	adminInfo *AdminInfo
}

func NewServer(client pb.EcommClient, adminInfo *AdminInfo) *Server {
	return &Server{
		client:    client,
		adminInfo: adminInfo,
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
	sem := semaphore.NewWeighted(10)
	for _, ev := range res.Events {
		wg.Add(1)
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}
		go func(ev *pb.NotificationEvent) {
			defer wg.Done()
			defer sem.Release(1)

		}(ev)

	}
	go func() {
		wg.Wait()
	}()
	return nil
}
func (s *Server) sendNotification(ctx context.Context, ev *pb.NotificationEvent) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.adminInfo.Email)
	m.SetHeader("To", ev.UserEmail)
	m.SetHeader("Subject", "email from ecomm")
	m.SetBody("text/plain", fmt.Sprintf("Order %d is %s", strings.ToLower(ev.OrderStatus.String())))

	return nil

}
