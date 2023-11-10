package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "first-service",
		"auto.offset.reset": "latest"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	topics := []string{"topic2"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println(fmt.Sprintf("Error while consuming message %v", msg), err)
			continue
		}
		fmt.Println("message:", string(msg.Key), string(msg.Value))
		time.Sleep(time.Second)
	}

}

// func main() {
// 	cfg := config.Load()
// 	lg := logger.NewLogger(cfg.Environment, "debug")
// 	strg, err := postgres.NewStorage(context.Background(), cfg)
// 	if err != nil {
// 		log.Fatalf("failed to connect to database: %v", err)
// 	}
// 	clients, err := grpc_client.New(cfg)
// 	if err != nil {
// 		log.Fatalf("failed to connect to services: %v", err)
// 	}
// 	s := grpc.SetUpServer(lg, strg, clients)
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	log.Printf("server listening at %v", lis.Addr())

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
