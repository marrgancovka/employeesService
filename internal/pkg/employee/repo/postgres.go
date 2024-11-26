package repo

import (
	"context"
	"employees/gen"
	"employees/internal/models"
	"employees/internal/pkg/txer"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	DB *pgxpool.Pool
}

type PostgresRepo struct {
	db      *pgxpool.Pool
	queries *gen.Queries
}

func New(p Params) *PostgresRepo {
	return &PostgresRepo{
		db:      p.DB,
		queries: gen.New(p.DB),
	}
}

func (r *PostgresRepo) CreateEmployee(ctx context.Context, employee *models.Employee) (int32, error) {
	return txer.InTxWithValue(ctx, r.db, func(txConn *pgx.Tx) (int32, error) {
		tx := r.queries.WithTx(*txConn)
		departmentID, err := tx.CreateDepartment(ctx, gen.CreateDepartmentParams{
			Name:      employee.Department.Name,
			Phone:     employee.Department.Phone,
			CompanyID: employee.CompanyID,
		})
		if err != nil {
			return 0, err
		}
		createEmployeeID, err := tx.CreateEmployee(ctx, gen.CreateEmployeeParams{
			Name:           employee.Name,
			Surname:        employee.Surname,
			Phone:          employee.Phone,
			CompanyID:      employee.CompanyID,
			DepartmentID:   pgtype.Int4{Int32: departmentID, Valid: true},
			PassportType:   employee.Passport.Type,
			PassportNumber: employee.Passport.Number,
		})
		if err != nil {
			return 0, err
		}
		return createEmployeeID, err
	})
}

func (r *PostgresRepo) DeleteEmployee(ctx context.Context, id int32) error {
	err := r.queries.DeleteEmployee(ctx, id)
	return err
}

func (r *PostgresRepo) GetListCompanyEmployees(ctx context.Context, id int32) ([]*models.Employee, error) {
	employees, err := r.queries.GetListCompanyEmployee(ctx, id)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
func (r *PostgresRepo) GetListDepartmentEmployees(ctx context.Context, idCompany, idDepartment int32) ([]*models.Employee, error) {

}
func (r *PostgresRepo) EditEmployee(ctx context.Context, employee *models.Employee) error {
	return txer.InTx(ctx, r.db, func(txConn *pgx.Tx) error {

		oldEmployee, err := r.GetEmployeeByID(ctx, employee.ID)
		if err != nil {
			return err
		}

		err = r.queries.UpdateEmployee(ctx, gen.UpdateEmployeeParams{
			ID:             employee.ID,
			Name:           lo.Ternary(employee.Name == "", oldEmployee.Name, employee.Name),
			Surname:        lo.Ternary(employee.Surname == "", oldEmployee.Surname, employee.Surname),
			Phone:          lo.Ternary(employee.Phone == "", oldEmployee.Phone, employee.Phone),
			CompanyID:      lo.Ternary(employee.CompanyID == 0, oldEmployee.CompanyID, employee.CompanyID),
			PassportType:   lo.Ternary(employee.Passport.Type == "", oldEmployee.Passport.Type, employee.Passport.Number),
			PassportNumber: lo.Ternary(employee.Passport.Number == "", oldEmployee.Passport.Number, employee.Passport.Number),
		})
		return err

	})
}
func (r *PostgresRepo) CreateCompany(ctx context.Context, name string) (int32, error) {

}
func (r *PostgresRepo) CreateDepartment(ctx context.Context, department *models.Department) (int32, error) {

}
func (r *PostgresRepo) GetEmployeeByID(ctx context.Context, id int32) (*models.Employee, error) {

}
