package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	// the url below returns a json obj
	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// the above Get func returns a pointer of http response
	fmt.Printf("Response type: %T\n", resp) 
	defer resp.Body.Close()

	// here we can read the body/content that shall be returned in bytes
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// we convert it to string to make it readable
	content := string(bytes)
	fmt.Println(" here is the content: ", content)
}
