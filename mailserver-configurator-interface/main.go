package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rakyll/statik/fs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"gomailadmin/mailserver-configurator-interface/password"
	_ "gomailadmin/mailserver-configurator-interface/statik"

	_ "github.com/go-sql-driver/mysql"
)

var (
	version = "development"
)

func init() {
	// Config logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Init")
}

type MailServerConfiguratorInterface struct {
	DBConn              *sql.DB
	Config              Config
	PasswordHashBuilder password.PasswordHashBuilder
}

func NewMailServerConfiguratorInterface(config Config) *MailServerConfiguratorInterface {
	hb := password.GetPasswordHashBuilder(config.PasswordScheme)

	return &MailServerConfiguratorInterface{Config: config, PasswordHashBuilder: hb}
}

func (m *MailServerConfiguratorInterface) connectToDb() {
	log.Debug().Msg("Try to connect to Database")
	db, err := sql.Open("mysql", m.Config.DatabaseURI)

	if err != nil {
		log.Fatal().Err(err).Msg("Can`t connect to db")
	}
	m.DBConn = db

	log.Debug().Msg("Ping Database")

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Debug().Msg("Connection to Database ok")
}

func http_ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (m *MailServerConfiguratorInterface) http_status(w http.ResponseWriter, r *http.Request) {
	err := m.DBConn.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Ok"))
}

var authConfig auth

func defineRouter(config Config) chi.Router {
	log.Debug().Msg("Setup API-Routen")
	r := chi.NewRouter()

	m := NewMailServerConfiguratorInterface(config)
	m.connectToDb()

	if config.V3Config {
		log.Info().Msgf("Run with v3 config")
	} else {
		redis := newRedisConnection(config)
		authConfig = NewAuthFromEnv(redis, config)
	}

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-APITOKEN", "x-apitoken"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	apiRouten := chi.NewRouter()
	if config.V3Config {
		apiRouten.Use(jwtauth.Verifier(tokenAuth))
		apiRouten.Use(jwtauth.Authenticator)
	} else {
		apiRouten.Use(authConfig.Handle)
		apiRouten.Post("/v1/logout", logout)
	}

	apiRouten.Get("/v1/domain", m.getDomains)
	apiRouten.Get("/v1/domain/{domain}", m.getDomainDetails)
	apiRouten.Post("/v1/domain", m.addDomain)
	apiRouten.Delete("/v1/domain", m.deleteDomain)
	apiRouten.Get("/v1/alias", m.getAliases)
	apiRouten.Post("/v1/alias", m.addAlias)
	apiRouten.Delete("/v1/alias", m.deleteAlias)
	apiRouten.Put("/v1/alias", m.updateAlias)
	apiRouten.Get("/v1/account", m.getAccounts)
	apiRouten.Post("/v1/account", m.addAccount)
	apiRouten.Delete("/v1/account", m.deleteAccount)
	apiRouten.Put("/v1/account", m.updateAccount)
	apiRouten.Put("/v1/account/password", m.updateAccountPassword)
	apiRouten.Get("/v1/tlspolicy", m.getTLSPolicy)
	apiRouten.Post("/v1/tlspolicy", m.addTLSPolicy)
	apiRouten.Put("/v1/tlspolicy", m.updateTLSPolicy)
	apiRouten.Delete("/v1/tlspolicy", m.deleteTLSPolicy)
	apiRouten.Get("/v1/features", m.getFeatureToggles)
	apiRouten.Get("/v1/version", getVersion)
	r.Get("/ping", http_ping)
	r.Get("/status", m.http_status)
	//r.Get("/test", test)

	publicRouten := chi.NewRouter()

	if config.V3Config {
		publicRouten.Post("/v1/login", m.login)
		publicRouten.Post("/v1/login/username", m.login) //Old route for old frontend, need to be removed
	} else {
		publicRouten.Post("/v1/login/username", loginUsername)
	}

	publicRouten.Post("/v1/features", m.getFeatureToggles)
	publicRouten.Get("/v1/features", m.getFeatureToggles)

	r.Mount("/api", apiRouten)
	r.Mount("/public", publicRouten)

	//Static files
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err)
	}

	r.Handle("/*", http.StripPrefix("", http.FileServer(statikFS)))

	return r
}

func main() {

	log.Debug().Msg("Start Go Mail Admin")
	log.Info().Msgf("Running version %v", version)

	config := NewConfig()

	router := defineRouter(config)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), router)
	if err != nil {
		log.Error().Err(err).Msg("HTTP Server stop")
	}

	log.Debug().Msg("Done, Shotdown")
}
