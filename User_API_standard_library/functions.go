package main

import (
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/lithammer/shortuuid"
)

//DB ...
var DB *sql.DB

type users struct {
	UniqueID    string `gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber int    `json:"phone_number"`
}

//NewUser function ..
func NewUser(w http.ResponseWriter, r *http.Request) {
	connect()
	var user users
	json.NewDecoder(r.Body).Decode(&user)
	id := shortuuid.New()
	user.UniqueID = id
	encodedPassword := b64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Password = encodedPassword

	sqlStatement := `
		INSERT INTO users (unique_id,name,email,username,password,date_of_birth,phone_number)
	VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := DB.Exec(sqlStatement, user.UniqueID, user.Name, user.Email, user.Username, user.Password, user.DateOfBirth, user.PhoneNumber)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("user added")
}

//GetUser function..
func GetUser(w http.ResponseWriter, r *http.Request) {
	connect()
	sqlStatement := `select * from users`
	result, _ := DB.Exec(sqlStatement)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//SearchUser function..
func SearchUser(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := r.URL.Query()
	userID := vars.Get("id")
	sqlStatement := `select * from users where unique_id=$1`
	result, _ := DB.Exec(sqlStatement, userID)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//LoginEndPoint function..
func LoginEndPoint(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := r.URL.Query()
	usrname := vars.Get("username")
	pasword := vars.Get("password")
	sqlStatement := `select * from users where username=$1 and password=$2`
	result, _ := DB.Exec(sqlStatement, usrname, pasword)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//EditUser function..
func EditUser(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := r.URL.Query()
	userID := vars.Get("id")
	var getUser users
	sqlStatement := `select * from users where unique_id=$1`
	_, err := DB.Exec(sqlStatement, userID)
	if err != nil {
		fmt.Println("Users not found")
	}
	json.NewDecoder(r.Body).Decode(&getUser)
	DB.Query("update users set name= ?, email=?,username=?,password=?,date_of_birth=?,phone_number=?", getUser.Name, getUser.Email, getUser.Username, getUser.Password, getUser.DateOfBirth, getUser.PhoneNumber)
	finalsqlStatement := `select * from users where unique_id=$1`
	result, err := DB.Exec(finalsqlStatement, userID)
	if err != nil {
		fmt.Println("User not found")
	}
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//DB Connection funcion..
func connect() {
	dsn := "host=localhost port=5432 user=postgres password=yuvi@123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Can not connect to the database")
	}
	DB = db
}
