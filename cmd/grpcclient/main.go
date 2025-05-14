package main

import (
	"fmt"
	"log"
	"time"

	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/grpc/client"
)

func main() {
	// Criar cliente
	orderClient, err := client.NewOrderClient("localhost:50051")
	if err != nil {
		log.Fatalf("Erro ao criar cliente: %v", err)
	}

	// Criar primeira order
	orderID1 := fmt.Sprintf("order_%d", time.Now().UnixNano())
	createResp1, err := orderClient.CreateOrder(orderID1, 100.0, 10.0)
	if err != nil {
		log.Fatalf("Erro ao criar order: %v", err)
	}
	fmt.Printf("Order criada: %v\n", createResp1)

	// Criar segunda order
	orderID2 := fmt.Sprintf("order_%d", time.Now().UnixNano())
	createResp2, err := orderClient.CreateOrder(orderID2, 200.0, 20.0)
	if err != nil {
		log.Fatalf("Erro ao criar order: %v", err)
	}
	fmt.Printf("Order criada: %v\n", createResp2)

	// Listar orders
	time.Sleep(time.Second) // Aguardar um pouco para garantir que as orders foram salvas
	listResp, err := orderClient.ListOrders()
	if err != nil {
		log.Fatalf("Erro ao listar orders: %v", err)
	}
	fmt.Printf("Orders: %v\n", listResp)
}
