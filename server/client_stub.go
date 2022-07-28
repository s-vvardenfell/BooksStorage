package server

import (
	"fmt"

	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientStub struct {
	storage.BooksStorageClient
}

func NewClientStub(host, port string) *ClientStub {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("cannot connect to host <%s> and port <%s>: %v", host, port, err)
	}
	return &ClientStub{
		storage.NewBooksStorageClient(conn),
	}
}
