package main

import (
	"strconv"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
)

type Config struct {
	PasswordScheme string
	Port           uint16
	Address        string
	V3Config       bool
	AuthUsername   string
	AuthPassword   string
	FeatureToggles struct {
		CatchAll          bool   `json:"catchall"`
		AuthMethode       string `json:"auth"`
		ShowDomainRecords bool   `json:"showDomainDetails"`
	}
	DatabaseURI           string
	CheckDnsRecords       bool
	RedisNetwork          string
	RedisAddress          string
	AuthMethod            string
	ApiKey                string
	ApiSecret             string
	HttpBasicAuthUsername string
	HttpBasicAuthPassword string

	DkimSelector string
	DkimValue    string
}

func NewConfig() Config {
	config := Config{}
	config.PasswordScheme = getConfigVariableWithDefault("PASSWORD_SCHEME", "SSHA512")
	port, err := strconv.ParseUint(getConfigVariableWithDefault("PORT", "3001"), 10, 16)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid port")
	}
	config.Port = uint16(port)
	config.Address = getConfigVariableWithDefault("ADDRESS", "localhost")

	config.V3Config = false
	if getConfigVariableWithDefault("V3", "off") == "on" {
		config.V3Config = true
		config.AuthUsername = getConfigVariable("AUTH_Username")

		if config.AuthUsername == "" {
			log.Fatal().Msgf("No Username is set, set GOMAILADMIN_AUTH_Username")
		}
		config.AuthPassword = getConfigVariable("AUTH_Password")
		if config.AuthPassword == "" {
			log.Fatal().Msgf("No Password is set, set GOMAILADMIN_AUTH_Password")
		}
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
		log.Error().Msgf("Use static secret, dont to this!")
	} else {
		config.FeatureToggles.AuthMethode = authConfig.Method
	}

	if getConfigVariableWithDefault("CATCHALL", "On") == "On" {
		config.FeatureToggles.CatchAll = true
	}

	if getConfigVariableWithDefault("SHOW_DNS_RECORDS", "On") == "On" {
		config.FeatureToggles.ShowDomainRecords = true
	}

	config.DatabaseURI = getConfigVariable("DB")
	if config.DatabaseURI == "" {
		log.Fatal().Msg("No DB Connection string set, set enviroment varieable GOMAILADMIN_DB")
	}

	if getConfigVariable("CHECK_DNS_RECORDS") == "On" {
		config.CheckDnsRecords = true
	}

	config.RedisNetwork = getConfigVariableWithDefault("REDIS_NETWORK", "tcp")
	config.RedisAddress = getConfigVariableWithDefault("REDIS_ADDRESS", "localhost:6379")

	config.AuthMethod = getConfigVariableWithDefault("AUTH_METHOD", "")
	config.ApiKey = getConfigVariable("APIKEY")
	config.ApiSecret = getConfigVariable("APISECRET")

	config.HttpBasicAuthUsername = getConfigVariable("AUTH_HTTPBasic_Username")
	config.HttpBasicAuthPassword = getConfigVariable("AUTH_HTTPBasic_Password")

	if config.ApiKey != "" || config.ApiSecret != "" {
		log.Warn().Msg("The old Enviroment variable GOMAILADMIN_APIKEY or GOMAILADMIN_APISECRET are set. The Auth method is forced to HTTPBasicAuth, please read the new auth docs to change your configuration")
	}

	config.DkimSelector = getConfigVariable("DKIM_SELECTOR")
	config.DkimValue = getConfigVariable("DKIM_VALUE")

	return config
}
