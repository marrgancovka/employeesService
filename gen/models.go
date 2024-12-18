// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gen

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Company struct {
	ID   int32
	Name string
}

type Department struct {
	ID        int32
	Name      string
	Phone     string
	CompanyID int32
	CreatedAt pgtype.Timestamptz
}

type Employee struct {
	ID             int32
	Name           string
	Surname        string
	Phone          string
	CompanyID      int32
	DepartmentID   int32
	PassportType   string
	PassportNumber string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}
