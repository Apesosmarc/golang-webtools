package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	cl := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:3000", nil)
	if err != nil {
		os.Exit(1)
	}

	reqT := flag.String("method", "GET", "make a request to hello world server")
	flag.Parse()
	fmt.Printf("request type: %v\n", *reqT)

	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("RESPONSE:")
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(string(r))

}
