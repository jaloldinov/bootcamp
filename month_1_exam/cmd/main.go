package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	fmt.Println(con.Task_8())
	// con.OrderPayment(&models.OrderPayment{OrderId: "ff9aa3v6-7dd2-4b2e-9376-93bc47391e82"})
}
