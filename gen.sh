#!/bin/sh
protoc --go-grpc_out=. proto/books_storage.proto
protoc --go_out=. proto/books_storage.proto