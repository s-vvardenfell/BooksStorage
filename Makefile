# Set env var with dsn to connect to database (need specify local db user/pass)
export DSN := user:pass@/kvadoru

# Run tests
tests:
	go test -v -cover ./...

# Compile protos
protos: 
	./gen.sh

# Run docker with database and server
dockerrun: protos
	docker-compose docker-compose.yml up --build

# Run server locally
server: protos
	go run main.go --config resources/config.yml