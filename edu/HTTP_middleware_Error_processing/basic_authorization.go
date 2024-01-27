package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "" {
			fields := strings.Fields(auth)
			if len(fields) == 2 && fields[0] == "Basic" {
				credentials, err := base64.StdEncoding.DecodeString(fields[1])
				if err == nil {
					cred := string(credentials)
					if cred == "userid:password" {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	})
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The answer is 42")
}

func main() {
	mux := http.NewServeMux()
	ans := http.HandlerFunc(answerHandler)
	mux.Handle("/", Authorization(ans))
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatal(err)
	}
}
