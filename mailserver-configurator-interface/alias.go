package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// Alias from MYSQL
type Alias struct {
	ID                  int    `json:"id"`
	SourceUsername      *string `json:"source_username"`
	SourceDomain        *string `json:"source_domain"`
	DestinationUsername *string `json:"destination_username"`
	DestinationDomain   *string `json:"destination_domain"`
	Enabled             bool   `json:"enabled"`
	PrintSource         string `json:"print_source"`
	PrintDestination    string `json:"print_destination"`
}

func getAliases(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * FROM aliases ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var aliases []Alias

	for result.Next() {
		var alias = Alias{}
		err := result.Scan(&alias.ID, &alias.SourceUsername, &alias.SourceDomain, &alias.DestinationUsername, &alias.DestinationDomain, &alias.Enabled)
		if err != nil {
			log.Fatal(err)
		}

		username := ""

		if alias.SourceUsername != nil {
			username = *alias.SourceUsername
		}

		alias.PrintSource = username + "@" + *alias.SourceDomain
		alias.PrintDestination = *alias.DestinationUsername + "@" + *alias.DestinationDomain
		aliases = append(aliases, alias)
	}
	ren := render.New()
	ren.JSON(w, http.StatusOK, aliases)
}
func addAlias(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var alias Alias
	json.Unmarshal(body, &alias)

	if alias.SourceUsername != nil && *alias.SourceUsername == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Source Username can`t be empty string, only null or string is valid"))
		return
	}

	stmt, err := db.Prepare("INSERT INTO aliases (`source_username`, `source_domain`, `destination_username`, `destination_domain`, `enabled`) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(alias.SourceUsername, alias.SourceDomain, alias.DestinationUsername, alias.DestinationDomain, alias.Enabled)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func deleteAlias(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var alias Alias
	json.Unmarshal(body, &alias)

	stmt, err := db.Prepare("DELETE FROM aliases WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(alias.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func updateAlias(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var alias Alias
	json.Unmarshal(body, &alias)

	stmt, err := db.Prepare("UPDATE aliases SET source_username = ?, source_domain = ?, destination_username = ?, destination_domain = ?, enabled = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = stmt.Exec(alias.SourceUsername, alias.SourceDomain, alias.DestinationUsername, alias.DestinationDomain, alias.Enabled, alias.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
