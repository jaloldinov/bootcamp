/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	pb "example-grpc-server/genproto"

	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedBranchServiceServer
	pb.UnimplementedStreamServiceServer
	branches []*pb.Branch
}

// SayHello implements helloworld.GreeterServer

// func (s *server) Create(ctx context.Context, req *pb.CreateBranch) (*pb.CreateResponse, error) {
// 	log.Printf("Received: %v", req.GetName())
// 	id := uuid.NewString()
// 	s.branches = append(s.branches, &pb.Branch{
// 		Id:      id,
// 		Name:    req.Name,
// 		Address: req.Address,
// 	})
// 	return &pb.CreateResponse{Id: id}, nil
// }

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBranchServiceServer(s, &server{})
	pb.RegisterStreamServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Count(req *pb.Request, res pb.StreamService_CountServer) error {
	fmt.Println("request:", req.GetNumber())
	for i := 1; i < int(req.GetNumber()); i++ {
		err := res.Send(&pb.Response{Count: int32(i)})
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("res sent:", i)
		time.Sleep(time.Second)
	}
	return nil
}
func (s *server) Sum(stream pb.StreamService_SumServer) error {
	var total int32

	for {
		value, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("respond:", total)
			return stream.SendAndClose(&pb.Response{
				Count: total,
			})
		}

		if err != nil {
			return err
		}

		fmt.Println("received number:", value.GetNumber())
		total += value.GetNumber()
	}
}

func (s *server) Sqr(stream pb.StreamService_SqrServer) error {
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Println("received number:", value.GetNumber())

		if err := stream.Send(&pb.Response{
			Count: int32(math.Pow(float64(value.GetNumber()), 2)),
		}); err != nil {
			return err
		}
	}
}

func (s *server) Fibonacci(req *pb.Request, res pb.StreamService_FibonacciServer) error {
	fmt.Println("request:", req.GetNumber())

	a, b := 0, 1
	for i := 0; a < int(req.GetNumber()); i++ {

		if err := res.Send(&pb.Response{Count: int32(a)}); err != nil {
			return err
		}

		fmt.Println("sent: ", a)
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) Translate(stream pb.StreamService_TranslateServer) error {

	wordsMap := map[string]string{
		"yellow": "sariq",
		"blue":   "moviy",
		"red":    "qizil",
		"green":  "yashil",
		"white":  "oq",
		"black":  "qora",
		"brown":  "kulrang",
		"pink":   "pushti",
	}
	for _, word := range wordsMap {
		value, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("received word:", value.GetWord())

		if word == value.GetWord() {
			if err := stream.Send(&pb.RespondWords{
				Word: word,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
