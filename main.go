package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	_ "github.com/go-sql-driver/mysql"
)

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

func defineRouten() chi.Router { //TODO: einheitliche Sprache
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
	r.Get("/api/v1/alias", getAliases)
	r.Post("/api/v1/alias", addAlias)
	r.Delete("/api/v1/alias", deleteAlias)
	r.Put("/api/v1/alias", updateAlias)
	r.Get("/api/v1/account", getAccounts)
	r.Post("/api/v1/account", addAccount)
	r.Delete("/api/v1/account", deleteAccount)
	r.Put("/api/v1/account", updateAccount)
	r.Put("/api/v1/account/password", updateAccountPassword)
	r.Get("/api/v1/tlspolicy", getTLSPolicy)
	r.Post("/api/v1/tlspolicy", addTLSPolicy)
	r.Put("/api/v1/tlspolicy", updateTLSPolicy)
	r.Delete("/api/v1/tlspolicy", deleteTLSPolicy)

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
