package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var branches = Branches{Data: make([]Branch, 0)}

func main() {
	data1 := Branch{
		ID:        1,
		Name:      "Omadbek",
		Adress:    "Andijan City",
		FoundedAt: "2002-02-25",
	}
	data2 := Branch{
		ID:        2,
		Name:      "Sarvarbek",
		Adress:    "Toshkent City",
		FoundedAt: "2002-02-25",
	}

	// CREATE BRANCH
	_, err := branches.CreateBranch(data1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successfully created")
	}

	// UPDATE BRANCH
	res, err := branches.UpdateBranch(1, data2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	// GET BRANCH WITH ID
	branch, err := branches.GetBranch(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(branch)
	}

	// GET ALL BRANCHES
	allData, err := branches.GetAllBranches("Sarvarbek", "City", 1, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(allData)
	}

	// DELETE BRANCH
	r, err := branches.DeleteBranch(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}

}

type Branch struct {
	ID        int
	Name      string
	Adress    string
	FoundedAt string
	Year      int
	CreatedAt string
}

type Branches struct {
	Data []Branch
}

// creates new branch
func (b *Branches) CreateBranch(newBranch Branch) (string, error) {
	newBranch.ID = len(b.Data) + 1
	newBranch.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	// Check if branch id already exists
	for _, branch := range b.Data {
		if branch.ID == newBranch.ID {
			return "", fmt.Errorf("branch with ID  %d already exists", newBranch.ID)
		}
	}
	// Add new branch slice
	b.Data = append(b.Data, newBranch)
	return "successfully created", nil
}

// update branch
func (b *Branches) UpdateBranch(branchID int, updatedBranch Branch) (string, error) {
	index := -1
	for i, branch := range b.Data {
		if branch.ID == branchID {
			index = i
			break
		}
	}
	if index == -1 {
		return "", fmt.Errorf("no branch found with ID %d", branchID)
	}

	// Check if the updated branch ID already exists
	for i, branch := range b.Data {
		if branch.ID == updatedBranch.ID && i != index {
			return "", fmt.Errorf("branch with ID %d already exists", updatedBranch.ID)
		}
	}

	updatedBranch.ID = branchID
	b.Data[index].FoundedAt = updatedBranch.FoundedAt
	b.Data[index] = updatedBranch
	return "successfully updated", nil
}

func (b *Branches) GetBranch(id int) (Branch, error) {
	for i := range b.Data {
		if b.Data[i].ID == id {
			foundedAt, _ := strconv.Atoi(b.Data[i].FoundedAt[:4])
			b.Data[i].Year = (time.Now().Year() - foundedAt)
			return b.Data[i], nil
		}
	}
	return Branch{}, fmt.Errorf("no branch found with ID %d", id)
}

// GET ALL BRANCHES WITH FILRER AND PAGINATION
func (b *Branches) GetAllBranches(name, adress string, limit, page int) ([]Branch, error) {

	searched := []Branch{}

	for _, branch := range branches.Data {
		if strings.Contains(branch.Name, name) && strings.Contains(branch.Adress, adress) {
			searched = append(searched, branch)
		}
	}

	start := (page - 1) * limit
	end := start + limit
	if end > len(searched) {
		end = len(searched)
	}
	return searched[start:end], nil
}

// DELETE BRANCH BASES ON GIVEN ID
func (b *Branches) DeleteBranch(id int) (string, error) {
	index := -1
	for i, branch := range b.Data {
		if branch.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return "", fmt.Errorf("no branch found with ID %d", id)
	}

	b.Data = append(b.Data[:index], b.Data[index+1:]...)

	return fmt.Sprintf("branch with ID %d deleted", id), nil
}
