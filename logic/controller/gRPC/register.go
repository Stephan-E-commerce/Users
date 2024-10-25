package gRPC

import (
	"context"
	"log"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/webapi"
)

// RegisterServer methods gRPC
type RegisterServer struct {
	pb.UnimplementedGreeterServer
	userUseCase *webapi.UserUseCase
}

// NewRegisterServer new server
func NewRegisterServer(userUseCase *webapi.UserUseCase) *RegisterServer {
	return &RegisterServer{userUseCase: userUseCase}
}

// RegisterUser func
func (s *RegisterServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		PasswordHash: in.GetPassword(),
	}

	err := s.userUseCase.Register(ctx, user, in.GetPassword())
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("User registered successfully: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil
}
