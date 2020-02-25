package main

import (
	"database/sql"
	"fmt"
	"github.com/99designs/basicauth-go"
	"github.com/go-chi/chi/middleware"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	_ "./statik"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectToDb() {
	log.Println("Try to connect to Database")
	dbString := getConfigVariable("DB")
	if dbString == "" {
		log.Fatal("No DB Connection string set")
	}
	var err error
	db, err = sql.Open("mysql", dbString)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ping Database")

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println("Connection to Database ok")
}

func getConfigVariable(name string) string {
	value := os.Getenv("GOMAILADMIN_"+name)
	return value
}

func http_ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func http_status(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Ok"))
}


func defineRouter() chi.Router {
	log.Println("Setup API-Routen")
	r := chi.NewRouter()

	//r.Use(basicauth.NewFromEnv("Need Auth", "GOMAILADMIN_USER"))
	apiKey := getConfigVariable("APIKEY")
	apiSecret := getConfigVariable("APISECRET")

	if apiKey != "" && apiSecret != "" {
		r.Use(basicauth.New("MyRealm", map[string][]string{
			apiKey: {apiSecret},
		}))
		log.Println("Enabled Basic auth for basic protection.")
	} else {
		log.Println("Run without Basic auth, make sure to protect the API at another layer")
	}

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

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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
	r.Get("/ping", http_ping)
	r.Get("/status", http_status)

	//Static files
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/*", http.StripPrefix("", http.FileServer(statikFS)))

	return r
}

func main() {
	log.Println("Start Go Mail Admin")
	connectToDb()
	router := defineRouter()
	http.ListenAndServe(":3001", router)
	fmt.Println("Done")
}
