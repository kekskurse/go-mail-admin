package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// Domain from MYSQL
type Domain struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
}

func getDomains(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * FROM domains ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var domains []Domain

	for result.Next() {
		var domain = Domain{}
		err := result.Scan(&domain.ID, &domain.Domain)
		if err != nil {
			log.Fatal(err)
		}
		domains = append(domains, domain)
	}
	ren := render.New()
	ren.JSON(w, http.StatusOK, domains)
	//json.NewEncoder(w).Encode(domains)
}

func addDomain(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var domain Domain
	json.Unmarshal(body, &domain)

	stmt, err := db.Prepare("INSERT INTO domains(domain) VALUES(?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(domain.Domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteDomain(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var domain Domain
	json.Unmarshal(body, &domain)

	stmt, err := db.Prepare("DELETE FROM domains WHERE domain = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(domain.Domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
