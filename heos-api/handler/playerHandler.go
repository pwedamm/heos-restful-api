package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"heos-restful-api/heos-api/groups"
	"heos-restful-api/heos-api/internal"
	"heos-restful-api/heos-api/logger"
	"net/http"
)

func GetPlayerHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)
	var playPreviousResponse groups.HeosPlayerAnswer

	cmd := internal.Command{
		Group:   "player",
		Command: "get_players",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func GetPlayerInfoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_player_info",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func GetPlayStateHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_play_state",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// SetPlayStateHandler accepts play, pause and stop as params for state
func SetPlayStateHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["state"] = vars["state"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_play_state",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func GetNowPlayingMediaHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_now_playing_media",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func GetVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_volume",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// SetVolumeHandler level 0 to 100
func SetVolumeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["level"] = vars["level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_volume",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// VolumeUpHandler level 0 to 10
func VolumeUpHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["step_level"] = vars["step_level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "volume_up",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// VolumeDownHandler level 0 to 10
func VolumeDownHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["step_level"] = vars["step_level"]

	cmd := internal.Command{
		Group:   "player",
		Command: "volume_down",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// GetMuteHandler level 0 to 10
func GetMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_mute",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// SetMuteHandler state => on / off
func SetMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["state"] = vars["state"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_mute",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func ToggleMuteHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "toggle_mute",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func GetPlayModeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "get_play_mode",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// SetPlayModeHandler repeat => on_all, on_one, off || shuffle => on, off
func SetPlayModeHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]
	params["repeat"] = vars["repeat"]
	params["shuffle"] = vars["shuffle"]

	cmd := internal.Command{
		Group:   "player",
		Command: "set_play_mode",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func PlayNextHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "play_next",
	}
	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

func PlayPreviousHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	var playPreviousResponse groups.HeosPlayerAnswer
	params := make(map[string]string)
	params["pid"] = vars["player_id"]

	cmd := internal.Command{
		Group:   "player",
		Command: "play_previous",
	}

	res := sendPlayerGroupMessageToHeosSystem(w, cmd, params)
	if res != nil {

		json.Unmarshal(res, &playPreviousResponse)
		json.NewEncoder(w).Encode(playPreviousResponse)
	}
}

// sendPlayerGroupMessageToHeosSystem takes input params from handler, writes message to heos and write the result to http.response
func sendPlayerGroupMessageToHeosSystem(w http.ResponseWriter, cmd internal.Command, params map[string]string) []byte {
	heos, err := GetConnectedHeos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No connected devices. Error %v:", err.Error())
		return nil
	}

	response, err := heos.SendCmdToHeosSpeaker(cmd, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while communicating with heos %v", err.Error())
		return nil
	}
	logger.HeosLogger.Debugf("logging raw answer %s", response)
	return response
}
