package usecase

import (
	"context"
	"employees/internal/models"
	"employees/internal/pkg/employee"
	"go.uber.org/fx"
	"log/slog"
)

type Params struct {
	fx.In

	Repo   employee.Repository
	Logger *slog.Logger
}

type Usecase struct {
	repo employee.Repository
	log  *slog.Logger
}

func New(p Params) *Usecase {
	return &Usecase{
		repo: p.Repo,
		log:  p.Logger,
	}
}

func (uc *Usecase) CreateEmployee(ctx context.Context, employee *models.Employee) (int32, error) {
	id, err := uc.repo.CreateEmployee(ctx, employee)
	return id, err
}
func (uc *Usecase) DeleteEmployee(ctx context.Context, id int32) error {
	err := uc.repo.DeleteEmployee(ctx, id)
	return err
}
func (uc *Usecase) GetListCompanyEmployees(ctx context.Context, companyID int32) ([]*models.Employee, error) {
	listEmployees, err := uc.repo.GetListCompanyEmployees(ctx, companyID)
	return listEmployees, err
}

func (uc *Usecase) GetListDepartmentCompanyEmployees(ctx context.Context, departmentID int32) ([]*models.Employee, error) {
	listEmployees, err := uc.repo.GetListDepartmentEmployees(ctx, departmentID)
	return listEmployees, err
}

func (uc *Usecase) EditEmployee(ctx context.Context, employee *models.Employee) error {
	if err := uc.repo.EditEmployee(ctx, employee); err != nil {
		return err
	}
	return nil
}
func (uc *Usecase) CreateCompany(ctx context.Context, name string) (int32, error) {
	id, err := uc.repo.CreateCompany(ctx, name)
	return id, err
}
func (uc *Usecase) CreateDepartment(ctx context.Context, department *models.Department) (int32, error) {
	id, err := uc.repo.CreateDepartment(ctx, department)
	return id, err
}
