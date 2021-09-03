package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	"log"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)


func httpServer(){
	engine := gin.Default()
	engine.Static("css","src/css")
	engine.Static("js","src/js")
	engine.LoadHTMLGlob("src/*.html")

	engine.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{

		})
	})

	engine.GET("/signup",func(c *gin.Context){
		c.HTML(http.StatusOK,"signup.html",gin.H{
			//	サインアップする際はdbにアクセスする必要がある
			//	html側で{{.message}}で取得したものを表示可能
			"message" : mysqlQuery(),
		})
	})

	engine.Run(":3000")
}

func mysqlQuery() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db,err := sql.Open("mysql",os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 1
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = ?",id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name 
}

func main() {
	httpServer()
}
