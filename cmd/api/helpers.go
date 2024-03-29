package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// readJSON reads json from request body into data. We only accept a single json value in the body
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // max one megabyte in request body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// we only allow one entry in the json file
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

// badRequest sends a JSON response with status http.StatusBadRequest, describing the error
func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	payload.Error = true
	payload.Message = err.Error()

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(out)
	return nil
}

// writeJson function helps to write json response
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap ...string) error {
	var js []byte
	var err error
	if len(wrap) > 0 {

		wrapper := make(map[string]interface{})
		wrapper[wrap[0]] = data

		js, err = json.Marshal(wrapper)
		if err != nil {
			return err
		}
	} else {
		js, err = json.Marshal(data)
		if err != nil {
			return err
		}
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

// verify authentication token
func (app *application) verifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.jwt.secret), nil
	})

	// if token is invalid then return error
	if err != nil || !token.Valid {
		return nil, errors.New("unauthorized - invalid token")
	}

	// Get the custom claims from the token
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("unauthorized - invalid token")
	}

	// check if token is expired
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("unauthorized - token expired")
	}

	// return claims
	return claims, nil
}

// parseHeaderToken extracts the Authorization token from an HTTP response.
func (app *application) parseHeaderToken(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is present
	if authHeader == "" {
		return 0, errors.New("authorization header is missing")
	}

	// Split the header value to get the token part (e.g., "Bearer <token>")
	headerParts := strings.Split(authHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, errors.New("invalid authorization header format")
	}

	// Extract the token
	token := headerParts[1]

	claims, err := app.verifyToken(token)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, errors.New("invalid user id")
	}

	return userID, nil
}
