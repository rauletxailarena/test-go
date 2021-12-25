package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(resp.Body)
	// io.Copy(os.Stdout, resp.Body)
}
