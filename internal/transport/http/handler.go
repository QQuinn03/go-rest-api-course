package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	//Router  http.ServeMux
	Service CommentService

	Server *http.Server
}

//type commontService interface{}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()

	// create a listener , If only :{port} is used as addr argument then that means
	//HTTP server is reachable from all the ip addresses
	//(loopback, public ip, internal ip) of the machine.
	h.Server = &http.Server{ //customirzed http server
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	h.mapRoute()

	return h

}
func (h *Handler) mapRoute() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world lala")
	})

	h.Router.HandleFunc("api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("api/v1/comment/", h.PostComment).Methods("POST")
	h.Router.HandleFunc("api/v1/comment/{id}", h.DeleteComment).Methods("Delete")
	h.Router.HandleFunc("api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
}

/* listen and serve starts a http server and return error */
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("Shutting down gracefully")
	return nil

}
