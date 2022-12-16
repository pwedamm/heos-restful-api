package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

// PlayPresetHandler method is used to play radio stations triggered by knx buttons.
func PlayPresetHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPresetHandlerResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["pid"]
	params["preset"] = vars["preset"]

	cmd := internal.Command{
		Group:   "browse",
		Command: "play_preset",
	}

	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPresetHandlerResponse)
		json.NewEncoder(w).Encode(playPresetHandlerResponse)
	}
}
