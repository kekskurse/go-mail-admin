package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmptyDomainList(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/domain", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	resetDBForTest()
	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(m.getDomains)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var domains []Domain
	if err := json.NewDecoder(rr.Body).Decode(&domains); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(domains) != 0 {
		t.Errorf("Dont return empty domain list")
	}
}

func TestAddInvalidDomain(t *testing.T) {
	t.Skip("Domain validation need to be added")
	req, err := http.NewRequest("POST", "/v1/domain", bytes.NewBufferString("{\"domain\":\"invalide\"}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addDomain)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestAddDomain(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/domain", bytes.NewBufferString("{\"domain\":\"example.com\"}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addDomain)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestGetOneDomainList(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/domain", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(m.getDomains)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var domains []Domain
	if err := json.NewDecoder(rr.Body).Decode(&domains); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(domains) != 1 {
		t.Errorf("Dont return domain list")
	}

	if domains[0].Domain != "example.com" {
		t.Errorf("Wrong domain")
	}

	if domains[0].ID != 1 {
		t.Errorf("Wrong domain id")
	}
}

func TestDeleteNotExistingDomain(t *testing.T) {
	t.Skip("Check if domian existed must be added to controller first")
	req, err := http.NewRequest("DELETE", "/v1/domain", bytes.NewBufferString("{\"domain\":\"delete.com\"}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteDomain)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestDeleteDomain(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/domain", bytes.NewBufferString("{\"domain\":\"example.com\"}"))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteDomain)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestGetEmptyDomainList2(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/domain", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	m := NewMailServerConfiguratorInterface(NewConfig())
	m.connectToDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(m.getDomains)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var domains []Domain
	if err := json.NewDecoder(rr.Body).Decode(&domains); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(domains) != 0 {
		t.Errorf("Dont return empty domain list")
	}
}
