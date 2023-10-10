package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment:" %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.ServerConfig.Env, version)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintln(w, "I'm alive!")
	// fmt.Fprintf(w, "environment: %s\n", app.config.ServerConfig.Env)
	// fmt.Fprintf(w, "version: %s\n", version)
	w.Write([]byte(js))
}
