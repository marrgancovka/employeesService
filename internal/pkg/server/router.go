package server

import (
	_ "employees/docs"
	handlerEmployee "employees/internal/pkg/employee/delivery/http"
	"employees/internal/pkg/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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
	v1.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	employees := v1.PathPrefix("/employees").Subrouter()

	employees.HandleFunc("", p.Handler.CreateEmployee).Methods(http.MethodPost)
	employees.HandleFunc("/{id}", p.Handler.DeleteEmployee).Methods(http.MethodDelete)
	employees.HandleFunc("/{id}", p.Handler.UpdateEmployee).Methods(http.MethodPatch)

	companies := v1.PathPrefix("/companies").Subrouter()

	companies.HandleFunc("/{id}/employees", p.Handler.GetCompanyEmployees).Methods(http.MethodGet)
	companies.HandleFunc("", p.Handler.CreateCompany).Methods(http.MethodPost)

	departments := v1.PathPrefix("/departments").Subrouter()

	departments.HandleFunc("/{id}/employees", p.Handler.GetDepartmentCompanyEmployees).Methods(http.MethodGet)
	departments.HandleFunc("", p.Handler.CreateDepartment).Methods(http.MethodPost)

	router := &Router{
		handler: api,
	}

	p.Logger.Info("registered router")

	return router
}
