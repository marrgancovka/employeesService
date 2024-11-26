package delivery

import (
	"employees/internal/models"
	"employees/internal/pkg/employee/usecase"
	"employees/internal/pkg/utils"
	"employees/internal/pkg/utils/messages"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

type Params struct {
	fx.In

	Uc *usecase.Usecase
}

type Handler struct {
	uc *usecase.Usecase
}

func New(p Params) *Handler {
	return &Handler{
		uc: p.Uc,
	}
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee *models.Employee

	if err := utils.ReadRequestData(r, &employee); err != nil {
		utils.Send400(w, messages.BadRequest)
	}

	id, err := h.uc.CreateEmployee(r.Context(), employee)
	if err != nil {
		utils.Send400(w, messages.BadRequest)
	}
	utils.Send201(w, id)
}

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee *models.Employee

	if err := utils.ReadRequestData(r, &employee); err != nil {
		utils.Send400(w, messages.BadRequest)
	}

	err := h.uc.EditEmployee(r.Context(), employee)
	if err != nil {
		utils.Send400(w, messages.BadRequest)
	}
	utils.Send200(w, "employee updated")

}
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Send400(w, messages.BadRequest)
	}
	if err = h.uc.DeleteEmployee(r.Context(), int32(id)); err != nil {
		utils.Send404(w, messages.NotFound)
	}
	utils.Send200(w, "employee deleted")
}
func (h *Handler) GetCompanyEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Send400(w, messages.BadRequest)
	}
	listEmployees, err := h.uc.GetListCompanyEmployees(r.Context(), int32(id))
	if err != nil {
		utils.Send404(w, messages.NotFound)
	}
	utils.Send200(w, listEmployees)
}
func (h *Handler) GetDepartmentCompanyEmployees(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {

}
