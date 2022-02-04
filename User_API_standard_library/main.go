package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerequest() {

	http.HandleFunc("/createusers", NewUser)
	http.HandleFunc("/getusers", GetUser)
	http.HandleFunc("/search/{id}", SearchUser)
	http.HandleFunc("/loginendpoint/{username,password}", LoginEndPoint)
	http.HandleFunc("/update/{id}", EditUser)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	fmt.Println("Go Task")
	handlerequest()

}
