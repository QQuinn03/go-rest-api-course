package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//handler -store pointer to our comments service
type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRouter() {
	fmt.Println("Set up router")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
