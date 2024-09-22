package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const (
	// InputFilePrefix file prefix can be relative or full path
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

type Reference struct {
	RawReference string
	Index        int
	KeyIndex     int
	Key          string
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

func parseReference(value string) *[]Reference {
	re := regexp.MustCompile(`{(.*?)}`)
	matches := re.FindAllStringSubmatch(value, -1)
	if len(matches) == 0 {
		return nil
	}

	references := make([]Reference, 0)
	for _, match := range matches {
		// match contains: [{0->0->id} 0->0->id]
		ref := match[1]
		items := strings.Split(ref, "->")

		index, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}

		if len(items) == 2 {
			references = append(references, Reference{
				RawReference: match[0],
				Index:        index,
				KeyIndex:     -1,
				Key:          items[1],
			})
			continue
		}

		keyIndex, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}

		references = append(references, Reference{
			RawReference: match[0],
			Index:        index,
			KeyIndex:     keyIndex,
			Key:          items[2],
		})
	}

	return &references
}

func parseReferenceAsMap(ref Reference, refValues map[int]interface{}) map[string]string {
	var rawMap map[string]interface{}

	// if KeyIndex is -1, it means the reference is already a map
	// if not, it means the reference is an array of map
	if ref.KeyIndex == -1 {
		rawMap, _ = refValues[ref.Index].(map[string]interface{})
	} else {
		mSlice, _ := refValues[ref.Index].([]interface{})
		rawMap = mSlice[ref.KeyIndex].(map[string]interface{})
	}

	refMap := make(map[string]string)
	for k, v := range rawMap {
		value, ok := v.(string)
		if !ok {
			continue
		}
		refMap[k] = value
	}

	return refMap
}

func replaceReferenceWithValue(text string, refValues map[int]interface{}) string {
	refs := parseReference(text)
	if refs == nil {
		return text
	}

	newValue := text
	for _, ref := range *refs {
		refMap := parseReferenceAsMap(ref, refValues)
		newValue = strings.ReplaceAll(text, ref.RawReference, refMap[ref.Key])
	}
	return newValue
}

func parseQueryParameter(config *RequestConfig, refValues map[int]interface{}) string {
	if len(config.Query) == 0 {
		return ""
	}

	query := "?"
	for k, v := range config.Query {
		query += fmt.Sprintf("%s=%s&", k, replaceReferenceWithValue(v, refValues))
	}

	return strings.TrimSuffix(query, "&")
}

func parseUrl(config *RequestConfig, refValues map[int]interface{}) string {
	return replaceReferenceWithValue(config.Url, refValues) + parseQueryParameter(config, refValues)
}

func parseFormBody(config *RequestConfig, refValues map[int]interface{}) (*bytes.Buffer, string) {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)

	for k, v := range config.Form {
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

		err := writer.WriteField(k, replaceReferenceWithValue(v, refValues))
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

func parseJSONBody(config *RequestConfig, refValues map[int]interface{}) (*bytes.Buffer, string) {
	newJSON := make(map[string]interface{})
	for k, v := range config.JSON {
		value, ok := v.(string)
		if !ok {
			newJSON[k] = v
			continue
		}

		newJSON[k] = replaceReferenceWithValue(value, refValues)
	}

	body, err := json.Marshal(newJSON)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(body), "application/json"
}
