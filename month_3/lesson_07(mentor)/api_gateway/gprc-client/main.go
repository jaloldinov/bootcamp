package main

import (
	branches "api_gateway/genproto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := branches.NewBranchServiceClient(conn)

	// createBranch(c)
	listBranches(c)
	getBranchByID(c)

	updateBranch(c)
	// check if branch update is successful
	getBranchByID(c)

	deleteBranch(c)
	// check if branch delete is successful
	getBranchByID(c)

}

func createBranch(c branches.BranchServiceClient) {
	name := ""
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	address := ""
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	r, err := c.Create(context.Background(), &branches.CreateBranchRequest{
		Name:    name,
		Address: address,
	})
	if err != nil {
		log.Fatalf("Could not create branch: %v", err)
	}

	log.Printf("New branch: %v", r)
}

func listBranches(c branches.BranchServiceClient) {
	limit := 0
	fmt.Print("Enter limit: ")
	fmt.Scan(&limit)
	page := 0
	fmt.Print("Enter page: ")
	fmt.Scan(&page)
	search := ""
	fmt.Print("Enter search: ")
	fmt.Scan(&search)

	r, err := c.List(context.Background(), &branches.ListRequest{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: search,
	})
	if err != nil {
		log.Fatalf("Could not list branches: %v", err)
	}

	log.Printf("List of branches: %v", r)
}

func getBranchByID(c branches.BranchServiceClient) {
	id := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&id)

	r, err := c.Get(context.Background(), &branches.IdRequest{Id: id})
	if err != nil {
		log.Fatalf("Could not get branch: %v", err)
	}

	log.Printf("Branch details: %v", r)
}

func updateBranch(c branches.BranchServiceClient) {
	id := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&id)

	name := ""
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	address := ""
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	r, err := c.Update(context.Background(), &branches.UpdateBranchRequest{
		Id:      id,
		Name:    name,
		Address: address,
	})
	if err != nil {
		log.Fatalf("Could not update branch: %v", err)
	}

	log.Printf("Updated branch: %v", r)
}

func deleteBranch(c branches.BranchServiceClient) {
	id := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&id)

	r, err := c.Delete(context.Background(), &branches.IdRequest{Id: id})
	if err != nil {
		log.Fatalf("Could not delete branch: %v", err)
	}
	log.Printf("Deleted branch: %v", r)
}
