package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.nawakatha.com")
	if err != nil {
		fmt.Println("ERROR: ")
		os.Exit(1)
	}

	fmt.Println(resp)
}
