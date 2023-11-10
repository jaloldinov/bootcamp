package main

import (
	"fmt"
	"lesson_20/config"
	"lesson_20/handler"
	"lesson_20/storage/memory"
	"net/http"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage("data/branch.json", "data/staff.json", "data/sale.json", "data/transaction.json", "data/tariff.json")
	handler := handler.NewHandler(strg, *cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/branch/", handler.BranchHandler)

	fmt.Println("server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// func getAllPerson(rw http.ResponseWriter, req *http.Request) {
	// 	if req.Method == http.MethodGet {

	// 		values := req.URL.Query()
	// 		search := values.Get("search")
	// 		limit := values.Get("limit")
	// 		limitN, _ := strconv.Atoi(limit)
	// 		fmt.Println(search)
	// 		fmt.Println(limitN)
	// 		data, err := json.Marshal(persons)
	// 		if err != nil {
	// 			fmt.Println("error:", err.Error())
	// 			return
	// 		}
	// 		rw.Header().Add("content-type", "application/json")
	// 		rw.Write(data)
	// 		rw.WriteHeader(http.StatusOK)
	// 	}

}
