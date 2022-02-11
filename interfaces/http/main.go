package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes : ", len(p))
	return len(p), nil
}
