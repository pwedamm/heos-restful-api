package handler

import (
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func PlayPresetHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["pid"]
	params["preset"] = vars["preset"]

	cmd := internal.Command{
		Group:   "browse",
		Command: "play_preset",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}
