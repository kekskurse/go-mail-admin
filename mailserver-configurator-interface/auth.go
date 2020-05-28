package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
)

type auth struct {
	Method string

	//For HTTP Basic auth
	HTTPBasicAuthUsername string
	HTTPBasicAuthPassword string

	//For Admin Mail
	AdminMailMails []string
	AdminMailAPIKeys []string

}

// NewAuthFromEnv load the env based of go-mail-admin to make it a bit easyer to use
func NewAuthFromEnv() auth {
	a := auth{}

	authMethod := getConfigVariableWithDefault("AUTH_METHOD", "")

	apiKey := getConfigVariable("APIKEY")
	apiSecret := getConfigVariable("APISECRET")

	//Switch it to the httpBasicAuth Block if the old Auth enviroment is removed
	httpBasicAuthUsername := getConfigVariable("AUTH_HTTPBasic_Username")
	httpBasicAuthPassword := getConfigVariable("AUTH_HTTPBasic_Password")

	if apiKey != "" && apiSecret != "" {
		if authMethod == "AdminMail" {
			panic("Auth method is set to AdminMail but APIKEY and APISECRET is set!")
		}
		authMethod = "HTTPBasicAuth"
		httpBasicAuthUsername = apiKey
		httpBasicAuthPassword = apiSecret
		log.Println("WARNING! The old Enviroment varieable GOMAILADMIN_APIKEY or GOMAILADMIN_APISECRET are set. The Auth method is forced to HTTPBasicAuth, please read the new auth docs to change your configuration")
		// At a panic in the feature
	}

	a.Method = authMethod

	if a.Method == "HTTPBasicAuth" {
		log.Println("Auth: Enabled HTTPBasicAuth")
		a.HTTPBasicAuthUsername = httpBasicAuthUsername
		a.HTTPBasicAuthPassword = httpBasicAuthPassword
	}

	if a.Method == "" {
		log.Println("No Auth Method is set. Auth Method is forced to None. Please read the auth doc and set a variable. It is DEPRECATED to start go-mail-admin without!")
		a.Method = "None"
	}

	if a.Method == "AdminMail" {
		log.Println("Auth: Enabled AdminMail")
		adminMails := strings.Split(getConfigVariable("AUTH_AdminMail_MAIL"), ",")
		apiKeys := strings.Split(getConfigVariable("AUTH_AdminMail_API"), ",")

		if len(adminMails) == 1 && adminMails[0] == "" {
			log.Println("Auth: AdminMail is used but not Admin E-Mail address is set, no one can use the webfrontend")
		} else {
			a.AdminMailMails = adminMails
		}

		if len(apiKeys) == 1 && apiKeys[0] == "" {
			log.Println("Auth: AdminMail is used but not API Keys are set, if you want to use the API you have to login via mail/password")
		} else {
			a.AdminMailAPIKeys = apiKeys
		}

		if len(a.AdminMailAPIKeys) == 0 && len(a.AdminMailMails) == 0 {
			panic("Auth: AdminMail is used but no Admin E-Mail address or API-Key is set")
		}

	}

	return a
}

func (a *auth) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check HTTP-Basic Auth
		if a.Method == "HTTPBasicAuth" {
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

			next.ServeHTTP(w, r)
			return
		}

		if a.Method == "None" {
			next.ServeHTTP(w, r)
			return
		}

		if a.Method == "AdminMail" {
			panic("Auth method AdminMail is not done yet")
		}

		panic("No valid Auth Method is set")

	})
}

func (a *auth) httpBasicAuthUnauthorized(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}


func (a *auth) httpBasicAuthCheck(username string, password string) (ok bool, err error) {
	ok = false //Just make clear that the default is false

	if username == a.HTTPBasicAuthUsername && password == a.HTTPBasicAuthPassword {
		ok = true
		return
	}

	log.Println("HTTP Basic auth for user >" + username + "< failed")

	return
}