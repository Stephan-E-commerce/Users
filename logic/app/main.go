package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	grpcc "github.com/stepundel1/E-commerce/Users/logic/controller/gRPC"
	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"github.com/stepundel1/E-commerce/Users/logic/usecase/repo"
	"github.com/stepundel1/E-commerce/pkg/postgres"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	// Настройка прослушивания на порту 50051
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Не удалось прослушать: %v", err)
	}

	// Загрузка .env файла
	err = godotenv.Load("/Users/stepansalikov/Desktop/E-commerce/.env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	// Получение строки подключения к базе данных из переменной окружения
	connection := os.Getenv("DB_CONNECTION")
	if connection == "" {
		log.Fatal("Переменная окружения DB_CONNECTION не установлена")
	}

	// Подключение к базе данных
	pool, err := postgres.New(connection)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Создание репозитория пользователей
	userRepo := repo.NewUserRepo(pool)

	// Создание gRPC сервера
	s := grpc.NewServer()

	// Регистрация сервиса GreeterServer
	registerServer := grpcc.NewServer(userRepo) // Используем конструктор
	pb.RegisterGreeterServer(s, registerServer)

	log.Printf("Сервер слушает на %v", lis.Addr())

	// Запуск gRPC сервера
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
