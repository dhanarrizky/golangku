package main

import (
	"go-web-native/config"
	"go-web-native/controller/categorycontroller"
	homecontroller "go-web-native/controller/homecontroller"
	"log"
	"net/http"

)

func main() {

	config.ConnectDB()
	// pangil home page
	http.HandleFunc("/", homecontroller.Welcome)

	//categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("server Ranning on port :8080")
	http.ListenAndServe(":8080", nil)
}
