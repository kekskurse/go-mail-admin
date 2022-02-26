package main

import (
	"net/http"

	"github.com/unrolled/render"
)

func (m *MailServerConfiguratorInterface) getFeatureToggles(w http.ResponseWriter, r *http.Request) {
	ren := render.New()
	ren.JSON(w, http.StatusOK, m.Config.FeatureToggles)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	ren := render.New()
	res := struct {
		Version string `json:"version"`
	}{
		Version: version,
	}
	ren.JSON(w, http.StatusOK, res)
}
