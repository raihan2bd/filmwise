package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// add custom type for context key
type userIDKey string

// enableCORS adds the Access-Control-Allow-Origin header to all responses.
func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

// authenticate checks whether a request is coming from an authenticated user.
func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		// if auth header is empty then return error
		if authHeader == "" {
			app.errorJSON(w, errors.New("authorization header is required"))
			return
		}

		headerParts := strings.Split(authHeader, " ")

		// if auth header is not two parts then return error
		if len(headerParts) != 2 {
			app.errorJSON(w, errors.New("invalid auth header"))
			return
		}

		// if auth header doesn't include Bearer then return error
		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("unauthorized - no bearer"))
			return
		}

		// verify token
		tokenString := headerParts[1]

		// parse token, and return claims if there is not error
		claims, err := app.verifyToken(tokenString)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		// check user_type is allowed
		if claims.UserType == "admin" || claims.UserType == "user" {
			// add user_id to context
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized - user does not is not valid"))
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey("user_id"), userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			// next.ServeHTTP(w, r)
		} else {
			app.errorJSON(w, errors.New("unauthorized - user does not have permission"))
			return
		}

	}) // end of http.HandlerFunc
}

// auth middleware for admin

// authenticate checks whether a request is coming from an authenticated user.
func (app *application) adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		// if auth header is empty then return error
		if authHeader == "" {
			app.errorJSON(w, errors.New("authorization header is required"))
			return
		}

		headerParts := strings.Split(authHeader, " ")

		// if auth header is not two parts then return error
		if len(headerParts) != 2 {
			app.errorJSON(w, errors.New("invalid auth header"))
			return
		}

		// if auth header doesn't include Bearer then return error
		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("unauthorized - no bearer"))
			return
		}

		// verify token
		tokenString := headerParts[1]

		// parse token, and return claims if there is not error
		claims, err := app.verifyToken(tokenString)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		// check user_type is allowed
		if claims.UserType == "admin" {
			// add user_id to context
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized - user does not valid"))
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey("user_id"), userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			// next.ServeHTTP(w, r)
		} else {
			app.errorJSON(w, errors.New("unauthorized - user does not have permission"))
			return
		}

	}) // end of http.HandlerFunc
}
