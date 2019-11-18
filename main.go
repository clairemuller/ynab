package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	token := getToken()
	resp := sendHeader(token)

	fmt.Println(resp)
}

func getToken() string {
	file, err := os.Open("token.txt")
	check(err)
	defer file.Close()

	txt, err := ioutil.ReadAll(file)
	check(err)
	return string(txt)
}

func sendHeader(token string) *http.Response {
	req, err := http.NewRequest("GET", "https://api.youneedabudget.com/v1/budgets", nil)
	check(err)

	bb := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", bb)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	return resp
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
