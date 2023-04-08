package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// struct for query parameters
type queryParam struct {
	key   string
	value interface{}
}

// helper function to validate query parameters and return an error if invalid
func validateQueryParam(key, value string) (interface{}, error) {
	switch key {
	case "s":
		return strings.ToLower(value), nil
	case "page", "limit", "genre", "year":
		return strconv.Atoi(value)
	case "order_by":
		switch value {
		case "rating", "runtime", "old", "name":
			return value, nil
		default:
			return "", errors.New("invalid order_by value")
		}
	default:
		return "", errors.New("invalid query parameter")
	}
}

// writeJson function helps to write json response
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// errorJSON function helps to write error json response
func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, statusCode, theError, "error")
}
