package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func GetGroupsHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "group",
		Command: "get_groups",
	}
	sendGroupGroupMessageToHeosSystem(w, cmd, params)
}

func GetGroupInfoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["gid"] = vars["gid"]

	cmd := internal.Command{
		Group:   "group",
		Command: "get_group_info",
	}
	sendGroupGroupMessageToHeosSystem(w, cmd, params)
}

func GetGroupVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["gid"] = vars["gid"]

	cmd := internal.Command{
		Group:   "group",
		Command: "get_volume",
	}
	sendGroupGroupMessageToHeosSystem(w, cmd, params)
}

func SetGroupVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["gid"] = vars["gid"]
	params["level"] = vars["level"]

	cmd := internal.Command{
		Group:   "group",
		Command: "set_volume",
	}
	sendGroupGroupMessageToHeosSystem(w, cmd, params)
}

// SetGroupHandler pid => List of comma separated player_id's where each player id is returned by 'get_players' or 'get_groups' command; first
//
//	player_id in list is group leader
func SetGroupHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["p_ids"]

	cmd := internal.Command{
		Group:   "group",
		Command: "set_group",
	}
	sendGroupGroupMessageToHeosSystem(w, cmd, params)
}

// sendPlayerGroupMessageToHeosSystem takes input params from handler, writes message to heos and write the result to http.response
func sendGroupGroupMessageToHeosSystem(w http.ResponseWriter, cmd internal.Command, params map[string]string) {
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

	var msg groups.HeosGroupAnswer

	json.Unmarshal(response, &msg)

	log.Infof("logging raw answer %s", response)
	json.NewEncoder(w).Encode(msg)

}
