package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Can't run program, no config file provided")
	}

	run()
}

func run() {
	configs := parseConfig(os.Args[1])
	var responses map[int]interface{}

	for i, config := range configs {
		queryParameter := parseQueryParameter(&config)
		fullUrl := config.Url + queryParameter

		buffer, ctype := parseFormBody(&config)
		if buffer == nil {
			buffer, ctype = parseJSONBody(&config)
		}

		request, err := http.NewRequest(config.Method, fullUrl, buffer)
		if err != nil {
			panic(err)
		}

		if ctype != "" {
			request.Header.Set("Content-Type", ctype)
		}

		for k, v := range config.Headers {
			request.Header.Set(k, v)
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}

		body := response.Body
		responseBody, err := io.ReadAll(body)
		if err != nil {
			panic(err)
		}

		err = body.Close()
		if err != nil {
			panic(err)
		}

		var responseJSON interface{}
		err = json.Unmarshal(responseBody, &responseJSON)
		if err != nil {
			panic(err)
		}

		responses[i] = responseJSON
		prettyPrint(responseJSON)
	}
}

func prettyPrint(v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", content)
}
