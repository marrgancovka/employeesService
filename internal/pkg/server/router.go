package server

import (
	handlerEmployee "employees/internal/pkg/employee/delivery/http"
	"employees/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
)

type RouterParams struct {
	fx.In

	Handler *handlerEmployee.Handler
	Logger  *slog.Logger
}

type Router struct {
	handler *mux.Router
}

func NewRouter(p RouterParams) *Router {
	api := mux.NewRouter().PathPrefix("/api").Subrouter()
	api.Use(middleware.CORSMiddleware)

	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/employees", p.Handler.CreateEmployee).Methods(http.MethodPost)
	v1.HandleFunc("/employees/{id}", p.Handler.DeleteEmployee).Methods(http.MethodDelete)
	v1.HandleFunc("/companies/{id}/employees", p.Handler.GetCompanyEmployees).Methods(http.MethodGet)
	v1.HandleFunc("/departments/{id}/employees", p.Handler.GetDepartmentCompanyEmployees).Methods(http.MethodGet)
	v1.HandleFunc("/employees/{id}", p.Handler.UpdateEmployee).Methods(http.MethodPatch)
	v1.HandleFunc("/companies", p.Handler.CreateCompany).Methods(http.MethodPost)
	v1.HandleFunc("/departments", p.Handler.CreateDepartment).Methods(http.MethodPost)

	router := &Router{
		handler: api,
	}

	p.Logger.Info("start server")

	return router
}