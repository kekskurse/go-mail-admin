package main

import (
	"fmt"
	"net/http"
	"log"
)

type auth struct {
	HTTPBasicAuthEnabled bool
	HTTPBasicAuthUser string
	HTTPBasicAuthPassword string

}

// NewAuthFromEnv load the env based of go-mail-admin to make it a bit easyer to use
func NewAuthFromEnv() auth {
	a := auth{}

	apiKey := getConfigVariable("APIKEY")
	apiSecret := getConfigVariable("APISECRET")

	a.HTTPBasicAuthEnabled = false
	if apiKey != "" && apiSecret != "" {
		a.HTTPBasicAuthEnabled = true
		a.HTTPBasicAuthUser = apiKey
		a.HTTPBasicAuthPassword = apiSecret
		log.Println("AUTH: Enabled HTTP Basic Auth")
	} else {
		log.Println("AUTH: Run without any Auth-Protection")
	}
	return a
}

func (a auth) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check HTTP-Basic Auth
		if a.HTTPBasicAuthEnabled {
			username, password, ok := r.BasicAuth()
			if !ok {
				a.httpBasicAuthUnauthorized(w, "MyRealm")
				return
			}

			ok, err := a.httpBasicAuthCheck(username, password)
			if err != nil {
				log.Panic(err)
			}
			if !ok {
				a.httpBasicAuthUnauthorized(w, "MyRealm")
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (a auth) httpBasicAuthUnauthorized(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}


func (a auth) httpBasicAuthCheck(username string, password string) (ok bool, err error) {
	ok = false //Just make clear that the default is false
	if !a.HTTPBasicAuthEnabled {
		return
	}

	if username == a.HTTPBasicAuthUser && password == a.HTTPBasicAuthPassword {
		ok = true
		return
	}

	return
}