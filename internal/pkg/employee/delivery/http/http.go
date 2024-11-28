package http

import (
	"employees/internal/models"
	"employees/internal/pkg/employee"
	"employees/internal/pkg/utils"
	"employees/internal/pkg/utils/messages"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
	"strconv"
)

type Params struct {
	fx.In

	Uc     employee.Usecase
	Logger *slog.Logger
}

type Handler struct {
	uc  employee.Usecase
	log *slog.Logger
}

func New(p Params) *Handler {
	return &Handler{
		uc:  p.Uc,
		log: p.Logger,
	}
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeData *models.Employee

	if err := utils.ReadRequestData(r, &employeeData); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	id, err := h.uc.CreateEmployee(r.Context(), employeeData)
	if err != nil {
		h.log.Error("create employee", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	h.log.Info("created new employee", "id", id)
	utils.Send201(w, id)
}

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeData *models.Employee

	if err := utils.ReadRequestData(r, &employeeData); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	employeeData.ID = int32(id)
	err = h.uc.EditEmployee(r.Context(), employeeData)
	if err != nil {
		h.log.Error("edit employee", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	h.log.Info("updated employee")
	utils.Send200(w, "employee updated")

}

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	if err = h.uc.DeleteEmployee(r.Context(), int32(id)); err != nil {
		h.log.Error("delete employee", "error", err.Error())
		utils.Send404(w, messages.NotFound)
	}

	h.log.Info("deleted employee")
	utils.Send200(w, "employee deleted")
}

func (h *Handler) GetCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	listEmployees, err := h.uc.GetListCompanyEmployees(r.Context(), int32(id))
	if err != nil {
		h.log.Error("get list of employees", "error", err.Error())
		utils.Send404(w, messages.NotFound)
	}

	h.log.Info("got list of company employees")
	utils.Send200(w, listEmployees)
}

func (h *Handler) GetDepartmentCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	departmentIDStr := vars["id"]
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		h.log.Error("parse department id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	listEmployees, err := h.uc.GetListDepartmentCompanyEmployees(r.Context(), int32(departmentID))
	if err != nil {
		h.log.Error("get list of department employees", "error", err.Error())
		utils.Send404(w, messages.NotFound)
	}

	h.log.Info("got list of department employees")
	utils.Send200(w, listEmployees)
}

func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var name string
	if err := utils.ReadRequestData(r, &name); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	companyID, err := h.uc.CreateCompany(r.Context(), name)
	if err != nil {
		h.log.Error("create company", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	h.log.Info("created company", "id", companyID)
	utils.Send201(w, companyID)
}

func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var department *models.Department
	if err := utils.ReadRequestData(r, &department); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	departmentID, err := h.uc.CreateDepartment(r.Context(), department)
	if err != nil {
		h.log.Error("create department", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
	}

	h.log.Info("created department", "id", departmentID)
	utils.Send201(w, departmentID)
}
