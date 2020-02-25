package main

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"gopkg.in/unrolled/render.v1"
)

//Password Stuff
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytesMaskImprSrc Generate random string in golang
func RandStringBytesMaskImprSrc(n int) string {
	// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Account from MYSQL
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Password string `json:"-"`
	Quota    int    `json:"quota"`
	Enabled  bool   `json:"enabled"`
	SendOnly bool   `json:"sendonly"`
	Print string `json:"print"`
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
	Print string `json:"print"`
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * FROM accounts ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var accounts []Account

	for result.Next() {
		var account = Account{}
		err := result.Scan(&account.ID, &account.Username, &account.Domain, &account.Password, &account.Quota, &account.Enabled, &account.SendOnly)
		if err != nil {
			log.Fatal(err)
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

	//Hash password
	salt := RandStringBytesMaskImprSrc(5)
	hasher := sha512.New()
	hasher.Write([]byte(account.Password))
	hasher.Write([]byte(salt))

	pwHash := "{SSHA512}" + base64.StdEncoding.EncodeToString([]byte(string(hasher.Sum(nil))+salt))

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

	salt := RandStringBytesMaskImprSrc(5)
	hasher := sha512.New()
	hasher.Write([]byte(account.Password))
	hasher.Write([]byte(salt))

	pwHash := "{SSHA512}" + base64.StdEncoding.EncodeToString([]byte(string(hasher.Sum(nil))+salt))

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
