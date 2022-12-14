package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

func GetPlayerHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "player",
		Command: "get_players",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func GetPlayerInfoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_player_info",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func GetPlayStateHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_play_state",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// SetPlayStateHandler accepts play, pause and stop as params for state
func SetPlayStateHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["state"] = vars["state"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_play_state",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func GetNowPlayingMediaHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_now_playing_media",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func GetVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_volume",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// SetVolumeHandler level 0 to 100
func SetVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["level"] = vars["level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_volume",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// VolumeUpHandler level 0 to 10
func VolumeUpHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["step_level"] = vars["step_level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "volume_up",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// VolumeDownHandler level 0 to 10
func VolumeDownHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["step_level"] = vars["step_level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "volume_down",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// GetMuteHandler level 0 to 10
func GetMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_mute",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// SetMuteHandler state => on / off
func SetMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["state"] = vars["state"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_mute",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func ToggleMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "toggle_mute",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func GetPlayModeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_play_mode",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// SetPlayModeHandler repeat => on_all, on_one, off || shuffle => on, off
func SetPlayModeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["repeat"] = vars["repeat"]
	params["shuffle"] = vars["shuffle"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_play_mode",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func PlayNextHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "play_next",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func PlayPreviousHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "play_previous",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

// sendPlayerGroupMessageToHeosSystem takes input params from handler, writes message to heos and write the result to http.response
func sendPlayerGroupMessageToHeosSystem(w http.ResponseWriter, cmd internal.Command, params map[string]string) {
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

	var msg groups.HeosPlayerAnswer

	json.Unmarshal(response, &msg)

	log.Infof("logging raw answer %s", response)
	json.NewEncoder(w).Encode(msg)

}
