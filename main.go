package main

import (
	"net/http"

	"github.com/inconshreveable/log15"
)

func main() {

	log := log15.New()
	log.Debug("starting")
	mux := http.NewServeMux()
	mux.Handle("/standard", &EndpointHandler{log})
	mux.Handle("/closure", ClosureHandler(log)) // then here we call the outer function, "closing" around our logger.

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Error(
			"Server Exited",
			"err", err,
		)
	}

}

// This is the standard handler. We create a struct, add a method and pass it to the muxer.
// We have a single instance variable, log
type EndpointHandler struct {
	log log15.Logger
}

func (h *EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("Request Received - standard")
}

// end standard handler

// This is the closure style handler. We also have one "variable", but it is immutable.
func ClosureHandler(log log15.Logger) http.HandlerFunc {

	// this has the same behavior, but we define it as an "anonymous" (it has no name) function so we can close around it and still access the incoming parameters. (log in this case)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Request Received - closure")
	}
}

// end closure handler
