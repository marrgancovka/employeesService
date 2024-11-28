package repo

import (
	"context"
	"employees/gen"
	"employees/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"log/slog"
)

type Params struct {
	fx.In
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

type PostgresRepo struct {
	db      *pgxpool.Pool
	queries *gen.Queries
	log     *slog.Logger
}

func New(p Params) *PostgresRepo {
	return &PostgresRepo{
		db:      p.DB,
		queries: gen.New(p.DB),
		log:     p.Logger,
	}
}

func (r *PostgresRepo) CreateEmployee(ctx context.Context, employee *models.Employee) (int32, error) {
	createEmployeeID, err := r.queries.CreateEmployee(ctx, gen.CreateEmployeeParams{
		Name:           employee.Name,
		Surname:        employee.Surname,
		Phone:          employee.Phone,
		CompanyID:      employee.CompanyID,
		DepartmentID:   employee.Department.ID,
		PassportType:   employee.Passport.Type,
		PassportNumber: employee.Passport.Number,
	})
	if err != nil {
		r.log.Error("create employee", "error", err)
		return 0, err
	}
	return createEmployeeID, nil
}

func (r *PostgresRepo) DeleteEmployee(ctx context.Context, id int32) error {
	err := r.queries.DeleteEmployee(ctx, id)
	if err != nil {
		r.log.Error("delete employee", "error", err)
		return err
	}
	return nil
}

func (r *PostgresRepo) GetListCompanyEmployees(ctx context.Context, id int32) ([]*models.Employee, error) {
	employees, err := r.queries.GetListCompanyEmployee(ctx, id)
	if err != nil {
		r.log.Error("get employees", "error", err)
		return nil, err
	}

	listEmployees := make([]*models.Employee, len(employees))
	for i, employee := range employees {
		listEmployees[i] = &models.Employee{
			ID:        employee.ID,
			Name:      employee.Name,
			Surname:   employee.Surname,
			Phone:     employee.Phone,
			CompanyID: employee.CompanyID,
			Passport: models.Passport{
				Type:   employee.PassportType,
				Number: employee.PassportNumber,
			},
			Department: models.Department{
				Name:  employee.Name_2,
				Phone: employee.Phone_2,
			},
		}
	}

	return listEmployees, nil
}
func (r *PostgresRepo) GetListDepartmentEmployees(ctx context.Context, idDepartment int32) ([]*models.Employee, error) {
	employees, err := r.queries.GetListCompanyDepartmentEmployee(ctx, idDepartment)
	if err != nil {
		r.log.Error("get employees", "error", err)
		return nil, err
	}

	listEmployees := make([]*models.Employee, len(employees))
	for i, employee := range employees {
		listEmployees[i] = &models.Employee{
			ID:        employee.ID,
			Name:      employee.Name,
			Surname:   employee.Surname,
			Phone:     employee.Phone,
			CompanyID: employee.CompanyID,
			Passport: models.Passport{
				Type:   employee.PassportType,
				Number: employee.PassportNumber,
			},
			Department: models.Department{
				Name:  employee.Name_2,
				Phone: employee.Phone_2,
			},
		}
	}

	return listEmployees, nil
}
func (r *PostgresRepo) EditEmployee(ctx context.Context, employee *models.Employee) error {
	r.log.Debug("edit employee", "employee", employee)
	oldEmployee, err := r.GetEmployeeByID(ctx, employee.ID)
	if err != nil {
		r.log.Error("get employee", "error", err)
		return err
	}
	r.log.Debug("old employee", "old", oldEmployee)
	err = r.queries.UpdateEmployee(ctx, gen.UpdateEmployeeParams{
		ID:             employee.ID,
		Name:           lo.Ternary(employee.Name == "", oldEmployee.Name, employee.Name),
		Surname:        lo.Ternary(employee.Surname == "", oldEmployee.Surname, employee.Surname),
		Phone:          lo.Ternary(employee.Phone == "", oldEmployee.Phone, employee.Phone),
		CompanyID:      lo.Ternary(employee.CompanyID == 0, oldEmployee.CompanyID, employee.CompanyID),
		DepartmentID:   lo.Ternary(employee.Department.ID == 0, oldEmployee.Department.ID, employee.Department.ID),
		PassportType:   lo.Ternary(employee.Passport.Type == "", oldEmployee.Passport.Type, employee.Passport.Number),
		PassportNumber: lo.Ternary(employee.Passport.Number == "", oldEmployee.Passport.Number, employee.Passport.Number),
	})
	if err != nil {
		r.log.Error("update employee", "error", err)
		return err
	}

	return nil

}
func (r *PostgresRepo) CreateCompany(ctx context.Context, name string) (int32, error) {
	companyID, err := r.queries.CreateCompany(ctx, name)
	if err != nil {
		r.log.Error("create company", "error", err)
		return 0, err
	}

	return companyID, nil
}
func (r *PostgresRepo) CreateDepartment(ctx context.Context, department *models.Department) (int32, error) {
	departmentID, err := r.queries.CreateDepartment(ctx, gen.CreateDepartmentParams{
		Name:      department.Name,
		Phone:     department.Phone,
		CompanyID: department.CompanyID,
	})
	if err != nil {
		r.log.Error("create department", "error", err)
		return 0, err
	}

	return departmentID, nil
}
func (r *PostgresRepo) GetEmployeeByID(ctx context.Context, id int32) (*models.Employee, error) {
	employee, err := r.queries.GetEmployeeByID(ctx, id)
	if err != nil {
		r.log.Error("get employee", "error", err)
		return nil, err
	}

	modelEmployee := &models.Employee{
		ID:        employee.ID,
		Name:      employee.Name,
		Surname:   employee.Surname,
		Phone:     employee.Phone,
		CompanyID: employee.CompanyID,
		Passport: models.Passport{
			Type:   employee.PassportType,
			Number: employee.PassportNumber,
		},
		Department: models.Department{
			ID: employee.DepartmentID,
		},
	}

	return modelEmployee, nil
}
