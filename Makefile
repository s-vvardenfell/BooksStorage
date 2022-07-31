# Set env var with dsn to connect to database
export DSN := s.vvardenfell:Zxasqw12@/kvadoru

# Run tests (need specify user/pass)
tests:
	go test -v -cover ./...

# Compile protos
protos: 
	./gen.sh

# Run docker with database and server
dockerrun: protos
	docker-compose -f docker-compose.yml up --build

server: protos
	go run main.go --config resources/config.yml