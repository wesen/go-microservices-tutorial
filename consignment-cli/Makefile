build:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o consignment-cli cli.go
	docker build -t shippy-cli-consignment -f Dockerfile.fast .

run:
	docker run -e MICRO_SERVER_ADDRESS=:50051 -p 50051:50051 shippy-clirvice-consignment
