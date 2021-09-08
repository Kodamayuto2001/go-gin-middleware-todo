package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
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
			//	本番では
			//	uri:	api/signup (POST)
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

type UserField struct {
	// ID			int		`json:"id"`
	Name 		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

func userAdd(user *UserField) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db,err := sql.Open("mysql",os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name 		:= user.Name
	email 		:= user.Email
	password 	:= user.Password

	ins,err := db.Prepare("INSERT INTO members(name,email,password) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	ins.Exec(name,email,password)

	message := fmt.Sprintf("\nname: %v\nemail: %v\npassword: %v\n",name,email,password)
	return message
}

func main() {
	// httpServer()

	//	api仕様
	//	/api/v1/users
		//	/
		//	/add
		//	/update
		//	/delete

	engine := gin.Default()

	engine.POST("/api/v1/users/add",func(c *gin.Context){
		var user UserField
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"name"		:user.Name,
				"email"		:user.Email,
				"password"	:user.Password,
			},
		)
		
		//	クエリの実行（insert）
		message := userAdd(&user)
		c.String(http.StatusOK,message)
	})

	engine.Run(":3000")
}
