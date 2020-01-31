package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/unrolled/render.v1"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	_ "github.com/go-sql-driver/mysql"
)

// Domain from MYSQL
type Domain struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
}

var db *sql.DB

func connectToDb() {
	var err error
	db, err = sql.Open("mysql", "root:example@tcp(127.0.0.1:3306)/vmail")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping DB")

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("Ping OK")
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

func defineRouten() chi.Router {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Get("/api/v1/domain", getDomains)
	r.Post("/api/v1/domain", addDomain)
	r.Delete("/api/v1/domain", deleteDomain)

	return r
}

func main() {
	//defer db.Close()
	fmt.Println("Hey")
	connectToDb()
	router := defineRouten()
	http.ListenAndServe(":3001", router)
	fmt.Println("Done")
}
