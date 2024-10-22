package api

import (
	"net/http"
)

func CreateServer() {
	// System
    http.HandleFunc("/api/info", GetSystemInfoHandler)
	// Rules
	http.HandleFunc("/api/rules", GetRulesHandler)
	// LogsTodo
}