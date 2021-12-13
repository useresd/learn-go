build:
	go build -o build/server cmd/server/main.go
http:
	go run cmd/server/main.go