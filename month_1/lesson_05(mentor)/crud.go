package main

import (
	"errors"
	"fmt"
)

func main() {

	Omadbek := Person{
		FirstName: "Omadbek",
		LastName:  "Jaloldinov",
		Job:       "Backend Developer",
		Age:       21,
	}

	Updater := Person{
		FirstName: "Mohira",
		LastName:  "Jaloldinova",
		Job:       "ELS teacher",
		Age:       18,
	}

	// creates three data
	for i := 0; i < 3; i++ {
		createPerson(Omadbek)
	}

	// gets all data
	fmt.Println(getAll())

	// gets data by id
	fmt.Println(getById(1))

	// printer(containerArr, containerMap)
	// updates data by given id
	fmt.Println(updatePerson(1, Updater))
	// printer(containerArr, containerMap)

	fmt.Println(deletePerson(2))
	// printer(containerArr, containerMap)

}

var containerMap = map[int]Person{}
var containerArr = []Person{}

var success = errors.New("success: ")
var err = errors.New("something went wrong: ")

type Person struct {
	FirstName string
	LastName  string
	Job       string
	Age       int
	Id        int
}

func createPerson(p Person) string {

	id := len(containerArr) + 1
	p.Id = id

	containerArr = append(containerArr, p)
	containerMap[id] = p

	return "created"
}

func getAll() (error, map[int]Person) {

	if len(containerMap) != 0 {
		return success, containerMap
	} else {
		return err, containerMap
	}
}

func getById(id int) (error, Person) {

	elem, ok := containerMap[id]
	if ok {
		return success, elem
	} else {
		return err, elem
	}
}

func updatePerson(id int, newP Person) string {

	_, ok := containerMap[id]
	if ok {
		newP.Id = id
		containerMap[id] = newP
		return "updated"
	} else {
		return "something went wrong"
	}
}

func deletePerson(id int) string {

	_, ok := containerMap[id]

	if ok {
		delete(containerMap, id)
		return "deleted"
	} else {
		return "someting went wrong"
	}

}

func printer(arr []Person, mapp map[int]Person) {

	for _, v := range arr {
		fmt.Println("arr => ", v)
	}
	for i, v := range mapp {
		fmt.Println("map => ", i, ":", v)
	}
}
