package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model     `json:"-"`
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	JobDescription string `json:"job_desc"`
	EntryDate      int    `json:"entry_date"`
}
