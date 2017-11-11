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
	"os"
	"errors"
	"time"
	"log"
	"github.com/robrotheram/baldrick_engine/app/db"
	"github.com/robrotheram/baldrick_engine/app/rules"
	"github.com/robrotheram/baldrick_engine/app/messages"
	"github.com/robrotheram/baldrick_engine/app/api_handlers"
	"github.com/robrotheram/baldrick_engine/app/configuration"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"strings"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"

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
	api_handlers.CreateChanelHandlers("/api/v1",router)

	api_handlers.CreateAuthHandlers("/api/v1",router)


	router.HandleFunc("/", api_handlers.TestHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return router
}

func Run(function string, args... interface{})(x bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Invalid rule")
		}
	}()
	rules.Invoke(function, args...)
	x = true
	return x,err
}

func test(){
	arg := os.Args[1]
	a := append(os.Args[:0], os.Args[2:]...)
	a = a[:len(a) - 1]
	s := make([]interface{}, len(a))
	for i, v := range a {
		s[i] = v
	}
	_,  err := Run(arg,s...)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ran complete !!!!")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// Parse the command-line flags.
	test()
	testJWE();

	flag.Parse()
	configuration.ReadConfig("config.json")
	db.CreateSession();
	db.GenerateBot();

	messages.OpenConnect();

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	//StartDispatcher(*NWorkers)

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
func forever() {
	for {
		//fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}

func findARule(rules []db.Rule, msg string) db.Rule {
	result := strings.Split(msg, " ")
	for i := range rules {
		if (result[0] == rules[i].Prifix){
			return rules[i]
		}
	}
	return db.Rule{};
}