package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jl-23929/battery-manager/assets"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	files := []string{
		"templates/base.gohtml",
		"templates/partials/nav.gohtml",
		"templates/partials/footer.gohtml",
		"templates/pages/home.gohtml",
	}

	ts, err := template.ParseFS(assets.EmbeddedFiles, files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func viewBattery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	batteryID := r.PathValue("id")

	// TODO:  Fetch battery details from DB

	files := []string{
		"templates/base.gohtml",
		"templates/partials/nav.gohtml",
		"templates/partials/footer.gohtml",
		"templates/pages/viewbattery.gohtml",
	}

	ts, err := template.ParseFS(assets.EmbeddedFiles, files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", batteryID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func newBattery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	files := []string{
		"templates/base.gohtml",
		"templates/partials/nav.gohtml",
		"templates/partials/footer.gohtml",
		"templates/pages/newbattery.gohtml",
	}

	ts, err := template.ParseFS(assets.EmbeddedFiles, files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	batteryID := id.String()

	png, err := generateQRCode("/viewbattery/" + batteryID)
	if err != nil {
		panic(err)
	}

	type batteryData struct {
		BatteryID string
		Image     string
	}

	bd := batteryData{
		BatteryID: batteryID,
		Image:     base64.StdEncoding.EncodeToString(png),
	}

	err = ts.ExecuteTemplate(w, "base", bd)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	errs := []error{}

	// TODO: Check services

	if len(errs) > 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		for _, err := range errs {
			fmt.Fprintf(w, "%s\n", err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
