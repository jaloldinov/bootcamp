package main

import (
	"fmt"
)

type Category struct {
	Id       int
	Name     string
	Products map[int]Product
}

type Product struct {
	Id    int
	Name  string
	Price float64
}

type User struct {
	Id                 int
	Name               string
	FavoriteProductIDs []int
}

var categories = map[int]Category{
	1: {
		Id:   1,
		Name: "Fruits",
		Products: map[int]Product{
			1: {Id: 1, Name: "Apple", Price: 1.0},
			2: {Id: 2, Name: "Banana", Price: 0.5},
			3: {Id: 3, Name: "Orange", Price: 1.2},
			4: {Id: 4, Name: "Melon", Price: 1.2},
		},
	},

	2: {
		Id:   2,
		Name: "Technologies",
		Products: map[int]Product{
			5: {Id: 1, Name: "Laptop", Price: 1500.0},
			6: {Id: 2, Name: "Smartphone", Price: 1000.0},
			7: {Id: 3, Name: "Earphone", Price: 1000.0},
			8: {Id: 3, Name: "Charger", Price: 1000.0},
		},
	},
}

var users = []User{
	{Id: 1, Name: "Omadbek", FavoriteProductIDs: []int{2, 3, 5, 6}},
	{Id: 2, Name: "Mohira", FavoriteProductIDs: []int{1, 2, 3, 5, 6}},
}

func main() {

	for _, user := range users {

		fruitCount := 0
		techCount := 0

		for _, favoriteID := range user.FavoriteProductIDs {
			for _, category := range categories {
				if _, ok := category.Products[favoriteID]; ok {

					if category.Name == "Fruits" {
						fruitCount++
					} else if category.Name == "Technologies" {
						techCount++
					}
				}
			}
		}

		fmt.Printf("%s: Fruits: %d, Technologies: %d\n", user.Name, fruitCount, techCount)
	}
}
