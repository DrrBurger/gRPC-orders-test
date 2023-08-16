package main

import (
	"context"
	pb "gRPC-orders-test/proto/order_service"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	orders      map[string][]string
	lastOrderID int
}

func (s *server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	s.lastOrderID++
	orderID := "order_" + strconv.Itoa(s.lastOrderID)
	s.orders[in.UserId] = append(s.orders[in.UserId], orderID)
	return &pb.OrderResponse{OrderId: orderID}, nil
}

func (s *server) GetAllOrders(ctx context.Context, in *pb.GetAllOrdersRequest) (*pb.AllOrdersResponse, error) {
	return &pb.AllOrdersResponse{OrderIds: s.orders[in.UserId]}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{orders: make(map[string][]string)})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
