package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path"
	"strings"
)

const (
	InputFilePrefix = "file:"
)

type RequestConfig struct {
	Url     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Query   map[string]string      `json:"query"`
	Form    map[string]string      `json:"form"`
	JSON    map[string]interface{} `json:"json"`
}

func parseConfig(fileName string) []RequestConfig {
	var request []RequestConfig
	file, _ := readFile(fileName)

	err := json.Unmarshal(file, &request)
	if err != nil {
		panic(err)
	}

	return request
}

func parseQueryParameter(config *RequestConfig) string {
	if len(config.Query) == 0 {
		return ""
	}

	query := "?"
	for k, v := range config.Query {
		query += fmt.Sprintf("%s=%s&", k, v)
	}

	return strings.TrimSuffix(query, "&")
}

func parseFormBody(config *RequestConfig) (*bytes.Buffer, string) {
	if len(config.Form) == 0 {
		return bytes.NewBuffer(nil), ""
	}

	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)

	for k, v := range config.Form {
		// file: prefix indicates that the value is a file path
		// it can be full path or relative path
		if strings.HasPrefix(v, InputFilePrefix) {
			filePath := strings.TrimSpace(strings.TrimPrefix(v, InputFilePrefix))
			file, fileFullPath := readFile(filePath)

			part, err := writer.CreateFormFile(k, path.Base(fileFullPath))
			if err != nil {
				panic(err)
			}

			_, err = part.Write(file)
			if err != nil {
				panic(err)
			}
			continue
		}

		err := writer.WriteField(k, v)
		if err != nil {
			panic(err)
		}
	}

	err := writer.Close()
	if err != nil {
		panic(err)
	}

	return buffer, writer.FormDataContentType()
}

func parseJSONBody(config *RequestConfig) (*bytes.Buffer, string) {
	if len(config.JSON) == 0 {
		return bytes.NewBuffer(nil), ""
	}

	body, err := json.Marshal(config.JSON)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(body), "application/json"
}
