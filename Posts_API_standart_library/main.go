package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerequest() {

	http.HandleFunc("/createposts", NewPost)
	http.HandleFunc("/getposts", GetPost)
	http.HandleFunc("/editpost/{id}", EditPost)
	http.HandleFunc("/deletepost/{id}", DeletePost)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Go Task")
	handlerequest()

}
