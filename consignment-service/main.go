package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"sync"

	pb "github.com/wesen/go-microservices-tutorial/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Consignment = consignment
	res.Created = true
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
