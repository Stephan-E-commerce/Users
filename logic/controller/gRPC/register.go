package grpcc

import (
	"context"
	"log"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
)

// Реализация gRPC сервера
type RegisterServer struct {
	pb.UnimplementedGreeterServer
	userRepo *repo.UserRepo
}

// Конструктор для создания нового сервера
func NewServer(userRepo *repo.UserRepo) *RegisterServer {
	return &RegisterServer{userRepo: userRepo}
}

// Метод для регистрации пользователя
func (s *RegisterServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		PasswordHash: in.GetPassword(),
	}

	// Создание пользователя через репозиторий
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		log.Printf("Не удалось создать пользователя: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("Пользователь успешно зарегистрирован: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil
}
