package main

import (
	"database/sql"
	"embed"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rakyll/statik/fs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	_ "github.com/go-sql-driver/mysql"
	_ "gomailadmin/mailserver-configurator-interface/statik"
)

var (
	version                                = "development"
)

var db *sql.DB

//go:embed templates/*.tmpl
var embeddedTemplates embed.FS

func init() {
	// Config logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Init")

	//Init Database
	connectToDb()

	//Config Auth
	if getConfigVariableWithDefault("V3", "off") == "on" {
		checkAuthConfig()
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
		log.Error().Msgf("Use static secret, dont to this!")
	}
}

func checkAuthConfig() {
	if getConfigVariable("AUTH_Username") == "" {
		log.Fatal().Msgf("No Username is set, set GOMAILADMIN_AUTH_Username")
	}

	if getConfigVariable("AUTH_Password") == "" {
		log.Fatal().Msgf("No Password is set, set GOMAILADMIN_AUTH_Password")
	}
}

func connectToDb() {
	log.Debug().Msg("Try to connect to Database")
	dbString := getConfigVariable("DB")
	if dbString == "" {
		log.Fatal().Msg("No DB Connection string set, set enviroment varieable GOMAILADMIN_DB")
	}
	var err error
	db, err = sql.Open("mysql", dbString)

	if err != nil {
		log.Fatal().Err(err).Msg("Can`t connect to db")
	}

	log.Debug().Msg("Ping Database")

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Debug().Msg("Connection to Database ok")
}

func getConfigVariable(name string) string {
	value := os.Getenv("GOMAILADMIN_" + name)
	return value
}

func getConfigVariableWithDefault(name string, defaultValue string) string {
	value := os.Getenv("GOMAILADMIN_" + name)
	if value == "" {
		return defaultValue
	}
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

var authConfig auth

func defineRouter() chi.Router {
	log.Debug().Msg("Setup API-Routen")
	r := chi.NewRouter()

	if getConfigVariableWithDefault("V3", "off") == "on" {
		log.Info().Msgf("Run with v3 config")
	} else {
		redis := newRedisConnection()
		authConfig = NewAuthFromEnv(redis)
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
	if getConfigVariableWithDefault("V3", "off") == "on" {
		apiRouten.Use(jwtauth.Verifier(tokenAuth))
		apiRouten.Use(jwtauth.Authenticator)
	} else {
		apiRouten.Use(authConfig.Handle)
		apiRouten.Post("/v1/logout", logout)
	}


	apiRouten.Get("/v1/domain", getDomains)
	apiRouten.Get("/v1/domain/{domain}", getDomainDetails)
	apiRouten.Post("/v1/domain", addDomain)
	apiRouten.Delete("/v1/domain", deleteDomain)
	apiRouten.Get("/v1/alias", getAliases)
	apiRouten.Post("/v1/alias", addAlias)
	apiRouten.Delete("/v1/alias", deleteAlias)
	apiRouten.Put("/v1/alias", updateAlias)
	apiRouten.Get("/v1/account", getAccounts)
	apiRouten.Post("/v1/account", addAccount)
	apiRouten.Delete("/v1/account", deleteAccount)
	apiRouten.Put("/v1/account", updateAccount)
	apiRouten.Put("/v1/account/password", updateAccountPassword)
	apiRouten.Get("/v1/tlspolicy", getTLSPolicy)
	apiRouten.Post("/v1/tlspolicy", addTLSPolicy)
	apiRouten.Put("/v1/tlspolicy", updateTLSPolicy)
	apiRouten.Delete("/v1/tlspolicy", deleteTLSPolicy)
	apiRouten.Get("/v1/features", getFeatureToggles)
	apiRouten.Get("/v1/version", getVersion)
	r.Get("/ping", http_ping)
	r.Get("/status", http_status)
	//r.Get("/test", test)

	publicRouten := chi.NewRouter()

	if getConfigVariableWithDefault("V3", "off") == "on" {
		publicRouten.Post("/v1/login", login)
		publicRouten.Post("/v1/login/username", login) //Old route for old frontend, need to be removed
	} else {
		publicRouten.Post("/v1/login/username", loginUsername)
	}

	publicRouten.Post("/v1/features", getFeatureToggles)
	publicRouten.Get("/v1/features", getFeatureToggles)



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

	router := defineRouter()
	address := getConfigVariable("ADDRESS")
	port := getConfigVariable("PORT")
	if port == "" {
		port = "3001"
	}
	err := http.ListenAndServe(address+":"+port, router)
	log.Error().Err(err).Msg("HTTP Server stop")

	log.Debug().Msg("Done, Shotdown")
}



/*func test(w http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: embeddedTemplates,
		},
		Extensions: []string{".html", ".tmpl"},
	})

	r.HTML(w, http.StatusOK, "example", "world")
}*/