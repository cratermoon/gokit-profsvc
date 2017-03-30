package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"

	"github.com/cratermoon/gokit-profsvc/svcs"
	"github.com/gorilla/mux"
)

var author = expvar.NewString("author")
var authorContact = expvar.NewString("authorContact")

func main() {
	r := mux.NewRouter()
	svc.NewProfileService(r)

	r.Methods("GET").Path("/debug/vars").Handler(expvar.Handler())

	fmt.Println("All services ready")
	author.Set("Steven E. Newton")
	authorContact.Set("snewton@treetopllc.com")

	log.Fatal(http.ListenAndServe(":8080", r))
}
