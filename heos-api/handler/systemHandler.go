package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func AccountCheckHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	var playPreviousResponse groups.HeosPlayerAnswer
	cmd := internal.Command{
		Group:   "system",
		Command: "check_account",
	}

	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func AccountSignInHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["un"] = vars["email"]
	params["pw"] = vars["password"]

	cmd := internal.Command{
		Group:   "system",
		Command: "sign_in",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func AccountSignOutHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	var playPreviousResponse groups.HeosPlayerAnswer
	cmd := internal.Command{
		Group:   "system",
		Command: "sign_out",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func HeartBeattHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	var playPreviousResponse groups.HeosPlayerAnswer
	cmd := internal.Command{
		Group:   "system",
		Command: "heart_beat",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}
