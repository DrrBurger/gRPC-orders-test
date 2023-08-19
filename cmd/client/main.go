package main

import (
	"context"
	"fmt"
	pb "gRPC-orders-test/proto/order_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// Установка соединение с сервером.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// Имитируем запрос на получение order_id для user_id "user_123"
	userID := "user_123"
	response, err := client.GetOrder(context.Background(), &pb.GetOrderRequest{UserId: userID})
	if err != nil {
		log.Fatalf("ошибка при получении order_id: %v", err)
	}
	fmt.Printf("Order ID для %s: %s\n", userID, response.GetOrderId())

	// Имитируем запрос на получение всех order_ids для user_id "user_123"
	allOrdersResponse, err := client.GetAllOrders(context.Background(), &pb.GetAllOrdersRequest{UserId: userID})
	if err != nil {
		log.Fatalf("ошибка при получении всех order_ids: %v", err)
	}
	fmt.Printf("Все заказы для %s: %v\n", userID, allOrdersResponse.GetOrderIds())

}
