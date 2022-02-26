package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
)

var tokenAuth *jwtauth.JWTAuth

func (m *MailServerConfiguratorInterface) login(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msgf("Login new JWT Function")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	loginData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	loginResult := struct {
		Login bool   `json:"login"`
		Token string `json:"token"`
	}{
		Login: false,
		Token: "",
	}

	json.Unmarshal(body, &loginData)

	if loginData.Username == m.Config.AuthUsername {
		if loginData.Password == m.Config.AuthPassword {
			_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"admin": true})
			loginResult.Token = tokenString
			loginResult.Login = true
		}
	}

	ren := render.New()
	ren.JSON(w, http.StatusOK, loginResult)
}
