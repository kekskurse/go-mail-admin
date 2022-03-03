package main

import (
	"encoding/json"
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

func (m *MailServerConfiguratorInterface) getAccounts(w http.ResponseWriter, r *http.Request) {
	result, err := m.DBConn.Query("SELECT id, username, domain, password, quota, enabled, sendonly FROM accounts ORDER BY id")

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

func (m *MailServerConfiguratorInterface) addAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account AccountWithPassword
	json.Unmarshal(body, &account)

	pwHash, err := m.HashBuilder.Hash(account.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := m.DBConn.Prepare("INSERT INTO accounts (`username`, `domain`, `password`, `quota`, `enabled`, `sendonly`) VALUES(?, ?, ? ,? , ? , ?)")
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

func (m *MailServerConfiguratorInterface) deleteAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account Account
	json.Unmarshal(body, &account)

	stmt, err := m.DBConn.Prepare("DELETE FROM accounts WHERE id = ?")
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

func (m *MailServerConfiguratorInterface) updateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account Account
	json.Unmarshal(body, &account)

	stmt, err := m.DBConn.Prepare("UPDATE accounts SET quota = ?, enabled = ?, sendonly = ? WHERE id = ?")
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

func (m *MailServerConfiguratorInterface) updateAccountPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var account AccountWithPassword
	json.Unmarshal(body, &account)

	pwHash, err := m.HashBuilder.Hash(account.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	stmt, err := m.DBConn.Prepare("UPDATE accounts SET password = ? WHERE id = ?")
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
