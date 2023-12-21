package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phatsanphonna/go-grader/routes"
)

var HOST string = "localhost:6000"

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/python", routes.PostPythonCode)
	r.POST("/java", routes.PostJavaCode)
	r.RunTLS(":6001", "/fullchain.pem", "/privkey.pem")
	fmt.Println("Server is running on 6001")
}
