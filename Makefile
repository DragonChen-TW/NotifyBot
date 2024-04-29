OUTPUT = bin

default: run
run:
	go run cmd/run/run.go
serve:
	go run cmd/serve/serve.go
build:
	go build -o $(OUTPUT)/run cmd/run/run.go
	go build -o $(OUTPUT)/serve cmd/serve/serve.go
