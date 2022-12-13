package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func AccountCheckHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	cmd := internal.Command{
		Group:   "system",
		Command: "check_account",
	}

	sendSystemMessage(w, cmd, params)
}

func AccountSignInHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["un"] = vars["email"]
	params["pw"] = vars["password"]

	cmd := internal.Command{
		Group:   "system",
		Command: "sign_in",
	}
	sendSystemMessage(w, cmd, params)
}

func AccountSignOutHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "system",
		Command: "sign_out",
	}
	sendSystemMessage(w, cmd, params)
}

func HeartBeattHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "system",
		Command: "heart_beat",
	}
	sendSystemMessage(w, cmd, params)
}

func sendSystemMessage(w http.ResponseWriter, cmd internal.Command, params map[string]string) {
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
