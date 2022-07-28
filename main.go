package main

import (
	"fmt"
	"net"

	"github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/s-vvardenfell/BooksStorage/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServ := grpc.NewServer()
	rcs := server.New("mysql", "s.vvardenfell:Zxasqw12@/kvadoru")
	books_storage.RegisterBooksStorageServer(grpcServ, rcs)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", "50051"))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(grpcServ)
	// if cnfg.WithReflection {
	// 	reflection.Register(grpcServ)
	// }

	logrus.Info("Starting gRPC listener on port " + "50051")
	if err := grpcServ.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
