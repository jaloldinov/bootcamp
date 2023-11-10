package task1

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 1.transactionlar soni bo'yicha top branches
func CalculateTranTopBranches() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	branches, _ := readBranches("data/branches.json")

	branchCounts := make(map[int]int)
	for _, transaction := range transactions {
		branchCounts[transaction.BranchID]++
	}

	var sortedBranches []models.BranchTransactionCount
	for branchID, count := range branchCounts {
		for _, b := range branches {
			if branchID == b.ID {
				sortedBranches = append(sortedBranches, models.BranchTransactionCount{BranchID: branchID, BranchName: b.Name, Count: count})
			}
		}
	}

	sort.Slice(sortedBranches, func(i, j int) bool {
		return sortedBranches[i].Count > sortedBranches[j].Count
	})

	for _, v := range sortedBranches {
		fmt.Printf("Branch: %s: Total Transactions: %d\n", v.BranchName, v.Count)
	}
}

// ================================READERS======================================

func readTransaction(data string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &transactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return transactions, nil
}

func readBranches(data string) ([]models.Branch, error) {
	var branches []models.Branch

	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read branch: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal branch: %+v", err)
		return nil, err
	}

	return branches, nil
}
