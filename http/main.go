package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	response, err := http.Get("http://www.google.com")

	if err != nil {
		log.Fatal("Error:", err)
		os.Exit(1)
	}
	fmt.Println(*response)

	data := *response

	// reading header and printing it
	header := data.Header

	fmt.Println("\nPrinting header of the response")
	for key, value := range header {

		fmt.Println(key, ": ", value)
	}

	// reading body and printing it
	/*bs := make([]byte, 99999)
	body := data.Body

	count, err := body.Read(bs)

	fmt.Println(string(bs))
	fmt.Println(count)

	if err != nil {
		fmt.Println(err)
	}*/

	// second way to read the body and print it
	lw := logWriter{}
	fmt.Println("Printing using the second method")
	io.Copy(lw, data.Body)

}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("size of the byte size :", len(bs))
	return len(bs), nil
}
