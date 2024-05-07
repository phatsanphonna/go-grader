package routes

import (
	"fmt"
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

var restrictWord = []string{
	"__import", "importlib", "globals", "locals", "posix", "commands", "webbrowser", "spawnv", "popen", "subprocess", "system", "chdir", "makedirs", "removedirs", "renames", " os", ",os", "\nos", "\tos", " sys", ",sys", "\nsys", "\tsys", "setrecursionlimit", "exec", "open", "file", "eval", "MySQLdb", "socket", "multiprocessing", "builtins", "vars", "gets", "unistd.h", "exit",
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
	
	// Check for restricted words
	for _, word := range restrictWord {
		if strings.Contains(body.Code, word) {
			c.JSON(http.StatusOK, gin.H{
				"out":    "",
				"err":    fmt.Sprintf("Restricted word found, found %s", word),
				"status": 1,
			})
			return
		}
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
