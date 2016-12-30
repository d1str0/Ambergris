package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrMustNotBeDirectory = errors.New("argument cannot be a directory")
var ErrNot4Bytes = errors.New("did not read 4 bytes from file")

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Filename: %s\n", file.Name())
	fs, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\tName: %s\n", fs.Name())
	fmt.Printf("\tSize: %d\n", fs.Size())
	fmt.Printf("\tMode: %x\n", fs.Mode())
	fmt.Printf("\tMod Time: %s\n", fs.ModTime().String())
	fmt.Printf("\tIs Dir: %b\n", fs.IsDir())

	if fs.IsDir() {
		log.Fatal(ErrMustNotBeDirectory)
	}

	b := make([]byte, 4)
	n, err := file.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	if n != 4 {
		log.Fatal(ErrNot4Bytes)
	}
	fmt.Printf("First 4 bytes: %x\n", b)
}
