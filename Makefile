clean:
	go mod tidy
	go mod vendor
swag-gen:
	swag init -d cmd/
start:
	go run cmd/main.go
build:
	go build cmd/main.go