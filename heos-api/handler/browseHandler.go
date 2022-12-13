package handler

import (
	"encoding/json"
	"fmt"
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
	sendBrowseMessage(w, cmd, params)
}

func sendBrowseMessage(w http.ResponseWriter, cmd internal.Command, params map[string]string) {
	heos, err := GetConnectedHeos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No connected devices. Error %v:", err.Error())
		return
	}

	response, err := heos.SendCmdToHeosSpeaker(cmd, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while communicating with heos %v", err.Error())
		return
	}

	log.Infof("got answer %s", response)
	json.NewEncoder(w).Encode(response)
}
