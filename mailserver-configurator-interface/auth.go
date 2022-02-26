package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
)

type auth struct {
	Method string

	//For HTTP Basic auth
	HTTPBasicAuthUsername string
	HTTPBasicAuthPassword string

	//For Username
	Username string
	Password string

	//For Admin Mail
	AdminMailMails   []string
	AdminMailAPIKeys []string

	//Redis
	Redis redisConnection
}

// NewAuthFromEnv load the env based of go-mail-admin to make it a bit easyer to use
func NewAuthFromEnv(r redisConnection, config Config) auth {
	a := auth{}
	a.Redis = r

	//Switch it to the httpBasicAuth Block if the old Auth enviroment is removed
	httpBasicAuthUsername := config.HttpBasicAuthUsername
	httpBasicAuthPassword := config.HttpBasicAuthPassword

	if config.ApiKey != "" && config.ApiSecret != "" {
		if config.AuthMethod == "AdminMail" {
			panic("Auth method is set to AdminMail but APIKEY and APISECRET is set!")
		}
		config.AuthMethod = "HTTPBasicAuth"
		httpBasicAuthUsername = config.ApiKey
		httpBasicAuthPassword = config.ApiSecret
	}

	a.Method = config.AuthMethod

	if a.Method == "HTTPBasicAuth" {
		log.Debug().Msg("Auth: Enabled HTTPBasicAuth")
		a.HTTPBasicAuthUsername = httpBasicAuthUsername
		a.HTTPBasicAuthPassword = httpBasicAuthPassword
	}

	if a.Method == "" {
		log.Debug().Msg("No Auth Method is set. Auth Method is forced to None. Please read the auth doc and set a variable. It is DEPRECATED to start go-mail-admin without!")
		a.Method = "None"
	}

	if a.Method == "Username" {
		log.Debug().Msg("Auth: Enabled Username")
		a.Username = config.AuthUsername
		a.Password = config.AuthPassword
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
				log.Fatal().Err(err).Msg("Error check auth")
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

		if a.Method == "Username" {
			token := r.Header.Get("X-APITOKEN")
			v, _ := a.Redis.get("auth_" + token)
			if v == "1" {
				next.ServeHTTP(w, r)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Token invalide"))
				return
			}
		}

		/*if a.Method == "AdminMail" {
			panic("Auth method AdminMail is not done yet")
		}*/

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

	log.Debug().Msg(fmt.Sprintf("HTTP Basic auth for user >%s< failed", username))

	return
}

/*
	Call this function to create a valide token for api-keys
*/
func (a *auth) createToken() (token string, e error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 24)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	token = string(b)

	a.Redis.set("auth_"+token, "1", 60*15)

	return

}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Login bool   `json:"login"`
	Token string `json:"token"`
}

func loginUsername(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	loginResult := LoginResult{}
	loginResult.Login = false

	var loginData LoginData
	json.Unmarshal(body, &loginData)

	if authConfig.Username == loginData.Username && authConfig.Password == loginData.Password {
		loginResult.Login = true
		loginResult.Token, _ = authConfig.createToken()
	}

	ren := render.New()
	ren.JSON(w, http.StatusOK, loginResult)
}

func logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-APITOKEN")

	authConfig.Redis.delete("auth_" + token)

	w.WriteHeader(http.StatusUnauthorized)
	//w.Write([]byte(""))
}
