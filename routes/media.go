package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nawikart/cms/ctx/mediaCtx"
)

func MediaRoutes(r *mux.Router) {
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))
	r.HandleFunc("/media-upload", mediaCtx.Upload)
	r.HandleFunc("/media-api", mediaCtx.Api)
}
