package main

import (
	"fmt"
	"strings"
)

func main() {

	// =========== START   USER ==============
	Omadbek := User{
		Name:               "Omadbek",
		FavoriteProductIDs: []int{1, 2},
	}
	// UNCOMMENT THIS LINE BEFORE RUNNING UPDATE FUNC )
	// updater := User{
	// 	Name:               "Updated",
	// 	FavoriteProductIDs: []int{1, 2},
	// }

	for i := 0; i < 3; i++ {
		createUser(Omadbek)
	}
	// fmt.Println(getAll("M"))
	// fmt.Println(getById(2))
	// fmt.Println(updateUser(2, updater))
	// fmt.Println(getById(2))
	// fmt.Println(deleteUser(2))
	// fmt.Println(getById(2))
	// =========== END   USER ==============

}

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

var categoriesContainer = map[int]Category{
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

var usersContainer = map[int]User{
	1: {Id: 1, Name: "Omadbek", FavoriteProductIDs: []int{2, 3, 5, 6}},
	2: {Id: 2, Name: "Mohira", FavoriteProductIDs: []int{1, 2, 3, 5, 6}},
}

// ================= Category CRUD ==================

// ================= Product CRUD ==================
// ================= User CRUD ==================
func createUser(u User) string {

	id := len(usersContainer) + 1
	u.Id = id
	usersContainer[id] = u

	return "created"
}

func getAll(name string) (int, map[int]User) {

	var foundPersons = map[int]User{}
	count := 0

	for key, person := range usersContainer {
		fmt.Println(strings.Contains(person.Name, name))

		if strings.Contains(person.Name, name) {
			foundPersons[key] = person
			count++
		}
	}
	return count, foundPersons
}

func getById(id int) (string, User) {
	user, ok := usersContainer[id]

	if ok {
		return "found", user
	} else {
		return "not found", User{}
	}
}

func updateUser(id int, newUser User) string {

	_, ok := usersContainer[id]
	if ok {
		usersContainer[id] = newUser
		return "updated"
	} else {
		return "something wrong"
	}
}

func deleteUser(id int) string {

	_, ok := usersContainer[id]
	if ok {
		delete(usersContainer, id)
		return "deleted"
	} else {
		return "error"
	}
}
