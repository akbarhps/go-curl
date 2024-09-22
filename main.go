package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("Can't run program, no config file provided")
	}

	run(parseConfig(os.Args[1]))
}

func run(configs []RequestConfig) {
	responses := make(map[int]interface{})
	next := true

	fmt.Print("[")
	for i, config := range configs {
		ctype := ""
		buffer := bytes.NewBuffer(nil)

		if len(config.Form) > 0 {
			buffer, ctype = parseFormBody(&config, responses)
		}
		if len(config.JSON) > 0 {
			buffer, ctype = parseJSONBody(&config, responses)
		}

		url := parseUrl(&config, responses)
		request, err := http.NewRequest(config.Method, url, buffer)
		if err != nil {
			panic(err)
		}

		if len(ctype) > 0 {
			request.Header.Set("Content-Type", ctype)
		}
		for k, v := range config.Headers {
			request.Header.Set(k, replaceReferenceWithValue(v, responses))
		}

		timestamp := time.Now()
		response, err := (&http.Client{}).Do(request)
		if err != nil {
			panic(err)
		}

		body := response.Body
		bodyBuffer, err := io.ReadAll(body)
		if err != nil {
			panic(err)
		}

		if err = body.Close(); err != nil {
			panic(err)
		}

		var bodyStr interface{}
		if err = json.Unmarshal(bodyBuffer, &bodyStr); err != nil {
			bodyStr = "\"" + string(bodyBuffer) + "\""
			next = false
		}

		headers := make(map[string]interface{})
		for k, v := range response.Header {
			isJSON := strings.HasPrefix(v[0], "{") || strings.HasPrefix(v[0], "[")
			if !isJSON || !json.Valid([]byte(v[0])) {
				headers[k] = v[0]
				continue
			}

			var m map[string]interface{}
			if err := json.Unmarshal([]byte(v[0]), &m); err != nil {
				headers[k] = v[0]
				continue
			}
			headers[k] = m
		}

		jsonStr, err := json.MarshalIndent(map[string]interface{}{
			"_____url":    url,
			"____status":  response.Status,
			"___duration": time.Since(timestamp).String(),
			"__headers":   headers,
			"_body":       bodyStr,
		}, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", jsonStr)

		if !next || response.StatusCode != 200 {
			break
		}

		if i+1 < len(configs) {
			fmt.Println(",\n\n")
		}

		responses[i] = bodyStr
	}
	fmt.Print("]")
}
