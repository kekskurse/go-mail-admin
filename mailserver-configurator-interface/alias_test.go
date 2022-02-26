package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAliasWithoutValidDomain(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/alias", bytes.NewBufferString("{\"source_username\":\"test\", \"source_domain\":\"foobar\", \"destination_username\":\"foo\", \"destination_domain\":\"google.com\", \"enabled\":true}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()
	resetDBForTest()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(m.addAlias)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		b, _ := ioutil.ReadAll(rr.Body)
		s := string(b)
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d. (%v)", http.StatusOK, status, s)
	}
}

func TestAddAlias(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/alias", bytes.NewBufferString("{\"source_username\":\"test\", \"source_domain\":\"alias.com\", \"destination_username\":\"foo\", \"destination_domain\":\"google.com\", \"enabled\":true}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	stmt, _ := db.Prepare("INSERT INTO domains(domain) VALUES(?)")
	_, _ = stmt.Exec("alias.com")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(m.addAlias)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		b, _ := ioutil.ReadAll(rr.Body)
		s := string(b)
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d. (%v)", http.StatusOK, status, s)
	}
}

func TestGetAliases(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/aliase", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAliases)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var alias []Alias
	if err := json.NewDecoder(rr.Body).Decode(&alias); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(alias) != 1 {
		t.Errorf("Dont return domain list")
	}

	if *alias[0].SourceUsername != "test" {
		t.Errorf("Wrong SourceUsername")
	}

	if *alias[0].SourceDomain != "alias.com" {
		t.Errorf("Wrong SourceDomain")
	}

	if *alias[0].DestinationUsername != "foo" {
		t.Errorf("Wrong SourceDomain")
	}

	if *alias[0].DestinationDomain != "google.com" {
		t.Errorf("Wrong SourceDomain")
	}

	if alias[0].Enabled != true {
		t.Errorf("Enabled is wrong")
	}

	if alias[0].ID != 2 {
		t.Errorf("Alias has wrong id got %v excepted %v", alias[0].ID, 2)
	}
}

func TestUpdateAliasWrongDomain(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/alias", bytes.NewBufferString("{\"id\": 2, \"source_username\":\"test\", \"source_domain\":\"ne.com\", \"destination_username\":\"foo\", \"destination_domain\":\"google.com\", \"enabled\":true}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateAlias)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		b, _ := ioutil.ReadAll(rr.Body)
		s := string(b)
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d. (%v)", http.StatusOK, status, s)
	}
}

func TestUpdateAlias(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/alias", bytes.NewBufferString("{\"id\": 2, \"source_username\":\"updated\", \"source_domain\":\"alias.com\", \"destination_username\":\"foo\", \"destination_domain\":\"google.com\", \"enabled\":true}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	stmt, _ := db.Prepare("INSERT INTO domains(domain) VALUES(?)")
	_, _ = stmt.Exec("alias.com")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateAlias)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		b, _ := ioutil.ReadAll(rr.Body)
		s := string(b)
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d. (%v)", http.StatusOK, status, s)
	}
}

func TestGetAliasesAfterUpdate(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/aliase", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAliases)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var alias []Alias
	if err := json.NewDecoder(rr.Body).Decode(&alias); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(alias) != 1 {
		t.Errorf("Dont return alias list")
	}

	if *alias[0].SourceUsername != "updated" {
		t.Errorf("Wrong SourceUsername got %v excepted %v", *alias[0].SourceUsername, "updated")
	}

	if *alias[0].SourceDomain != "alias.com" {
		t.Errorf("Wrong SourceDomain")
	}

	if *alias[0].DestinationUsername != "foo" {
		t.Errorf("Wrong SourceDomain")
	}

	if *alias[0].DestinationDomain != "google.com" {
		t.Errorf("Wrong SourceDomain")
	}

	if alias[0].Enabled != true {
		t.Errorf("Enabled is wrong")
	}
}

func TestRemoveAlias(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/alias", bytes.NewBufferString("{\"id\": 2}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteAlias)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		b, _ := ioutil.ReadAll(rr.Body)
		s := string(b)
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d. (%v)", http.StatusOK, status, s)
	}
}

func TestGetAliasesEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/aliase", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAliases)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var alias []Alias
	if err := json.NewDecoder(rr.Body).Decode(&alias); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(alias) != 0 {
		t.Errorf("Dont return empty alias list")
	}
}
