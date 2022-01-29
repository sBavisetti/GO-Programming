package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {

	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal("An Error occured when opening a file : ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, f)

}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println("Printing the content of the file : ")
	fmt.Println(string(bs))
	fmt.Println("File Size : ", len(bs))
	return len(bs), nil

}
