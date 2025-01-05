package controller

import (
	"context"
	"log"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/webapi"
)

type GRPCServer struct {
	pb.UnimplementedGreeterServer
	UserUseCase *webapi.UserUseCase
}

// NewRegisterServer new server
func NewGRPCServer(userUseCase *webapi.UserUseCase) *GRPCServer {
	return &GRPCServer{UserUseCase: userUseCase}
}

func (server *GRPCServer) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	err := server.UserUseCase.LogIn(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		log.Printf("failed to login user: %v", err)
		return &pb.LoginUserResponse{Success: false}, err
	}

	log.Printf("User login successfully: %v", in.GetEmail())
	return &pb.LoginUserResponse{Success: true}, nil

}

func (server *GRPCServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		PasswordHash: in.GetPassword(),
	}

	err := server.UserUseCase.Register(ctx, user, in.GetPassword())
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("User registered successfully: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil
}
