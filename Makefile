OUTPUT = bin

default: run
run:
	go run cmd/run/run.go
build:
	go build -o $(OUTPUT)/run cmd/run/run.go