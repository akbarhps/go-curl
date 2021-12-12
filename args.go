package main

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"os"
	"strings"
)

type AppArgs struct {
	Method  string
	Url     string
	Body    map[string]string
	Headers map[string]string
}

var allowedMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
}

func ParseArgs() (*AppArgs, error) {
	args := &AppArgs{}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if strings.HasPrefix(arg, "-b=") {
			args.Body = make(map[string]string)
			value := strings.Replace(strings.TrimSpace(arg), "-b=", "", 1)

			if err := stringToMap(value, args.Body); err != nil {
				return nil, err
			}
			continue
		}

		if strings.HasPrefix(arg, "-h=") {
			args.Headers = make(map[string]string)
			value := strings.Replace(strings.TrimSpace(arg), "-h=", "", 1)

			if err := stringToMap(value, args.Headers); err != nil {
				return nil, err
			}
			continue
		}

		if contains(allowedMethods, strings.ToUpper(arg)) {
			args.Method = strings.ToUpper(arg)
			continue
		}

		if ok := govalidator.IsURL(arg); ok {
			args.Url = arg
		}
	}

	return args, nil
}

func stringToMap(s string, m map[string]string) error {
	if !strings.Contains(s, "=") {
		return errors.New("invalid format")
	}

	for _, v := range strings.Split(s, ";") {
		if v == "" {
			continue
		}

		kv := strings.Split(strings.TrimSpace(v), "=")
		m[kv[0]] = kv[1]
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
