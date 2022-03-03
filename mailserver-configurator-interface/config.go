package main

import (
	"fmt"
	"os"
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

func getConfigVariable(name string) string {
	value := os.Getenv("GOMAILADMIN_" + name)
	return value
}

func getConfigVariableWithDefault(name string, defaultValue string) string {
	value := getConfigVariable(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func parseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "off", "Off":
		return false, nil
	}
	return false, fmt.Errorf("unable to parse '%s' to boolean", str)
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

	v3Config, err := parseBool(getConfigVariableWithDefault("V3", "off"))
	if err != nil {
		log.Fatal().Err(err).Msgf("invalid value for GOMAILADMIN_V3")
	}
	config.V3Config = v3Config

	if config.V3Config {
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

	catchAll, err := parseBool(getConfigVariableWithDefault("CATCHALL", "On"))
	if err != nil {
		log.Fatal().Err(err).Msgf("invalid value for GOMAILADMIN_CATCHALL")
	}
	config.FeatureToggles.CatchAll = catchAll

	showDomainRecords, err := parseBool(getConfigVariableWithDefault("SHOW_DNS_RECORDS", "On"))
	if err != nil {
		log.Fatal().Err(err).Msgf("invalid value for GOMAILADMIN_SHOW_DNS_RECORDS")
	}
	config.FeatureToggles.ShowDomainRecords = showDomainRecords

	config.DatabaseURI = getConfigVariable("DB")
	if config.DatabaseURI == "" {
		log.Fatal().Msg("No DB Connection string set, set enviroment varieable GOMAILADMIN_DB")
	}

	checkDnsRecords, err := parseBool(getConfigVariableWithDefault("CHECK_DNS_RECORDS", "off"))
	if err != nil {
		log.Fatal().Err(err).Msgf("invalid value for GOMAILADMIN_CHECK_DNS_RECORDS")
	}
	config.CheckDnsRecords = checkDnsRecords

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
