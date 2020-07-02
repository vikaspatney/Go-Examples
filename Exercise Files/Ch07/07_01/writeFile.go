package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	
	content := "Hello from Go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()
	// We use io & ioutil to write to a file; 

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("all done with the file %v characters", ln)

	bytes := []byte(content)
	// we convert the content to bytes then send it as a param also 0644 is used for file permission
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}