package main

import (
	"fmt"
	"strings"
)

func main() {
	// gets all data with filter based on given age and name
	fmt.Println(getAll(21, "Jaloldinov"))

}

var containerMap = map[int]Person{
	1: {
		Id:        1,
		FirstName: "Sarvarbek",
		LastName:  "Fozilov",
		Job:       "Backend Developer",
		Age:       98,
	},
	2: {
		Id:        2,
		FirstName: "Sarvarbek",
		LastName:  "Usmonov",
		Job:       "Backend Developer",
		Age:       25,
	},
	3: {
		Id:        3,
		FirstName: "Guli",
		LastName:  "Nishonboyeva",
		Job:       "Designer",
		Age:       16,
	},
	4: {
		Id:        4,
		FirstName: "Mashxurbek",
		LastName:  "Nishonboyev",
		Job:       "Pupil",
		Age:       14,
	},
	5: {
		Id:        5,
		FirstName: "Mohira",
		LastName:  "Jaloldinova",
		Job:       "ELS Teacher",
		Age:       18,
	},
	6: {
		Id:        6,
		FirstName: "Omadbek",
		LastName:  "Jaloldinov",
		Job:       "Backend Developer",
		Age:       21,
	},
	7: {
		Id:        7,
		FirstName: "Omadbek",
		LastName:  "Jaloldinov",
		Job:       "Backend Developer",
		Age:       21,
	},
}

type Person struct {
	FirstName string
	LastName  string
	Job       string
	Age       int
	Id        int
}

func getAll(age int, name string) (int, map[int]Person) {

	var foundPersons = map[int]Person{}
	count := 0

	for key, person := range containerMap {
		if person.Age == age && (strings.Contains(person.FirstName, name) || strings.Contains(person.LastName, name)) {
			foundPersons[key] = person
			count++
		}
	}

	return count, foundPersons
}
