package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Mizuhar4/tracking-service/pb"
)

type server struct {
	pb.UnimplementedTrackingServiceServer
}

func (s *server) UpdateDriverLocation(ctx context.Context, req *pb.UpdateDriverLocationRequest) (*pb.UpdateDriverLocationResponse, error) {
	return &pb.UpdateDriverLocationResponse{
		DeliveryId: req.DeliveryId,
		UpdatedAt:  "2026-10-07T23:30:00Z",
	}, nil
}

func (s *server) AssignDelivery(ctx context.Context, req *pb.AssignDeliveryRequest) (*pb.AssignDeliveryResponse, error) {
	return &pb.AssignDeliveryResponse{
		AssignmentId: "assign-uuid-001",
		OrderId:      req.OrderId,
		DeliveryId:   "driver-uuid-99",
		Status:       "ASSIGNED",
	}, nil
}

func (s *server) GetOrderLocation(ctx context.Context, req *pb.GetOrderLocationRequest) (*pb.GetOrderLocationResponse, error) {
	return &pb.GetOrderLocationResponse{
		OrderId:     req.OrderId,
		DeliveryId:  "driver-uuid-99",
		Latitude:    -38.7396, // Coordenadas de referencia (ej. Temuco)
		Longitude:   -72.5984,
		LastUpdated: "2026-10-07T23:30:00Z",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052") // Usamos el puerto 50052 para este servicio
	if err != nil {
		log.Fatalf("no se pudo escuchar: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterTrackingServiceServer(srv, &server{})

	log.Println("Servidor gRPC de Tracking & Delivery Service escuchando en :50052")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
