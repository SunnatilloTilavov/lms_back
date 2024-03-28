package main

import (
	"fmt"
	"clone/lms_back/config"
	"clone/lms_back/api"
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

	c := api.New(store)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8008")
}
