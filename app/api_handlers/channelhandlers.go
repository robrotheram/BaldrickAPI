package api_handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robrotheram/baldrick_engine/app/db"
	"github.com/robrotheram/baldrick_engine/app/messages"
	"io/ioutil"
	"time"
	"gopkg.in/mgo.v2/bson"
)

var channels = []db.Channel{}

func CreateChanelHandlers(base string, router *mux.Router) *mux.Router  {
	sub := router.PathPrefix(base).Subrouter()
	sub.Methods("GET").Path("/channel").HandlerFunc(getChannels)
	sub.Methods("GET").Path("/channel/{channelid}").HandlerFunc(getChannel)
	sub.Methods("POST").Path("/channel").HandlerFunc(addChannel)
	return sub;
}

// swagger:route GET /channel test
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
func getChannels(w http.ResponseWriter, _ *http.Request) {
	// Query All
	c := db.Session.DB("YOUR_DB").C("Channels")
	var results []db.Channel;
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}

}

// swagger:route GET /channel/{ChannelID} listParams
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

func getChannel(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	botid := vars["channelid"]

	c := db.Session.DB("YOUR_DB").C("Channels")
	var results []db.Channel;
	fmt.Println(botid)
	err := c.Find(bson.M{"name":botid}).All(&results)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}
// swagger:route POST  /channel/ testthing
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
func addChannel(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m db.Channel
	b, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(b, &m)
	m.StartTime = time.Unix(time.Now().Unix(), 0).String()
	messages.CreateChannel(m)
}
