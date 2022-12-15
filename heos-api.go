package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	heos_api "heos-restful-api/heos-api"
	"heos-restful-api/heos-api/handler"
	"heos-restful-api/heos-api/logger"
	"net/http"
	"time"
)

func main() {

	config, _ := heos_api.GetConfiguration()
	logger.HeosLogger.SetLevel(config.GetLogLever())
	handler.SetDefaultHeosIp(config.HeosIp)
	// creating mux handler https://github.com/gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello)

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// connection
	router.HandleFunc("/connect/{ip}", handler.ConnectHandler)

	// system group calls
	router.HandleFunc("/system/check_account", handler.AccountCheckHandler)
	router.HandleFunc("/system/sign_in/{email}/{password}", handler.AccountSignInHandler)
	router.HandleFunc("/system/sign_out", handler.AccountSignOutHandler)
	router.HandleFunc("/system/heart_beat", handler.HeartBeattHandler)

	// player group calls
	router.HandleFunc("/player/get_players", handler.GetPlayerHandler)
	router.HandleFunc("/player/get_player_info/{player_id}", handler.GetPlayerInfoHandler)
	router.HandleFunc("/player/get_player_state/{player_id}", handler.GetPlayStateHandler)
	// state => play,pause,stop
	router.HandleFunc("/player/set_player_state/{player_id}/{state}", handler.SetPlayStateHandler)
	router.HandleFunc("/player/get_now_playing_media/{player_id}", handler.GetNowPlayingMediaHandler)
	router.HandleFunc("/player/get_volume/{player_id}", handler.GetVolumeHandler)
	// level => 0 to 100
	router.HandleFunc("/player/set_volume/{player_id}/{level}", handler.SetVolumeHandler)
	// step_level => 0 to 10
	router.HandleFunc("/player/volume_up/{player_id}/{step_level}", handler.VolumeUpHandler)
	router.HandleFunc("/player/volume_down/{player_id}/{step_level}", handler.VolumeDownHandler)
	router.HandleFunc("/player/get_mute/{player_id}", handler.GetMuteHandler)
	// state => on, off
	router.HandleFunc("/player/set_mute/{player_id}/{state}", handler.SetMuteHandler)
	router.HandleFunc("/player/toggle_mute/{player_id}", handler.ToggleMuteHandler)
	router.HandleFunc("/player/get_play_mode/{player_id}", handler.GetPlayModeHandler)
	// repeat => on_all, on_one, off || shuffle => on, off
	router.HandleFunc("/player/set_play_mode/{player_id}/{repeat}/{shuffle}", handler.SetPlayModeHandler)
	// get queue
	// play queue item
	// remove items from queue
	// save queue as playlist
	// clear queue
	// move queue
	router.HandleFunc("/player/play_next/{player_id}", handler.PlayNextHandler)
	router.HandleFunc("/player/play_previous/{player_id}", handler.PlayPreviousHandler)

	// group commands
	router.HandleFunc("/group/get_groups", handler.GetGroupsHandler)
	router.HandleFunc("/group/get_group_info/{gid}", handler.GetGroupInfoHandler)
	// pid => List of comma separated player_id's where each player id is returned by 'get_players' or 'get_groups' command; first
	//	player_id in list is group leader
	router.HandleFunc("/group/set_group/{p_ids}", handler.SetGroupHandler)

	router.HandleFunc("/group/get_volume/{gid}", handler.GetGroupVolumeHandler)
	// level => 0 to 100
	router.HandleFunc("/group/set_volume/{gid}/{level}", handler.SetGroupVolumeHandler)
	// GroupVolumeUp
	// GroupVolumeDown
	// GroupGetGroupMute
	// SetGroupMute
	// ToggleGroupMute

	// browse
	// preset => preset number from app
	router.HandleFunc("/browse/play_preset/{pid}/{preset}", handler.PlayPresetHandler)

	server := &http.Server{
		Addr:         config.GetServerIp() + ":" + config.GetServerPort(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Handler:      router,
	}

	logger.HeosLogger.Infof("starting webserver on: %s.", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.HeosLogger.Error(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {

	logger.HeosLogger.Debug("entered hello handler")
	fmt.Fprintf(w, "hello\n")
}
