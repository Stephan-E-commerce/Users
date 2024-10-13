package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/stepundel1/E-commerce/Users/logic/entity"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
	usecase "github.com/stepundel1/E-commerce/Users/logic/usecase/webapi"
	"github.com/stepundel1/E-commerce/pkg/postgres"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	userRepo    *repo.UserRepo
	userUseCase *usecase.UserUseCase
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("не удалось прослушать: %v", err)
	}

	err = godotenv.Load("/Users/stepansalikov/Desktop/E-commerce/.env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	connection := os.Getenv("DB_CONNECTION")
	if connection == "" {
		log.Fatal("DB_CONNECTION environment variable is not set")
	}

	// Убедитесь, что создаете экземпляр Postgres перед созданием UserRepo
	pool, err := postgres.New(connection) // Убедитесь, что это правильно
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	userRepo := repo.NewUserRepo(pool) // Передайте корректный pool здесь
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{userRepo: userRepo}) // Убедитесь, что сервер имеет userRepo
	log.Printf("сервер слушает на %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}
}

func (s *server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		Name:         in.GetName(),     // Использование proto для получения имени
		Email:        in.GetEmail(),    // Использование proto для получения email
		PasswordHash: in.GetPassword(), // Использование proto для получения пароля
	}

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
		return &pb.RegisterUserResponse{Success: false}, err
	}

	log.Printf("User registered successfully: %v", in.GetName())
	return &pb.RegisterUserResponse{Success: true}, nil // Возвращение proto ответа
}

// func (s *server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
// 	err := s.userUseCase.Register(ctx, in.GetEmail(), in.GetName(), in.GetPassword()) // Вызовите метод Register
// 	if err != nil {
// 		log.Fatalf("failed to register user: %v", err)
// 		return &pb.RegisterUserResponse{Success: false}, err
// 	}

// 	log.Printf("User registered successfully: %v", in.GetName())
// 	return &pb.RegisterUserResponse{Success: true}, nil // Возвращение proto ответа
// }
