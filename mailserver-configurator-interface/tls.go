package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// TLSPolicy from MYSQL
type TLSPolicy struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
	Policy string `json:"policy"`
	Params string `json:"params"`
}

func getTLSPolicy(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT id, domain, policy, IFNULL(params, \"\") FROM tlspolicies ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var policys []TLSPolicy

	for result.Next() {
		var policy = TLSPolicy{}
		err := result.Scan(&policy.ID, &policy.Domain, &policy.Policy, &policy.Params)
		if err != nil {
			log.Fatal(err)
		}
		policys = append(policys, policy)
	}
	ren := render.New()
	ren.JSON(w, http.StatusOK, policys)
}

func addTLSPolicy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var tlspolicy TLSPolicy
	json.Unmarshal(body, &tlspolicy)

	stmt, err := db.Prepare("INSERT INTO tlspolicies (`domain`, `policy`, `params`) VALUES(?, ?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(tlspolicy.Domain, tlspolicy.Policy, tlspolicy.Params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func updateTLSPolicy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var policy TLSPolicy
	json.Unmarshal(body, &policy)

	stmt, err := db.Prepare("UPDATE tlspolicies SET domain = ?, policy = ?, params = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(policy.Domain, policy.Policy, policy.Params, policy.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func deleteTLSPolicy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var policy TLSPolicy
	json.Unmarshal(body, &policy)

	stmt, err := db.Prepare("DELETE FROM tlspolicies WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(policy.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
