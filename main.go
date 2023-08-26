package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type RequestConfig struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Query   map[string]string `json:"query"`
	Form    map[string]string `json:"form"`
	Body    map[string]string `json:"body"`
}

func parseRequest(fileName string) *RequestConfig {
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(path.Join(execDir, fileName))
	if err != nil {
		log.Fatal(err)
	}

	var request *RequestConfig
	err = json.Unmarshal(content, &request)
	if err != nil {
		log.Fatal(err)
	}

	return request
}

func PrettyPrint(data interface{}) {
	content, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", content)
}

func main() {
	if len(os.Args) == 1 {
		panic("Can't run program, no file provided")
	}

	fileName := os.Args[1]
	requestConfig := parseRequest(fileName)
	fmt.Println("===== REQUEST =====")
	PrettyPrint(requestConfig)

	queryParameters := "?"
	for k, v := range requestConfig.Query {
		queryParameters += fmt.Sprintf("%s=%s&", k, v)
	}

	fullUrl := requestConfig.Url + queryParameters

	isFormRequest := len(requestConfig.Form) > 0
	var httpRequest *http.Request
	if isFormRequest {
		formBody := ""
		for k, v := range requestConfig.Form {
			formBody += fmt.Sprintf("%s=%s&", k, v)
		}

		httpBody := strings.NewReader(formBody)
		httpRequest, _ = http.NewRequest(requestConfig.Method, fullUrl, httpBody)
	} else {
		bodyJson, err := json.Marshal(requestConfig.Body)
		if err != nil {
			panic(err)
		}

		httpBody := bytes.NewBuffer(bodyJson)
		httpRequest, _ = http.NewRequest(requestConfig.Method, fullUrl, httpBody)
	}

	httpRequest.Header.Set("Accept", "application/json")
	for k, v := range requestConfig.Headers {
		httpRequest.Header.Set(k, v)
	}

	if isFormRequest {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n===== RESPONSE HEADER =====")
	for k, v := range httpResponse.Header {
		fmt.Printf("%s: %s\n", k, v)
	}

	httpResponseBody, err := io.ReadAll(httpResponse.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(httpResponse.Body)

	fmt.Println("\n\n===== RESPONSE BODY =====")
	httpResponseJson := make(map[string]interface{})
	err = json.Unmarshal(httpResponseBody, &httpResponseJson)
	if err != nil {
		fmt.Println("Error parsing JSON response body")
	} else {
		PrettyPrint(httpResponseJson)
	}
}
