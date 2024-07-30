clean:
	$(info Cleaning DIO intermediate files...)
	go clean

build: clean
	$(info DIO build started...)
	go build -o ./bin/dio.exe cmd/dio/main.go

run:
	go run cmd/dio/main.go

start:
	./bin/dio.exe

all: clean build start