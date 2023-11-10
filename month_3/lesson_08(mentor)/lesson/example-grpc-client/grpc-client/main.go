package main

import (
	"context"
	sale_service "example-grpc-client/genproto"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	streamService := sale_service.NewStreamServiceClient(conn)

	// Call the desired streaming function
	// Uncomment the function you want to execute

	// serverSideStreaming(streamService)
	// clientSideStreaming(streamService)
	// bidirectionalStreaming(streamService)
	// fibonacciStreaming(streamService)
	translateStreaming(streamService)
}

// Function for server-side streaming
func serverSideStreaming(client sale_service.StreamServiceClient) {
	resStream, err := client.Count(context.Background(), &sale_service.Request{Number: 12})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		resp, err := resStream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("resp received: ", resp.GetCount())
	}
}

// Function for client-side streaming
func clientSideStreaming(client sale_service.StreamServiceClient) {
	stream, err := client.Sum(context.Background())
	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&sale_service.Request{Number: int32(i)})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		fmt.Println("sent:", i)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Closing", err)
	}
	fmt.Println("Total sum", res.GetCount())
}

// Function for bidirectional streaming
func bidirectionalStreaming(client sale_service.StreamServiceClient) {
	stream, err := client.Sqr(context.Background())
	if err != nil {
		log.Fatalln("Opening stream", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&sale_service.Request{Number: int32(i)})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		fmt.Println("send:", i)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Closing", err)
		}
		fmt.Println("Received:", res.GetCount())
		time.Sleep(time.Second)
	}
}

// Function for Fibonacci streaming
func fibonacciStreaming(client sale_service.StreamServiceClient) {

	num := 20
	stream, err := client.Fibonacci(context.Background(), &sale_service.Request{Number: int32(num)})
	if err != nil {
		log.Fatalln("Opening stream", err)
	}
	fmt.Println("sent:", num)

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Closing", err)
		}
		fmt.Println("Received:", res.GetCount())
		time.Sleep(time.Second)
	}
}

// Function for bidirectional streaming
func translateStreaming(client sale_service.StreamServiceClient) {
	words := []string{"yellow", "red", "green", "yellow", "green", "blue"}
	stream, err := client.Translate(context.Background())
	if err != nil {
		log.Fatalln("Opening stream", err)
	}

	for _, word := range words {
		err := stream.Send(&sale_service.RequestWords{Word: word})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		fmt.Println("sent:", word)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Closing", err)
		}
		fmt.Println("Received:", res.GetWord())
		time.Sleep(time.Second)
	}
}
