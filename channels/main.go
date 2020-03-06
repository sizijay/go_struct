package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("ERROR: ", err)
		c <- link
	} else {
		fmt.Println("The link", link, " seems to be up and running")
		//c <- "The website seems to be up"
		c <- link
	}
}

func main() {
	links := []string{
		"http://www.google.com",
		"http://nawakatha.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://dsfjsdkfjsdlfjkdfsljfksdfjlsdjfkgldjsdg.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for range links {
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		//go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}
