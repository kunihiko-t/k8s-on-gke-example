package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var html []byte

type User struct {
	ID   int    `json:"id" gorm:"primary_key" `
	Name string `json:"name" sql:"size:200"`
}

func init() {
	db = GetDB()
	db.LogMode(true)
	html, _ = ioutil.ReadFile("index.html")
}

func CloseDB() {
	_ = db.Close()
}

func GetDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("postgres", "user=postgres dbname=example sslmode=disable host=db")
	if err != nil {
		fmt.Println("got an unexpected error during connecting db.")
		panic(err)
	} else {
		fmt.Println("Success: DB Access", db)
	}
	return
}

func main() {
	defer CloseDB()

	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {

		if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			u := User{}
			decoder.Decode(&u)
			fmt.Printf("Name: %v \n", u.Name)
			db.Create(&u)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		users := []User{}
		db.Find(&users)
		b, _ := json.Marshal(users)
		fmt.Fprint(w, string(b))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, string(html))
	})

	http.ListenAndServe(":9090", nil)
}
