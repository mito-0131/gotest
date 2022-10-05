package main

import (
    "fmt"
    "log"

	"database/sql"
    _"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Overtry09@tcp(localhost:3306)/mydb")
    fmt.Println("うううううう",db)
    if err != nil{
        log.Fatal("db error.")
    }
	defer db.Close()
	
	// insert
    ins, err := db.Prepare("INSERT INTO user(id, password) VALUES(?,?)")
    if err != nil {
        log.Fatal(err)
    }
    ins.Exec("Koji", "password123")
}