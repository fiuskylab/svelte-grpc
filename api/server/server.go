package server

import (
	"context"
	"net"

	"github.com/fiuskylab/svelte-grpc/api/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type (
	// Product will store data related to Products.
	Product struct {
		ID   string
		Name string
	}

	routeProduct struct {
		proto.UnimplementedProductServer

		products []Product
	}
)

func newProductRoute() *routeProduct {
	return &routeProduct{
		products: []Product{},
	}
}

func (r *routeProduct) Create(_ context.Context, p *proto.CreateRequest) (*proto.CreateResponse, error) {
	id := uuid.NewString()
	r.products = append(r.products, Product{
		ID:   id,
		Name: p.Name,
	})

	return &proto.CreateResponse{
		Id:   id,
		Name: p.Name,
	}, nil
}

func (r *routeProduct) List(_ context.Context, _ *proto.Empty) (*proto.ListResponse, error) {
	resp := &proto.ListResponse{
		Products: []*proto.CreateResponse{},
	}

	for _, p := range r.products {
		resp.Products = append(resp.Products, &proto.CreateResponse{
			Id:   p.ID,
			Name: p.Name,
		})
	}

	return resp, nil
}

// Listen will stablish a gRPC server.
func Listen() {
	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProductServer(grpcServer, newProductRoute())
	grpcServer.Serve(lis)
}
