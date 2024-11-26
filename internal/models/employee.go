package models

type Passport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Department struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Employee struct {
	ID         int32      `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyID  int32      `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}
