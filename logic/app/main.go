package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	controller "github.com/stepundel1/E-commerce/Users/logic/controller/gRPC"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
	"github.com/stepundel1/E-commerce/pkg/postgres"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	// 50051 port connection
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("err %v", err)
	}

	// import .env
	err = godotenv.Load("/Users/stepansalikov/Desktop/E-commerce/.env")
	if err != nil {
		log.Fatalf("err %v", err)
	}

	// import DB_CONNECTION postgres variable
	connection := os.Getenv("DB_CONNECTION")
	if connection == "" {
		log.Fatal("err")
	}

	// db connection
	pool, err := postgres.New(connection)
	if err != nil {
		log.Fatalf("err %v", err)
	}

	// new user repo
	userRepo := repo.NewUserRepo(pool)

	// grpc server connection
	s := grpc.NewServer()

	// GreeterServer
	registerServer := controller.NewRegisterServer(userRepo)
	pb.RegisterGreeterServer(s, registerServer)

	log.Printf("Server %v", lis.Addr())

	// Запуск gRPC сервера
	if err := s.Serve(lis); err != nil {
		log.Fatalf("err %v", err)
	}
}
