package employee

import (
	"context"
	"employees/internal/models"
)

type Usecase interface {
	CreateEmployee(ctx context.Context, employee *models.Employee) (int32, error)
	DeleteEmployee(ctx context.Context, id int32) error
	GetListCompanyEmployees(ctx context.Context, companyID int32) ([]*models.Employee, error)
	GetListDepartmentCompanyEmployees(ctx context.Context, departmentID int32) ([]*models.Employee, error)
	EditEmployee(ctx context.Context, employee *models.Employee) error
	CreateCompany(ctx context.Context, name string) (int32, error)
	CreateDepartment(ctx context.Context, department *models.Department) (int32, error)
}

type Repository interface {
	CreateEmployee(ctx context.Context, employee *models.Employee) (int32, error)
	DeleteEmployee(ctx context.Context, id int32) error
	GetListCompanyEmployees(ctx context.Context, companyID int32) ([]*models.Employee, error)
	GetListDepartmentEmployees(ctx context.Context, departmentID int32) ([]*models.Employee, error)
	EditEmployee(ctx context.Context, employee *models.Employee) error
	CreateCompany(ctx context.Context, name string) (int32, error)
	CreateDepartment(ctx context.Context, department *models.Department) (int32, error)
	GetEmployeeByID(ctx context.Context, id int32) (*models.Employee, error)
}
