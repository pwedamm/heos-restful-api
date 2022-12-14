package handler

import (
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

	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
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
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func AccountSignOutHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "system",
		Command: "sign_out",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}

func HeartBeattHandler(w http.ResponseWriter, req *http.Request) {

	params := make(map[string]string)

	cmd := internal.Command{
		Group:   "system",
		Command: "heart_beat",
	}
	sendPlayerGroupMessageToHeosSystem(w, cmd, params)
}
