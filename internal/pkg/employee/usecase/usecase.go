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

func (uc *Usecase) CreateEmployee(ctx context.Context, employee *models.CreateEmployee) (int32, error) {
	employeeData := &models.Employee{
		Name:      employee.Name,
		Surname:   employee.Surname,
		Phone:     employee.Phone,
		CompanyID: employee.CompanyID,
		Passport: models.Passport{
			Type:   employee.Passport.Type,
			Number: employee.Passport.Number,
		},
		Department: models.Department{
			ID: employee.DepartmentID,
		},
	}

	id, err := uc.repo.CreateEmployee(ctx, employeeData)
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

func (uc *Usecase) EditEmployee(ctx context.Context, employee *models.CreateEmployee) error {
	employeeData := &models.Employee{
		ID:        employee.ID,
		Name:      employee.Name,
		Surname:   employee.Surname,
		Phone:     employee.Phone,
		CompanyID: employee.CompanyID,
		Passport: models.Passport{
			Type:   employee.Passport.Type,
			Number: employee.Passport.Number,
		},
		Department: models.Department{
			ID: employee.DepartmentID,
		},
	}

	if err := uc.repo.EditEmployee(ctx, employeeData); err != nil {
		return err
	}
	return nil
}
func (uc *Usecase) CreateCompany(ctx context.Context, name string) (int32, error) {
	id, err := uc.repo.CreateCompany(ctx, name)
	return id, err
}
func (uc *Usecase) CreateDepartment(ctx context.Context, department *models.CreateDepartment) (int32, error) {
	departmentDB := &models.Department{
		Name:      department.Name,
		Phone:     department.Phone,
		CompanyID: department.CompanyID,
	}
	id, err := uc.repo.CreateDepartment(ctx, departmentDB)
	return id, err
}
