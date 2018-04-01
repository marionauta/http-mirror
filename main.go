package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var baseURL = "https://example.com"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("To change the default url: %s URL\n", os.Args[0])
	} else {
		baseURL = os.Args[1]
	}
	fmt.Println("Mirroring:", baseURL)
	fmt.Println("Ctrl+C to exit.")

	handler := http.HandlerFunc(mirror)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func mirror(w http.ResponseWriter, r *http.Request) {
	url := baseURL + r.URL.Path
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	copyHeaders(w.Header(), res.Header)
	io.Copy(w, res.Body)
}

func copyHeaders(to, from http.Header) {
	for k, vv := range from {
		for _, v := range vv {
			to.Add(k, v)
		}
	}
}
