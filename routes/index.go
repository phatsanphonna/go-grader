package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phatsanphonna/go-grader/file"
	"github.com/phatsanphonna/go-grader/grader"
)

type GetIndexBody struct {
	Input string `json:"input"`
	Code  string `json:"code"`
}

func PostPythonCode(c *gin.Context) {
	var body GetIndexBody

	if err := c.BindJSON(&body); err != nil {
		log.Fatalln(err)
	}

	var input string = body.Input

	if !strings.HasPrefix(input, "\n") {
		input = input + "\n"
	}

	file.WriteFile("tmp.py", body.Code)
	out, err, status := grader.ExecutePythonCode("tmp.py", input)

	c.JSON(http.StatusOK, gin.H{
		"out":    out,
		"err":    err,
		"status": status,
	})
}

func PostJavaCode(c *gin.Context) {
	var body GetIndexBody

	if err := c.BindJSON(&body); err != nil {
		log.Fatalln(err)
	}

	var input string = body.Input

	if !strings.HasPrefix(input, "\n") {
		input = input + "\n"
	}

	file.WriteFile("Main.java", body.Code)
	out, err, status := grader.ExecuteJavaCode("Main", input)

	c.JSON(http.StatusOK, gin.H{
		"out":    out,
		"err":    err,
		"status": status,
	})
}
