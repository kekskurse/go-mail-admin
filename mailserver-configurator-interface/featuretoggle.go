package main

import (
	"github.com/unrolled/render"
	"net/http"
)

type featureToggles struct {
	CatchAll bool `json:"catchall"`
	AuthMethode string `json:"auth"`
	ShowDomainRecords bool `json:"showDomainDetails"`
}

func NewFeatureToggleFromEnv() *featureToggles {
	ft := featureToggles{}
	//Catchall
	ft.CatchAll = false
	ft.ShowDomainRecords = false
	if getConfigVariable("CATCHALL") == "On" {
		ft.CatchAll = true
	}

	if getConfigVariableWithDefault("SHOW_DNS_RECORDS", "On") == "On" {
		ft.ShowDomainRecords = true
	}


	ft.AuthMethode = authConfig.Method

	return &ft
}

func getFeatureToggles(w http.ResponseWriter, r *http.Request) {
	ft := NewFeatureToggleFromEnv()
	ren := render.New()
	ren.JSON(w, http.StatusOK, ft)
}