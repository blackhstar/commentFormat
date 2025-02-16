package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)

	b_string := string(b[:])

	text := strings.Split(b_string, "\n")

	comment := "/*"

	max_size := 10

	// fmt.Println("max_size  ", max_size)

	for _, value := range text {

		if len(value) > max_size {
			max_size = len(value)
		}

	}

	for i := 0; i < max_size+8; i++ {
		comment += "*"
	}
	comment += "*\n*"
	for i := 1; i < max_size+10; i++ {
		comment += " "
	}
	comment += "*\n"

	for _, value := range text {

		comment += "*    " + value
		for i := len(value) + 5; i < max_size+10; i++ {
			comment += " "
		}
		comment += "*\n"
	}

	comment += "*"
	for i := 0; i < max_size+9; i++ {
		comment += " "
	}
	comment += "*\n"

	for i := 0; i < max_size+9; i++ {
		comment += "*"
	}

	comment += "*/"

	// fmt.Printf("%s\n", b)

	newFile, err := os.Create("comment.txt")

	err = os.WriteFile("comment.txt", []byte(comment), 0666)

	newFile.Close()
}
