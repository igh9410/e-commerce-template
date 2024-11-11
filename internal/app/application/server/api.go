package server

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
	pb "github.com/igh9410/e-commerce-template/internal/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// API implements pb.ProductServiceServer
type API struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
}

func NewAPI(productService service.ProductService) *API {
	return &API{
		productService: productService,
	}
}

// RegisterAllServices registers all gRPC and gRPC-Gateway services
func RegisterAllServices(grpcServer *grpc.Server, mux *runtime.ServeMux, api *API) error {
	// Register gRPC services
	pb.RegisterProductServiceServer(grpcServer, api)

	// Add more services here as needed

	// Register gRPC-Gateway handlers
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, ":50051", opts); err != nil {
		log.Fatalf("Failed to register gRPC-Gateway handler: %v", err)
		return err
	}

	// Add more handlers here as needed

	return nil
}
