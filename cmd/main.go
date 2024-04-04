package main

import (
	"fmt"
	"clone/lms_back/config"
	"clone/lms_back/api"
	"clone/lms_back/service"
	"clone/lms_back/storage/postgres"
	"context"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(),cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store)
	c := api.New(services,store)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8008")
}
