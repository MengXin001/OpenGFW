package api

import (
    "encoding/json"
    "net/http"

	"github.com/apernet/OpenGFW/ruleset"
)
var rs ruleset.Ruleset

func LoadRules(ruleset ruleset.Ruleset) {
	// Ruleset
	rs = ruleset
}

func GetRulesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}