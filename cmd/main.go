package main

import (
	"clone/lms_back/config"
	"clone/lms_back/controller"
	"clone/lms_back/storage"
	"fmt"

	"net/http"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/branche", con.Branche)
	 http.HandleFunc("/teacher", con.Teacher)
	 http.HandleFunc("/group", con.Group)

	fmt.Println("programm is running on localhost:8080...")
	http.ListenAndServe(":8080", nil)

}
