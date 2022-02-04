package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

//DB ...
var DB *sql.DB

type posts struct {
	UniqueID  string `gorm:"primaryKey"`
	AuthorID  string `json:"author_id"`
	PostedOn  string `json:"posted_on"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Thumbnail string `json:"thumbnail"`
}

//NewPost function ..
func NewPost(w http.ResponseWriter, r *http.Request) {
	connect()
	var post posts
	sqlStatement := `
		INSERT INTO posts (unique_id,author_id,posted_on,title,body,thumbnail)
	VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := DB.Exec(sqlStatement, post.UniqueID, post.AuthorID, post.PostedOn, post.Title, post.Body, post.Thumbnail)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("post added")
}

//GetPost function..
func GetPost(w http.ResponseWriter, r *http.Request) {
	connect()
	sqlStatement := `select * from posts order by posted_on ASC`
	result, _ := DB.Exec(sqlStatement)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//EditPost function..
func EditPost(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := r.URL.Query()
	postID := vars.Get("id")
	var getpost posts
	sqlStatement := `select * from posts where unique_id=$1`
	_, err := DB.Exec(sqlStatement, postID)
	if err != nil {
		fmt.Println("Post not found")
	}
	json.NewDecoder(r.Body).Decode(&getpost)
	DB.Query("update posts set author_id= ?, posted_on=?,title=?,body=?,thumbnail=?", getpost.AuthorID, getpost.PostedOn, getpost.Title, getpost.Body, getpost.Thumbnail)
	finalsqlStatement := `select * from posts where unique_id=$1`
	result, err := DB.Exec(finalsqlStatement, postID)
	if err != nil {
		fmt.Println("Post not found")
	}
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

// //DeletePost function..
func DeletePost(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := r.URL.Query()
	postID := vars.Get("id")
	sqlStatement := `delete from posts where unique_id= $1`
	_, err := DB.Exec(sqlStatement, postID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("Post Deleted")
}

// DB Connection funcion..
func connect() {
	dsn := "host=localhost port=5432 user=postgres password=yuvi@123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Can not connect to the database")
	}
	DB = db

}
