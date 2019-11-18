package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("token.txt")
	check(err)
	defer file.Close()

	txt, err := ioutil.ReadAll(file)
	check(err)
	token := txt

	fmt.Println(string(token))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
