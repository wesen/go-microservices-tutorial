FROM golang:1.12 AS builder

WORKDIR /go
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN GO111MODULE=on go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
RUN apt-get update -y && apt-get install -y unzip

COPY . /go/src/github.com/wesen/go-microservices-tutorial
RUN sh /go/src/github.com/wesen/go-microservices-tutorial/scripts/get-protoc.sh
WORKDIR /go/src/github.com/wesen/go-microservices-tutorial/consignment-service
RUN protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
RUN GO111MODULE=on go get -u
RUN GOOS=linux GOARCH=amd64 go build -o consignment-service main.go

FROM alpine:latest
RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/wesen/go-microservices-tutorial/consignment-service/consignment-service /app/consignment-service

CMD ["./consignment-service"]