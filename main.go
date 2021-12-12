package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	run()
}

func run() {
	args, err := ParseArgs()
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(args.Body)
	if err != nil {
		panic(err)
	}

	reqBody := bytes.NewBuffer(body)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(args.Method, args.Url, reqBody)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	for k, v := range args.Headers {
		request.Header.Add(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	for k, v := range response.Header {
		fmt.Printf("%s: %s\n", k, v)
	}

	resBodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println("\n")
	fmt.Println(string(resBodyBytes))
}
