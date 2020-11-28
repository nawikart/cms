package routes

import (
	"github.com/gorilla/mux"
	"github.com/nawikart/cms/ctx/dashboardCtx"
)

func DashboardRoutes(r *mux.Router) {
	r.HandleFunc("/dashboard-hotelcalav", dashboardCtx.Api)
}
