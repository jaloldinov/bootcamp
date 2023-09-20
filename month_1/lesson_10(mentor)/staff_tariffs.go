package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	staf1 := Staff_tariff{
		Id:              1,
		Name:            "Premium",
		Type:            "fixed",
		amoun_for_cash:  12.2,
		amount_for_card: 12,
		Founded_at:      "2002-02-25",
	}

	staf2 := Staff_tariff{
		Id:              2,
		Name:            "Updated Premium",
		Type:            "percentage",
		amoun_for_cash:  1223.2,
		amount_for_card: 12323.1,
		Founded_at:      "2000-02-25",
	}

	// CREATE STAFF_TARIF
	res, err := staff_tariffs.CreateStaffTariff(staf1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
	fmt.Println(staff_tariffs)

	// GET TARIFF BY ID
	fmt.Println("GET BY ID:")
	tariff, err := staff_tariffs.GetTariffById(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(tariff)
	}

	// GET ALL DATA
	fmt.Println("GET ALL:")
	// 4)getAllda  name(search), type filter, pagination
	data, err := staff_tariffs.getAllTariffes("Pre", "fixed", 1, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(data)
	}

	// update staff
	result, err := staff_tariffs.UpdateStaffTariff(1, staf2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	fmt.Println(staff_tariffs)

	// DELETE STAFF
	result, err = staff_tariffs.DeleteTariff(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	fmt.Println(staff_tariffs)

}

type Staff_tariffs struct {
	Data []Staff_tariff
}

type Staff_tariff struct {
	Id              int
	Name            string
	Type            string
	amoun_for_cash  float64
	amount_for_card float64
	Founded_at      string
	Created_at      string
}

var staff_tariffs = Staff_tariffs{Data: make([]Staff_tariff, 0)}

func (s *Staff_tariffs) CreateStaffTariff(newTariff Staff_tariff) (string, error) {
	newTariff.Id = len(s.Data) + 1
	newTariff.Created_at = time.Now().Format("2006-01-02 15:04:05")
	for _, staff := range s.Data {
		if staff.Id == newTariff.Id {
			return "", fmt.Errorf("tariff with ID %d already exits", newTariff.Id)
		}
	}

	s.Data = append(s.Data, newTariff)
	return "created", nil
}

func (s *Staff_tariffs) GetTariffById(id int) (Staff_tariff, error) {
	for _, s := range s.Data {
		if s.Id == id {
			return s, nil
		}
	}
	return Staff_tariff{}, fmt.Errorf("no tariff found with ID %d", id)
}

func (s *Staff_tariffs) getAllTariffes(name, typeName string, limit, page int) ([]Staff_tariff, error) {
	filtered := []Staff_tariff{}
	checker := strings.Contains

	for i, v := range s.Data {
		if checker(v.Name, name) && checker(v.Type, typeName) {
			filtered = append(filtered, s.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit
	if len(filtered) < limit && page > 1 {
		return []Staff_tariff{}, fmt.Errorf("page not found")
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	return filtered[start:end], nil
}

func (s *Staff_tariffs) UpdateStaffTariff(id int, updatedTariff Staff_tariff) (string, error) {
	index := -1
	for i, staff := range s.Data {
		if staff.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return "", fmt.Errorf("no tariff found with ID %d", id)
	}
	// Check if the updated tariff ID already exists
	for i, staff := range s.Data {
		if staff.Id == updatedTariff.Id && i != id {
			return "", fmt.Errorf("tariff with ID %d already exits", updatedTariff.Id)
		}
	}

	updatedTariff.Id = id
	s.Data[index] = updatedTariff
	return "updated", nil
}

func (s *Staff_tariffs) DeleteTariff(id int) (string, error) {
	index := -1
	for i, tariff := range s.Data {
		if tariff.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return "", fmt.Errorf("no tariff found with ID %d", id)
	}

	s.Data = append(s.Data[:index], s.Data[index+1:]...)

	return fmt.Sprintf("tariff with ID %d deleted", id), nil
}
