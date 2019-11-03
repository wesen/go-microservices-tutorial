build:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o consignment-service main.go
	docker build -t shippy-service-consignment -f Dockerfile.fast .

run:
	docker run -e MICRO_SERVER_ADDRESS=:50051 -p 50051:50051 shippy-service-consignment
