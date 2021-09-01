package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "database/sql"
	// "github.com/joho/godotenv"
	// "os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.Static("css","src/css")
	engine.Static("js","src/js")
	engine.LoadHTMLGlob("src/*.html")

	engine.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{

		})
	})

	engine.GET("/signup",func(c *gin.Context){
		c.HTML(http.StatusOK,"signup.html",gin.H{})
	})

	engine.Run(":3000")
}
