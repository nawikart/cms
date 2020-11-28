package main

import (
	"fmt"
	// "html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nawikart/cms/ctx/crudCtx"
	"github.com/nawikart/cms/ctx/formGeneratorCtx"
	"github.com/nawikart/cms/routes"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/form-generator/api", formGeneratorCtx.Api)
	r.HandleFunc("/crud/api", crudCtx.Api)

	routes.AuthRoutes(r)
	routes.DashboardRoutes(r)
	routes.MediaRoutes(r)
	routes.NuxtjsRoutes(r)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", r)
}
