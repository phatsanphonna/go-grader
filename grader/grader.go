package grader

import (
	"io"
	"os/exec"
	"strings"
	"time"
)

const (
	OK           int = 0
	ERROR        int = 1
	TIMEOUT      int = 2
	MEMORY_ERROR int = 3
)

func ExecutePythonCode(filename string, input string) (string, string, int) {
	cmd := exec.Command("python3", filename)
	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}

	defer stdin.Close()

	var outbuf, errbuf strings.Builder

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	var state int = OK

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	timer := time.AfterFunc(1*time.Second, func() {
		err := cmd.Process.Kill()
		if err != nil {
			panic(err)
		}
		state = TIMEOUT
	})

	io.WriteString(stdin, input)
	cmd.Wait()

	timer.Stop()

	stdout := outbuf.String()
	stderr := errbuf.String()

	if len(stderr) > 0 {
		state = ERROR
	}

	return stdout, stderr, state
}

func ExecuteJavaCode(filename string, input string) (string, string, int) {
	exec.Command("javac", filename+".java").Run()
	cmd := exec.Command("java", filename)
	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}

	defer stdin.Close()

	var outbuf, errbuf strings.Builder

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	var state int = OK

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	timer := time.AfterFunc(1*time.Second, func() {
		err := cmd.Process.Kill()
		if err != nil {
			panic(err)
		}
		state = TIMEOUT
	})

	io.WriteString(stdin, input)
	cmd.Wait()

	timer.Stop()

	stdout := outbuf.String()
	stderr := errbuf.String()

	if len(stderr) > 0 {
		state = ERROR
	}

	return stdout, stderr, state
}
