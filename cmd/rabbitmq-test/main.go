package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Iniciando teste de conexão com RabbitMQ...")

	// URL de conexão
	url := "amqp://guest:guest@127.0.0.1:5672/"

	// Tentar conectar várias vezes
	var conn *amqp.Connection
	var err error
	for i := 0; i < 5; i++ {
		fmt.Printf("Tentativa %d de conexão...\n", i+1)
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		fmt.Printf("Erro na tentativa %d: %v\n", i+1, err)
		time.Sleep(time.Second * 2)
	}

	if err != nil {
		log.Fatalf("Não foi possível conectar ao RabbitMQ após 5 tentativas: %v", err)
	}
	defer conn.Close()

	fmt.Println("Conexão estabelecida com sucesso!")

	// Abrir canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir canal: %v", err)
	}
	defer ch.Close()

	// Declarar fila
	_, err = ch.QueueDeclare(
		"test_queue", // nome
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Erro ao declarar fila: %v", err)
	}

	fmt.Println("Fila declarada com sucesso!")

	// Publicar mensagem
	err = ch.Publish(
		"",           // exchange
		"test_queue", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Teste de mensagem"),
		},
	)
	if err != nil {
		log.Fatalf("Erro ao publicar mensagem: %v", err)
	}

	fmt.Println("Mensagem publicada com sucesso!")

	// Consumir mensagem
	msgs, err := ch.Consume(
		"test_queue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Erro ao consumir mensagem: %v", err)
	}

	// Aguardar mensagem
	msg := <-msgs
	fmt.Printf("Mensagem recebida: %s\n", msg.Body)

	fmt.Println("Teste concluído com sucesso!")
}
