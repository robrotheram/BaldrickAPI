package api_handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robrotheram/baldrick_engine/app/db"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
)

var members = []db.Bot{}

func CreateBotHandlers(base string, router *mux.Router) *mux.Router  {
	sub := router.PathPrefix(base).Subrouter()
	sub.Methods("GET").Path("/bots").HandlerFunc(getBots)
	sub.Methods("GET").Path("/bot/{botid}").HandlerFunc(getBot)
	sub.Methods("POST").Path("/bot").HandlerFunc(addBot)
	return sub;
}

// swagger:route GET /bots test
//
// Get list of something
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//    Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       500: errorResponse
//       200: Bot
func getBots(w http.ResponseWriter, _ *http.Request) {
	// Query All
	c := db.Session.DB("YOUR_DB").C("BOT")
	var results []db.Bot;
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}

}

// swagger:route GET /bot/{BotID} listParams
//
// Get list of something
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       500: errorResponse
//       200: Bot

func getBot(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	botid := vars["botid"]

	c := db.Session.DB("YOUR_DB").C("BOT")
	var results []db.Bot;
	fmt.Println(botid)
	err := c.Find(bson.M{"name":botid}).All(&results)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}
// swagger:route POST  /bot/ testthing
//
// Get list of something
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       500: errorResponse
//       200: Bot
func addBot(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m db.Bot
	b, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(b, &m)
	members = append(members, m)
	j, _ := json.Marshal(members)
	w.Write(j)
}
