package api_handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robrotheram/baldrick_engine/app/db"
	"io/ioutil"
)


func CreateRuleHandlers(base string, router *mux.Router) *mux.Router  {
	sub := router.PathPrefix(base).Subrouter()
	sub.Methods("GET").Path("/rules").HandlerFunc(getBots)
	sub.Methods("GET").Path("/rule/{botid}").HandlerFunc(getBot)
	sub.Methods("POST").Path("/rule").HandlerFunc(addBot)
	return sub;
}

func getRules(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello this is a test");
}

func getRule(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	botid := vars["botid"]
	fmt.Fprintf(w, "Hello this is a test bot id is:"+botid);
}

func addRule(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m db.Bot
	b, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(b, &m)
	members = append(members, m)
	j, _ := json.Marshal(members)
	w.Write(j)
}
