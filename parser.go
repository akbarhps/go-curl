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
	references := make([]Reference, 0)

	re := regexp.MustCompile(`{(.*?)}`)
	matches := re.FindAllStringSubmatch(value, -1)
	if len(matches) == 0 {
		return nil
	}

	for _, match := range matches {
		// match contains: [{0->0->id} 0->0->id]
		ref := match[1]
		// split ref by ->
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
			// NOTE: skip if the value is not a string
			continue
		}
		refMap[k] = value
	}

	return refMap
}

func replaceReferenceWithValue(oldValue string, ref Reference, refValues map[int]interface{}) string {
	refMap := parseReferenceAsMap(ref, refValues)
	return strings.ReplaceAll(oldValue, ref.RawReference, refMap[ref.Key])
}

func parseQueryParameter(config *RequestConfig, refValues map[int]interface{}) string {
	if len(config.Query) == 0 {
		return ""
	}

	query := "?"
	for k, v := range config.Query {
		value := v
		refs := parseReference(v)

		if refs != nil {
			for _, ref := range *refs {
				value = replaceReferenceWithValue(value, ref, refValues)
			}
		}

		query += fmt.Sprintf("%s=%s&", k, value)
	}

	return strings.TrimSuffix(query, "&")
}

func replacePathInUrl(url string, refValues map[int]interface{}) string {
	refs := parseReference(url)
	if refs == nil {
		return url
	}

	newUrl := url
	for _, ref := range *refs {
		newUrl = replaceReferenceWithValue(newUrl, ref, refValues)
	}

	return newUrl
}

func parseUrl(config *RequestConfig, refValues map[int]interface{}) string {
	return replacePathInUrl(config.Url, refValues) + parseQueryParameter(config, refValues)
}

func parseFormBody(config *RequestConfig, refValues map[int]interface{}) (*bytes.Buffer, string) {
	if len(config.Form) == 0 {
		return bytes.NewBuffer(nil), ""
	}

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

		refs := parseReference(v)
		if refs != nil {
			for _, ref := range *refs {
				v = replaceReferenceWithValue(v, ref, refValues)
			}
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

func parseJSONBody(config *RequestConfig, refValues map[int]interface{}) (*bytes.Buffer, string) {
	if len(config.JSON) == 0 {
		return bytes.NewBuffer(nil), ""
	}

	newJSON := make(map[string]interface{})
	for k, v := range config.JSON {
		value, ok := v.(string)
		if !ok {
			newJSON[k] = v
			continue
		}

		refs := parseReference(value)
		if refs == nil {
			newJSON[k] = v
			continue
		}

		for _, ref := range *refs {
			value = replaceReferenceWithValue(value, ref, refValues)
		}
		newJSON[k] = value
	}

	body, err := json.Marshal(newJSON)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(body), "application/json"
}
