package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("This much of bytes were written: %v\n", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://www.nawakatha.com")
	if err != nil {
		fmt.Println("ERROR: ")
		os.Exit(1)
	}

	//fmt.Println(resp)
	//bs := make([]byte, 99999)
	//i, _ := resp.Body.Read(bs)
	//fmt.Println(bs[:i])

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
