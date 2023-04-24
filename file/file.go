package file

import "os"

func WriteFile(filename string, code string) int {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	totalBytes, err := f.WriteString(code)

	if err != nil {
		panic(err)
	}

	return totalBytes
}
