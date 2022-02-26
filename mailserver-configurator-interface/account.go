package main

import (
	"encoding/json"
	"gomailadmin/mailserver-configurator-interface/password"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/unrolled/render"
)

// Account from MYSQL
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Password string `json:"-"`
	Quota    int    `json:"quota"`
	Enabled  bool   `json:"enabled"`
	SendOnly bool   `json:"sendonly"`
	Print    string `json:"print"`
}

// AccountWithPassword from MYSQL
type AccountWithPassword struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
	Quota    int    `json:"quota"`
	Enabled  bool   `json:"enabled"`
	SendOnly bool   `json:"sendonly"`
	Print    string `json:"print"`
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT id, username, domain, password, quota, enabled, sendonly FROM accounts ORDER BY id")

	if err != nil {
		log.Fatal().Err(err).Msg("Error execute query")
	}
	defer result.Close()

	var accounts []Account

	for result.Next() {
		var account = Account{}
		err := result.Scan(&account.ID, &account.Username, &account.Domain, &account.Password, &account.Quota, &account.Enabled, &account.SendOnly)
		if err != nil {
			log.Fatal().Err(err).Msg("Error Scan query result")
		}
		account.Print = account.Username + "@" + account.Domain
		accounts = append(accounts, account)
	}
	ren := render.New()
	ren.JSON(w, http.StatusOK, accounts)
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account AccountWithPassword
	json.Unmarshal(body, &account)

	pwHash, err := password.GetPasswordHashBuilder("ssha512").Hash(account.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := db.Prepare("INSERT INTO accounts (`username`, `domain`, `password`, `quota`, `enabled`, `sendonly`) VALUES(?, ?, ? ,? , ? , ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(account.Username, account.Domain, pwHash, account.Quota, account.Enabled, account.SendOnly)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account Account
	json.Unmarshal(body, &account)

	stmt, err := db.Prepare("DELETE FROM accounts WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(account.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account Account
	json.Unmarshal(body, &account)

	stmt, err := db.Prepare("UPDATE accounts SET quota = ?, enabled = ?, sendonly = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(account.Quota, account.Enabled, account.SendOnly, account.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateAccountPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account AccountWithPassword
	json.Unmarshal(body, &account)

	pwHash, err := password.GetPasswordHashBuilder("ssha512").Hash(account.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := db.Prepare("UPDATE accounts SET password = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(pwHash, account.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
