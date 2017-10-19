// Package main Baldrick API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: 192.168.99.100:8080
//     BasePath: /api/v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Robrotheram<robrotheram@gmail.com> https://robrotheram.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//
//
//     Extensions:
//     ---
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//     ---
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"github.com/robrotheram/baldrick_engine/app/messages"
	"github.com/robrotheram/baldrick_engine/app/api_handlers"
	"github.com/robrotheram/baldrick_engine/app/configuration"
	"github.com/robrotheram/baldrick_engine/app/db"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"

)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "0.0.0.0:8080", "Address to listen for HTTP requests on")
)

func createRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	api_handlers.CreateBotHandlers("/api/v1",router)
	api_handlers.CreateRuleHandlers("/api/v1",router)
	api_handlers.CreateUserHandlers("/api/v1",router)
	router.HandleFunc("/", api_handlers.TestHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))



	return router
}

func main() {
	// Parse the command-line flags.
	flag.Parse()
	configuration.ReadConfig("config.json")
	db.CreateSession();
	db.GenerateBot();

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	//StartDispatcher(*NWorkers)
	messages.CreateListener("hello")
	messages.CreateListener("test1")
	messages.CreateListener("test2")
	messages.CreateListener("test3")
	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	//http.HandleFunc("/work", Collector)
	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)


	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(createRouters())); err != nil {
		fmt.Println(err.Error())
	}
	db.Close();
}