package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id   int
	Name string
}

var DataSlice []User = []User{
	{
		Id:   1,
		Name: "Omadbek",
	},
	{
		Id:   2,
		Name: "Sarvarbek",
	},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/user/", handleUser)

	fmt.Println("server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodGet:
		// GetAllUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPut:
		//
	case http.MethodDelete:
		//
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while io.ReadAll:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		log.Println("error while Unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.Id = len(DataSlice) + 1
	DataSlice = append(DataSlice, newUser)

	responseData, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	// Convert DataSlice slice to JSON
	jsonData, err := json.Marshal(DataSlice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/user/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range DataSlice {
		if user.Id == id {
			jsonData, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			break
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while io.ReadAll:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("error while Unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for index, v := range DataSlice {
		if v.Id == id {
			DataSlice = append(DataSlice[:index], DataSlice[index+1:]...)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
