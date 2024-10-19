package controller

import (
	"context"
	"log"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
)

// RegisterServer methods gRPC
type RegisterServer struct {
	pb.UnimplementedGreeterServer
	userRepo *repo.UserRepo
}

// NewRegisterServer new server
func NewRegisterServer(userRepo *repo.UserRepo) *RegisterServer {
	return &RegisterServer{userRepo: userRepo}
}

// RegisterUser func
func (s *RegisterServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		PasswordHash: in.GetPassword(),
	}

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("User registered successfully: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil
}
