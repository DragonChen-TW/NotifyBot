default: run
run:
	go run cmd/run/run.go
build:
	go build -o run cmd/run/run.go