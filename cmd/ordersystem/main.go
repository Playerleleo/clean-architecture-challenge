package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/configs"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/event"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/database"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/graph"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/grpc/client"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/grpc/pb"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/grpc/service"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/web/handlers"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/web/webserver"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/usecase"
	"github.com/leonardogomesdossantos/clean-architecture-challenge/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()

	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := handlers.NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("POST", "/order", webOrderHandler.Create)
	webserver.AddHandler("GET", "/order", webOrderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(createOrderUseCase, listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// Teste do gRPC
	time.Sleep(time.Second) // Aguardar o servidor iniciar
	orderClient, err := client.NewOrderClient(fmt.Sprintf("localhost:%s", configs.GRPCServerPort))
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

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	go http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

	// Manter o programa rodando
	select {}
}
