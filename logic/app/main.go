package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	userRepo repo.UserRepo
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	user := entity.User{
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		PasswordHash: in.GetPassword(),
	}

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("User registered successfully: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil

}
func (s *server) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := s.userRepo.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		log.Printf("failed to find user: %v", err)
		return &pb.LoginUserResponse{Success: false}, err
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
	if err != nil {
		log.Printf("invalid password for user: %v", in.GetEmail())
		return &pb.LoginUserResponse{Success: false}, nil
	}

	log.Printf("User logged in successfully: %v", in.GetEmail())
	return &pb.LoginUserResponse{Success: true}, nil
}
