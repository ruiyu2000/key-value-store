package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var (
		node    = args[0]
		command = args[1]
		key     = args[2]
		method  string
		body    io.Reader
	)

	switch command {
	case "get":
		method = http.MethodGet
	case "set":
		method = http.MethodPost
		value := args[3]
		body = strings.NewReader("value=" + value)
	case "delete":
		method = http.MethodDelete
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", node, key), body)
	if err != nil {
		return
	}

	if command == "set" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}

	c := http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{}},
	}

	response, err := c.Do(request)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
