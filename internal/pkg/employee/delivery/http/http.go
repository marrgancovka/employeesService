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

// CreateEmployee godoc
// @Summary      Создать сотрудника
// @Description  Создать нового сотрудника
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        request body models.CreateEmployee true "employee data"
// @Success      201  {object} models.ResponseID
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /employees [post]
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeData *models.CreateEmployee

	if err := utils.ReadRequestData(r, &employeeData); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	h.log.Debug("create employee", "data", employeeData)
	id, err := h.uc.CreateEmployee(r.Context(), employeeData)
	if err != nil {
		h.log.Error("create employee", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	h.log.Info("created new employee", "id", id)
	utils.Send201(w, models.ResponseID{ID: id})
}

// UpdateEmployee godoc
// @Summary      Изменить данные сотрудника
// @Description  Изменить данные о сотруднике
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id path string true "employee id"
// @Param        request body models.CreateEmployee true "employee data"
// @Success      200  {object} utils.MessageResponse
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /employees/{id} [patch]
func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeData *models.CreateEmployee

	if err := utils.ReadRequestData(r, &employeeData); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	employeeData.ID = int32(id)
	h.log.Debug("update employee", "data", employeeData)
	err = h.uc.EditEmployee(r.Context(), employeeData)
	if err != nil {
		h.log.Error("edit employee", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	h.log.Info("updated employee")
	utils.Send200(w, utils.MessageResponse{Msg: "employee updated"})

}

// DeleteEmployee godoc
// @Summary      Удалить сотрудника
// @Description  Удалить сотрудника
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id path string true "employee id"
// @Success      200  {object} utils.MessageResponse
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /employees/{id} [delete]
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	if err = h.uc.DeleteEmployee(r.Context(), int32(id)); err != nil {
		h.log.Error("delete employee", "error", err.Error())
		utils.Send404(w, messages.NotFound)
		return
	}

	h.log.Info("deleted employee")
	utils.Send200(w, utils.MessageResponse{Msg: "employee deleted"})
}

// GetCompanyEmployees godoc
// @Summary      Получить сотрудников компании
// @Description  Вывести список сотрудников компании
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id path string true "company id"
// @Success      200  {object} []models.Employee
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /companies/{id}/employees [get]
func (h *Handler) GetCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("parse employee id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	listEmployees, err := h.uc.GetListCompanyEmployees(r.Context(), int32(id))
	if err != nil {
		h.log.Error("get list of employees", "error", err.Error())
		utils.Send404(w, messages.NotFound)
		return
	}

	h.log.Info("got list of company employees")
	utils.Send200(w, listEmployees)
}

// GetDepartmentCompanyEmployees godoc
// @Summary      Получить сотрудников отдела компании
// @Description  Вывести список сотрудников отдела компании
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id path string true "department id"
// @Success      200  {object} []models.Employee
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /departments/{id}/employees [get]
func (h *Handler) GetDepartmentCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	departmentIDStr := vars["id"]
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		h.log.Error("parse department id", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	listEmployees, err := h.uc.GetListDepartmentCompanyEmployees(r.Context(), int32(departmentID))
	if err != nil {
		h.log.Error("get list of department employees", "error", err.Error())
		utils.Send404(w, messages.NotFound)
		return
	}

	h.log.Info("got list of department employees")
	utils.Send200(w, listEmployees)
}

// CreateCompany godoc
// @Summary      Создать компанию
// @Description  Создать новую компанию
// @Tags         companies
// @Accept       json
// @Produce      json
// @Param        name body models.Company true "company name"
// @Success      200  {object} models.ResponseID
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /companies [post]
func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company *models.Company
	if err := utils.ReadRequestData(r, &company); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	companyID, err := h.uc.CreateCompany(r.Context(), company.Name)
	if err != nil {
		h.log.Error("create company", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	h.log.Info("created company", "id", companyID)
	utils.Send201(w, models.ResponseID{ID: companyID})
}

// CreateDepartment godoc
// @Summary      Создать отдел
// @Description  Создать новый отдел компании
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        request body models.CreateDepartment true "department data"
// @Success      200  {object} models.ResponseID
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      500  {object} string
// @Router       /departments [post]
func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var department *models.CreateDepartment
	if err := utils.ReadRequestData(r, &department); err != nil {
		h.log.Error("read request data", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	departmentID, err := h.uc.CreateDepartment(r.Context(), department)
	if err != nil {
		h.log.Error("create department", "error", err.Error())
		utils.Send400(w, messages.BadRequest)
		return
	}

	h.log.Info("created department", "id", departmentID)
	utils.Send201(w, models.ResponseID{ID: departmentID})
}
