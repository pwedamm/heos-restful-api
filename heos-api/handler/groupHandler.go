package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func GetGroupsHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	var setGroupResponse groups.HeosGroupAnswer
	cmd := internal.Command{
		Group:   "group",
		Command: "get_groups",
	}
	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &setGroupResponse)
		json.NewEncoder(w).Encode(setGroupResponse)
	}
}

func GetGroupInfoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var setGroupResponse groups.HeosGroupAnswer
	params := make(map[string]string)
	params["gid"] = vars["gid"]

	cmd := internal.Command{
		Group:   "group",
		Command: "get_group_info",
	}
	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &setGroupResponse)
		json.NewEncoder(w).Encode(setGroupResponse)
	}
}

func GetGroupVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var setGroupResponse groups.HeosGroupAnswer
	params := make(map[string]string)
	params["gid"] = vars["gid"]

	cmd := internal.Command{
		Group:   "group",
		Command: "get_volume",
	}
	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &setGroupResponse)
		json.NewEncoder(w).Encode(setGroupResponse)
	}
}

func SetGroupVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var setGroupResponse groups.HeosGroupAnswer
	params := make(map[string]string)
	params["gid"] = vars["gid"]
	params["level"] = vars["level"]

	cmd := internal.Command{
		Group:   "group",
		Command: "set_volume",
	}
	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &setGroupResponse)
		json.NewEncoder(w).Encode(setGroupResponse)
	}
}

// SetGroupHandler pid => List of comma separated player_id's where each player id is returned by 'get_players' or 'get_groups' command; first
//
//	player_id in list is group leader
func SetGroupHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var setGroupResponse groups.HeosGroupAnswer
	params := make(map[string]string)
	params["pid"] = vars["p_ids"]

	cmd := internal.Command{
		Group:   "group",
		Command: "set_group",
	}

	res := sendMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &setGroupResponse)
		json.NewEncoder(w).Encode(setGroupResponse)
	}
}
