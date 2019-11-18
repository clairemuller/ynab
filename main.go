package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("token.txt")
	check(err)
	defer file.Close()

	txt, err := ioutil.ReadAll(file)
	check(err)
	token := string(txt)

	req, err := http.NewRequest("GET", "https://api.youneedabudget.com/v1/budgets", nil)
	check(err)

	bb := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", bb)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
