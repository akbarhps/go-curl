package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

type RequestConfig struct {
	Url     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Query   map[string]string      `json:"query"`
	Form    map[string]string      `json:"form"`
	Body    map[string]interface{} `json:"body"`
}

func getFile(filePath string) (string, []byte) {
	fullPath := filePath

	// If file path is relative, get the full path
	if strings.HasPrefix(filePath, "./") {
		execDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		fullPath = path.Join(execDir, filePath)
	}

	file, err := os.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}

	return fullPath, file
}

func parseRequestConfig(fileName string) []RequestConfig {
	var request []RequestConfig
	_, content := getFile(fileName)

	err := json.Unmarshal(content, &request)
	if err != nil {
		panic(err)
	}

	return request
}

func PrettyPrint(data interface{}) {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", content)
}

func main() {
	if len(os.Args) == 1 {
		panic("Can't run program, no file provided")
	}

	fileName := os.Args[1]
	requestConfigs := parseRequestConfig(fileName)
	size := len(requestConfigs)

	for i, requestConfig := range requestConfigs {
		fmt.Println("Request #", i+1, "of", size)
		fmt.Println("::REQUEST::")
		PrettyPrint(requestConfig)

		queryParameters := "?"
		for k, v := range requestConfig.Query {
			queryParameters += fmt.Sprintf("%s=%s&", k, v)
		}
		queryParameters = strings.TrimSuffix(queryParameters, "&")

		fullUrl := requestConfig.Url + queryParameters

		isFormRequest := len(requestConfig.Form) > 0
		var httpRequest *http.Request
		if isFormRequest {
			buf := new(bytes.Buffer)
			m := multipart.NewWriter(buf)

			for k, v := range requestConfig.Form {
				if strings.HasPrefix(v, "file:") {
					filePath := strings.TrimSpace(strings.TrimPrefix(v, "file:"))
					fullPath, file := getFile(filePath)

					part, err := m.CreateFormFile(k, path.Base(fullPath))
					if err != nil {
						panic(err)
					}

					_, err = part.Write(file)
					if err != nil {
						panic(err)
					}
				} else {
					err := m.WriteField(k, v)
					if err != nil {
						panic(err)
					}
				}
			}

			err := m.Close()
			if err != nil {
				panic(err)
			}

			httpRequest, _ = http.NewRequest(requestConfig.Method, fullUrl, buf)
			httpRequest.Header.Set("Content-Type", m.FormDataContentType())
		} else {
			bodyJson, err := json.Marshal(requestConfig.Body)
			if err != nil {
				panic(err)
			}

			httpBody := bytes.NewBuffer(bodyJson)
			httpRequest, _ = http.NewRequest(requestConfig.Method, fullUrl, httpBody)
			httpRequest.Header.Set("Content-Type", "application/json")
		}

		httpRequest.Header.Set("Accept", "application/json")
		for k, v := range requestConfig.Headers {
			httpRequest.Header.Set(k, v)
		}

		client := &http.Client{}
		httpResponse, err := client.Do(httpRequest)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n\n::RESPONSE HEADER::")
		fmt.Println("Status Code:", httpResponse.StatusCode)
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

		fmt.Println("\n\n::RESPONSE BODY::")

		httpResponseJson := make(map[string]interface{})
		err = json.Unmarshal(httpResponseBody, &httpResponseJson)
		if err != nil {
			httpResponseJson := make([]map[string]interface{}, 0)
			err = json.Unmarshal(httpResponseBody, &httpResponseJson)

			if err != nil {
				fmt.Println(string(httpResponseBody))
			} else {
				PrettyPrint(httpResponseJson)
			}
		} else {
			PrettyPrint(httpResponseJson)
		}

		if i != size-1 {
			fmt.Println("\n\n\n")
		}
	}
}
