package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	desc "github.com/KEmsO-DY/auth/pkg/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceServer struct {
	desc.UnimplementedAuthServiceServer
}

func (s *ServiceServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Got Create request with: {name:%v,email:%v,role:%v}", req.Name, req.Email, req.UserRole)
	if req.Password != req.PasswordConfirm {
		return &desc.CreateResponse{}, errors.New("Failure to confirm password")
	}
	return &desc.CreateResponse{
		Id: 1,
	}, nil
}

func (s *ServiceServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Got Get request with: id=%v", req.Id)
	RT := time.Now()
	return &desc.GetResponse{
		Id:        1,
		Name:      "void",
		Email:     "void",
		UserRole:  desc.Role_ROLE_STATUS_USER,
		CreatedAt: timestamppb.New(RT),
		UpdatedAt: timestamppb.New(RT),
	}, nil
}

func (s *ServiceServer) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	log.Printf("Got Update request with: id=%v, name=%v, email=%v", req.Id, req.Name, req.Email)
	if req.Email == nil || req.Name == nil {
		return &desc.UpdateResponse{}, errors.New("Cant update field with nil")
	}
	return &desc.UpdateResponse{}, nil
}

func (s *ServiceServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Printf("Got Get request with: id=%v", req.Id)
	return &desc.DeleteResponse{}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Cant listen this port")
	}
	s := grpc.NewServer()                               // Объект сервера, не слушает порт не знает про AuthService не принимает запросы
	desc.RegisterAuthServiceServer(s, &ServiceServer{}) // связывает с gRPC-сервером реализацию сервиса
	reflection.Register(s)                              //introspection
	log.Print("Starting gRPC server on :50051")
	if err := s.Serve(listner); err != nil {
		log.Fatal("Cant make up server")
	}
}
