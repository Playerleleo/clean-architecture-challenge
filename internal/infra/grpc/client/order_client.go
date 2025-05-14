package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	client pb.OrderServiceClient
}

func NewOrderClient(address string) (*OrderClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao servidor gRPC: %v", err)
	}

	client := pb.NewOrderServiceClient(conn)
	return &OrderClient{
		client: client,
	}, nil
}

func (c *OrderClient) CreateOrder(id string, price, tax float32) (*pb.CreateOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CreateOrderRequest{
		Id:    id,
		Price: price,
		Tax:   tax,
	}

	resp, err := c.client.CreateOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar order: %v", err)
	}

	return resp, nil
}

func (c *OrderClient) ListOrders() (*pb.ListOrdersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListOrdersRequest{}
	resp, err := c.client.ListOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar orders: %v", err)
	}

	return resp, nil
}

func RunClient() {
	grpcServer := os.Getenv("GRPC_SERVER")
	if grpcServer == "" {
		grpcServer = "localhost:50051"
	}

	conn, err := grpc.Dial(grpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// Criar primeiro pedido com ID único baseado no timestamp
	createOrder1, err := client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Id:    fmt.Sprintf("order_%d", time.Now().UnixNano()),
		Price: 100.0,
		Tax:   10.0,
	})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}
	fmt.Printf("Order created: %v\n", createOrder1)

	// Criar segundo pedido com ID único baseado no timestamp
	createOrder2, err := client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Id:    fmt.Sprintf("order_%d", time.Now().UnixNano()),
		Price: 200.0,
		Tax:   20.0,
	})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}
	fmt.Printf("Order created: %v\n", createOrder2)

	// Listar pedidos
	time.Sleep(time.Second) // Aguardar um pouco para garantir que os pedidos foram salvos
	listOrders, err := client.ListOrders(context.Background(), &pb.ListOrdersRequest{})
	if err != nil {
		log.Fatalf("could not list orders: %v", err)
	}
	fmt.Printf("Orders: %v\n", listOrders)
}
