package main

import (
	"fmt"
	"net/http"

	"yt_alexmux/internal/learnGoFast/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("Server is running on port 8000")
}
