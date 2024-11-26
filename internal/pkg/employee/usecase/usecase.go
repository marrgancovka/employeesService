package usecase

import (
	"context"
	"employees/internal/models"
	"employees/internal/pkg/employee/repo"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Repo *repo.PostgresRepo
}

type Usecase struct {
	repo *repo.PostgresRepo
}

func New(p Params) *Usecase {
	return &Usecase{
		repo: p.Repo,
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
func (uc *Usecase) GetListDepartmentCompanyEmployees(ctx context.Context, companyID, departmentID int32) ([]*models.Employee, error) {

}
func (uc *Usecase) EditEmployee(ctx context.Context, employee *models.Employee) error {
	if err := uc.repo.EditEmployee(ctx, employee); err != nil {
		return err
	}
	return nil
}
func (uc *Usecase) CreateCompany(ctx context.Context, name string) (int32, error) {

}
func (uc *Usecase) CreateDepartment(ctx context.Context, department *models.Department) (int32, error) {

}
