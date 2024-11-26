-- name: CreateCompany :one
INSERT INTO companies (name)
VALUES ($1) RETURNING id;

-- name: CreateDepartment :one
INSERT INTO departments (name, phone, company_id)
VALUES ($1, $2, $3) ON CONFLICT (name, company_id) DO NOTHING
RETURNING id;

-- name: CreateEmployee :one
INSERT INTO employees (name, surname, phone, company_id, department_id, passport_type, passport_number)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;

-- name: GetListCompanyEmployee :many
SELECT e.id,
       e.name,
       e.surname,
       e.phone,
       e.company_id,
       e.passport_type,
       e.passport_number,
       d.name,
       d.phone
FROM employees e
         JOIN departments d ON e.department_id = d.id
WHERE e.company_id = $1
ORDER BY e.id asc;

-- name: GetListCompanyDepartmentEmployee :many
SELECT e.id,
       e.name,
       e.surname,
       e.phone,
       e.company_id,
       e.passport_type,
       e.passport_number,
       d.name,
       d.phone
FROM employees e
         JOIN departments d ON e.department_id = d.id
WHERE e.company_id = $1
  AND e.department_id = $2
ORDER BY e.id asc;

-- name: DeleteEmployee :exec
DELETE
FROM employees
WHERE id = $1;

-- name: UpdateEmployee :exec
UPDATE employees
SET name=$2,
    surname=$3,
    phone=$4,
    company_id=$5,
    passport_type=$6,
    passport_number=$7,
    updated_at=now()
WHERE id = $1;

-- name: GetEmployeeByID :one
SELECT id,
       name,
       surname,
       phone,
       company_id,
       passport_type,
       passport_number,
       department_id
FROM employees
WHERE id = $1;


-- name: GetDepartmentByID :one
SELECT name, phone
FROM departments
WHERE id = $1;

-- name: GetDepartmentID :one
SELECT department_id
FROM employees
WHERE id = $1;