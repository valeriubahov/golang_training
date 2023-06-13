package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	// fetch the URL
	resp, err := http.Get("http://google.com")

	//the response does not have a visible BODY attribute

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}

	// we retrieve the body content by using the Reader inteface which resp.Body implements
	io.Copy(lw, resp.Body)

	// Copy needs a Writer and Reader interface as parameters
	// os.Stdout use a type File to write - File implements a Writer interface therefore, it use Writer interface
	// resp.body implements Reader interface

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
