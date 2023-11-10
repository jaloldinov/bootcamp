package main

import (
	"context"
	"fmt"
	"log"
	"net"
	branches "server/genproto"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/fake"
	"google.golang.org/grpc"
)

type server struct {
	branches.UnimplementedBranchServiceServer
	branches []*branches.Branch
}

// inital data
func (s *server) CreateBranch() {
	branch := &branches.Branch{
		Id:      uuid.New().String(),
		Name:    fake.Brand(),
		Address: fake.StreetAddress(),
	}
	s.branches = append(s.branches, branch)
}

func main() {

	// adding inital data for 10 times
	fmt.Println("Start adding...")
	ser := &server{}
	for i := 1; i <= 10; i++ {
		ser.CreateBranch()
		fmt.Printf("created %d\n", i)
	}
	fmt.Println("finished")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	branches.RegisterBranchServiceServer(s, ser)

	log.Printf("server listening on: %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *branches.CreateBranchRequest) (*branches.CreateBranchResponse, error) {
	log.Printf("Received: %v", req)
	id := uuid.NewString()
	s.branches = append(s.branches, &branches.Branch{
		Id:      id,
		Name:    req.Name,
		Address: req.Address,
	})
	return &branches.CreateBranchResponse{}, nil
}

func (s *server) List(ctx context.Context, req *branches.ListBranchRequest) (*branches.ListBranchResponse, error) {
	log.Printf("Received: %v", req)

	page := req.Page
	limit := req.Limit
	searchQuery := req.Search

	// Filter branches based on the search
	filteredBranches := make([]*branches.Branch, 0)
	for _, branch := range s.branches {
		if searchQuery == "" || strings.Contains(strings.ToLower(branch.Name), strings.ToLower(searchQuery)) {
			filteredBranches = append(filteredBranches, branch)
		}
	}

	// Calculate the total number of filtered branches
	totalBranches := int32(len(filteredBranches))

	// Check if req.Limit is 0
	if limit == 0 {
		limit = int32(len(filteredBranches))
	}

	// Check if req.Page is 0
	if page == 0 {
		page = 1
	}

	startIndex := (page - 1) * limit
	endIndex := page * limit

	if startIndex >= int32(len(filteredBranches)) {
		return &branches.ListBranchResponse{
			Branches: nil,
			Count:    totalBranches,
		}, nil
	}

	if endIndex > int32(len(filteredBranches)) {
		endIndex = int32(len(filteredBranches))
	}

	// Get the branches within the specified range
	pagedBranches := filteredBranches[startIndex:endIndex]

	return &branches.ListBranchResponse{
		Branches: pagedBranches,
		Count:    totalBranches,
	}, nil
}

func (s *server) Get(ctx context.Context, req *branches.IdRequest) (*branches.GetBranchResponse, error) {
	log.Printf("Received id: %s", req.Id)

	branch := &branches.GetBranchResponse{}

	for _, b := range s.branches {
		if req.Id == b.Id {
			branch.Branch = b
			return branch, nil
		}
	}

	return nil, fmt.Errorf("branch not found")
}

func (s *server) Update(ctx context.Context, req *branches.UpdateBranchRequest) (*branches.Response, error) {
	log.Printf("Received: %v", req)

	for i, b := range s.branches {
		if req.Id == b.Id {
			s.branches[i].Name = req.Name
			s.branches[i].Address = req.Address
			return &branches.Response{Message: "updated"}, nil
		}
	}

	return nil, fmt.Errorf("branch not found with that id")
}

func (s *server) Delete(ctx context.Context, req *branches.IdRequest) (*branches.Response, error) {
	log.Printf("Received: %v", req)

	for i, b := range s.branches {
		if req.Id == b.Id {
			s.branches = append(s.branches[:i], s.branches[i+1:]...)
			return &branches.Response{Message: "deleted"}, nil
		}
	}
	return nil, fmt.Errorf("branch not found with that id")
}
