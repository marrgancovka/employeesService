package http

import (
	"bytes"
	"employees/internal/models"
	mockEmployee "employees/internal/pkg/employee/mocks"
	"employees/internal/pkg/logger"
	"employees/internal/pkg/utils/messages"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateCompany(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, company models.Company)
	testTable := []struct {
		name         string
		inputBody    string
		inputCompany models.Company
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:      "success",
			inputBody: `{"name":"test_company"}`,
			inputCompany: models.Company{
				Name: "test_company",
			},
			mockBehavior: func(m *mockEmployee.MockUsecase, company models.Company) {
				m.EXPECT().CreateCompany(gomock.Any(), company.Name).Return(int32(1), nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":1}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.inputCompany)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/companies", handler.CreateCompany)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/companies", bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())

		})
	}
}

func TestHandler_CreateDepartment(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, department *models.CreateDepartment)
	testTable := []struct {
		name            string
		inputBody       string
		inputDepartment *models.CreateDepartment
		mockBehavior    mockBehavior
		expectedCode    int
		expectedBody    string
	}{
		{
			name:      "success",
			inputBody: `{"name":"test_department", "phone": "896753617", "company_id":1}`,
			inputDepartment: &models.CreateDepartment{
				Name:      "test_department",
				Phone:     "896753617",
				CompanyID: 1,
			},
			mockBehavior: func(m *mockEmployee.MockUsecase, department *models.CreateDepartment) {
				m.EXPECT().CreateDepartment(gomock.Any(), department).Return(int32(1), nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":1}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.inputDepartment)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/departments", handler.CreateDepartment)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/departments", bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}

}

func TestHandler_CreateEmployee(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, employee *models.CreateEmployee)

	testTable := []struct {
		name          string
		inputBody     string
		inputEmployee *models.CreateEmployee
		mockBehavior  mockBehavior
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "ok",
			inputBody: `{"company_id": 1,"department_id": 1,"name": "katya","passport": {"number": "7878 898989",  "type": "РФ"  }, "phone": "93097383","surname": "ivanova"}`,
			inputEmployee: &models.CreateEmployee{
				Name:      "katya",
				Surname:   "ivanova",
				Phone:     "93097383",
				CompanyID: 1,
				Passport: models.Passport{
					Number: "7878 898989",
					Type:   "РФ",
				},
				DepartmentID: 1,
			},
			mockBehavior: func(m *mockEmployee.MockUsecase, employee *models.CreateEmployee) {
				m.EXPECT().CreateEmployee(gomock.Any(), employee).Return(int32(1), nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":1}`,
		},
		{
			name:      "fail",
			inputBody: `{"company_id": 0,"department_id": 1,"name": "katya","passport": {"number": "7878 898989",  "type": "РФ"  }, "phone": "93097383","surname": "ivanova"}`,
			inputEmployee: &models.CreateEmployee{
				Name:      "katya",
				Surname:   "ivanova",
				Phone:     "93097383",
				CompanyID: 0,
				Passport: models.Passport{
					Number: "7878 898989",
					Type:   "РФ",
				},
				DepartmentID: 1,
			},
			mockBehavior: func(m *mockEmployee.MockUsecase, employee *models.CreateEmployee) {
				m.EXPECT().CreateEmployee(gomock.Any(), employee).Return(int32(0), fmt.Errorf(messages.BadRequest))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"msg":"%v"}`, messages.BadRequest),
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.inputEmployee)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/employees", handler.CreateEmployee)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestHandler_UpdateEmployee(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, employee *models.CreateEmployee)

	testTable := []struct {
		name          string
		inputBody     string
		inputEmployee *models.CreateEmployee
		employeeID    string
		mockBehavior  mockBehavior
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "ok: all fields",
			inputBody: `{"company_id": 2,"department_id": 2,"name": "ivan","passport": {"number": "1212 343434",  "type": "РБ"  }, "phone": "3413434","surname": "petrov"}`,
			inputEmployee: &models.CreateEmployee{
				ID:        1,
				Name:      "ivan",
				Surname:   "petrov",
				Phone:     "3413434",
				CompanyID: 2,
				Passport: models.Passport{
					Number: "1212 343434",
					Type:   "РБ",
				},
				DepartmentID: 2,
			},
			employeeID: "1",
			mockBehavior: func(m *mockEmployee.MockUsecase, employee *models.CreateEmployee) {
				m.EXPECT().EditEmployee(gomock.Any(), employee).Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"msg":"employee updated"}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.inputEmployee)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/employees/{id}", handler.UpdateEmployee)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/employees/"+tt.employeeID, bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestHandler_DeleteEmployee(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, id int32)
	testTable := []struct {
		name         string
		employeeID   string
		ID           int32
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:       "ok",
			employeeID: "1",
			ID:         int32(1),
			mockBehavior: func(m *mockEmployee.MockUsecase, id int32) {
				m.EXPECT().DeleteEmployee(gomock.Any(), id).Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"msg":"employee deleted"}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.ID)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/employees/{id}", handler.DeleteEmployee)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/employees/"+tt.employeeID, nil)

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})

	}
}

func TestHandler_GetCompanyEmployees(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, companyID int32)
	testTable := []struct {
		name         string
		companyID    int32
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			companyID: 4,
			mockBehavior: func(m *mockEmployee.MockUsecase, companyID int32) {
				m.EXPECT().GetListCompanyEmployees(gomock.Any(), companyID).Return([]*models.Employee{
					{
						ID:        5,
						Name:      "ruslan",
						Surname:   "ruslanov",
						Phone:     "89776677",
						CompanyID: 4,
						Passport: models.Passport{
							Type:   "0989",
							Number: "0989",
						},
						Department: models.Department{
							Name:  "marketing",
							Phone: "89",
						},
					},
					{
						ID:        6,
						Name:      "masha",
						Surname:   "naumova",
						Phone:     "12095",
						CompanyID: 4,
						Passport: models.Passport{
							Type:   "rf",
							Number: "487484",
						},
						Department: models.Department{
							Name:  "marketing",
							Phone: "89",
						},
					},
					{
						ID:        7,
						Name:      "anna",
						Surname:   "petrova",
						Phone:     "18385",
						CompanyID: 4,
						Passport: models.Passport{
							Type:   "rf",
							Number: "407484",
						},
						Department: models.Department{
							Name:  "dev",
							Phone: "1234",
						},
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `[
  {
    "id": 5,
    "name": "ruslan",
    "surname": "ruslanov",
    "phone": "89776677",
    "company_id": 4,
    "passport": {
      "type": "0989",
      "number": "0989"
    },
    "department": {
      "name": "marketing",
      "phone": "89"
    }
  },
  {
    "id": 6,
    "name": "masha",
    "surname": "naumova",
    "phone": "12095",
    "company_id": 4,
    "passport": {
      "type": "rf",
      "number": "487484"
    },
    "department": {
      "name": "marketing",
      "phone": "89"
    }
  },
  {
    "id": 7,
    "name": "anna",
    "surname": "petrova",
    "phone": "18385",
    "company_id": 4,
    "passport": {
      "type": "rf",
      "number": "407484"
    },
    "department": {
      "name": "dev",
      "phone": "1234"
    }
  }
]`,
		},
		{
			name:      "empty",
			companyID: 10,
			mockBehavior: func(m *mockEmployee.MockUsecase, companyID int32) {
				m.EXPECT().GetListCompanyEmployees(gomock.Any(), companyID).Return([]*models.Employee{}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: "[]",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.companyID)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/companies/{id}/employees", handler.GetCompanyEmployees)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/companies/%v/employees", tt.companyID), nil)

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestHandler_GetDepartmentCompanyEmployees(t *testing.T) {
	type mockBehavior func(m *mockEmployee.MockUsecase, departmentID int32)
	testTable := []struct {
		name         string
		departmentID int32
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:         "ok",
			departmentID: 2,
			mockBehavior: func(m *mockEmployee.MockUsecase, departmentID int32) {
				m.EXPECT().GetListDepartmentCompanyEmployees(gomock.Any(), departmentID).Return([]*models.Employee{
					{
						ID:        5,
						Name:      "ruslan",
						Surname:   "ruslanov",
						Phone:     "89776677",
						CompanyID: 4,
						Passport: models.Passport{
							Type:   "0989",
							Number: "0989",
						},
						Department: models.Department{
							Name:  "marketing",
							Phone: "89",
						},
					},
					{
						ID:        6,
						Name:      "masha",
						Surname:   "naumova",
						Phone:     "12095",
						CompanyID: 4,
						Passport: models.Passport{
							Type:   "rf",
							Number: "487484",
						},
						Department: models.Department{
							Name:  "marketing",
							Phone: "89",
						},
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `[
  {
    "id": 5,
    "name": "ruslan",
    "surname": "ruslanov",
    "phone": "89776677",
    "company_id": 4,
    "passport": {
      "type": "0989",
      "number": "0989"
    },
    "department": {
      "name": "marketing",
      "phone": "89"
    }
  },
  {
    "id": 6,
    "name": "masha",
    "surname": "naumova",
    "phone": "12095",
    "company_id": 4,
    "passport": {
      "type": "rf",
      "number": "487484"
    },
    "department": {
      "name": "marketing",
      "phone": "89"
    }
  }
]`,
		},
		{
			name:         "empty",
			departmentID: 122,
			mockBehavior: func(m *mockEmployee.MockUsecase, departmentID int32) {
				m.EXPECT().GetListDepartmentCompanyEmployees(gomock.Any(), departmentID).Return([]*models.Employee{}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: "[]",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecaseEmployee := mockEmployee.NewMockUsecase(ctrl)
			tt.mockBehavior(mockUsecaseEmployee, tt.departmentID)

			handler := &Handler{
				uc:  mockUsecaseEmployee,
				log: logger.SetupLogger(),
			}

			router := mux.NewRouter()
			router.HandleFunc("/departments/{id}/employees", handler.GetDepartmentCompanyEmployees)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/departments/%v/employees", tt.departmentID), nil)

			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}
