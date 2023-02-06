package dtos

type Employee struct {
	Name           string `json:"name"`
	JobDescription string `json:"job_desc"`
	EntryDate      int    `json:"entry_date"`
}

type EditEmployee struct {
	UserID int    `json:"-"`
	Name   string `json:"name"`
}

type EmployeeOutput struct {
	Name           string `json:"name"`
	JobDescription string `json:"job_desc"`
	EntryDate      int    `json:"entry_date"`
}
