package main

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "net/http"
    //"fmt"
    _"github.com/go-sql-driver/mysql"
    "log"
    //"os"
    //"io"
    //"io/ioutil"
    "ginpro/logging"
)

// var db *sql.DB
// var err error 

func main() {
    logging.LoggingSettings("log.log")
	log.Println("log start.")

    db, err := sql.Open("mysql", "root:Overtry09@tcp(localhost:3306)/mydb")
    if err != nil {
		log.Fatal("OpenError: ", err)
	}
    engine:= gin.Default()
    // htmlのディレクトリを指定
    engine.LoadHTMLGlob("templates/*")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    })
	engine.POST("/hello", func(c *gin.Context) {
        if c.PostForm("logintry") == "logintry"{
            // MYSQL Connect
            var result string
            var (
                id   *string
                password *string
            )
            _ = password
            rows, err := db.Query("select * from user")
            if err != nil{
                log.Fatal(err)
            }
            Flg := false
            inputId := c.PostForm("username")
            inputPass := c.PostForm("password")

            for rows.Next() {
                if err := rows.Scan(&id, &password); err != nil {
                    log.Fatalf("failed to scan row: %s", err)
                }
                log.Printf("id: %d\n", *id)
                log.Printf("pass: %d\n", *password)
                if *id == inputId && *password == inputPass{
                    Flg = true
                }
            }
            if err = rows.Err(); err != nil {
                log.Fatalln(err)
            }

            if Flg == true{
                result = "OK"
            } else {
                result = "NG"
            }
            log.Println(result, "hello parameter")
            c.HTML(http.StatusOK, "index.html", gin.H{
                // htmlに渡す変数を定義
                "message": result,
            })
        } else{
            c.HTML(http.StatusOK, "index.html", gin.H{
                // htmlに渡す変数を定義
                "message": "user add",
            })
            // insert
            ins, err := db.Prepare("INSERT INTO user(id, password) VALUES(?,?)")
            if err != nil {
                log.Fatal(err)
            }
            ins.Exec(c.PostForm("username"), c.PostForm("password"))
        }
    })

    engine.Run(":3000")
}

