package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("python", "./main.py")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}

	defer stdin.Close()

	var outbuf, errbuf strings.Builder

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	io.WriteString(stdin, "A\n")
	cmd.Wait()

	stdout := outbuf.String()
	stderr := errbuf.String()

	if len(stderr) != 0 {
		fmt.Println(stderr)
		return
	}

	fmt.Println(stdout)
}
