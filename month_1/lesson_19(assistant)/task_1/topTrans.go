package task1

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 1.transactionlar soni bo'yicha top branches
func CalculateTranTopBranches() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	branches, _ := readBranches("data/branches.json")

	branchName := make(map[int]string)
	for _, b := range branches {
		branchName[b.ID] = b.Name
	}
	transCount := make(map[int]int)
	for _, t := range transactions {
		transCount[t.BranchID]++
	}

	for id, b := range branchName {
		fmt.Printf("%s - %d\n", b, transCount[id])
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
