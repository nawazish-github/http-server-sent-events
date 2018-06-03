package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(resp http.ResponseWriter, req *http.Request) {

	counter := 0

	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "text/event-stream")
	resp.Header().Set("Connection", "Keep-Alive")
	resp.Header().Set("Cache-Control", "No-Cache")
	resp.Write([]byte("\r\n"))

	for {
		if _, err := resp.Write([]byte(strconv.Itoa(counter))); err != nil {
			fmt.Println("client went off...")
			break
		}
		resp.Write([]byte("\r\n"))
		counter++
		resp.Write([]byte(strconv.Itoa(counter)))
		resp.Write([]byte("\r\n"))
		counter++
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("handler exiting...")
}
