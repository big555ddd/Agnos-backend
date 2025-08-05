package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"table:patients"`

	ID           string    `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	FirstNameTH  string    `bun:"first_name_th" json:"first_name_th"`
	MiddleNameTH string    `bun:"middle_name_th" json:"middle_name_th"`
	LastNameTH   string    `bun:"last_name_th" json:"last_name_th"`
	FirstNameEN  string    `bun:"first_name_en" json:"first_name_en"`
	MiddleNameEN string    `bun:"middle_name_en" json:"middle_name_en"`
	LastNameEN   string    `bun:"last_name_en" json:"last_name_en"`
	DateOfBirth  time.Time `bun:"date_of_birth,type:date" json:"date_of_birth"`
	PatientHN    string    `bun:"patient_hn" json:"patient_hn"`
	NationalID   string    `bun:"national_id,unique,nullzero" json:"national_id"`
	PassportID   string    `bun:"passport_id,unique,nullzero" json:"passport_id"`
	PhoneNumber  string    `bun:"phone_number" json:"phone_number"`
	Email        string    `bun:"email" json:"email"`
	Gender       string    `bun:"gender,type:char(1)" json:"gender"`
	Hospital     string    `bun:"hospital,notnull" json:"hospital"`

	_ struct{} `bun:"index:(first_name_th, first_name_en),index:(middle_name_th, middle_name_en),index:(last_name_th, last_name_en)"`
	_ struct{} `bun:"index:date_of_birth"`
	_ struct{} `bun:"index:email"`
	_ struct{} `bun:"index:phone_number"`
	_ struct{} `bun:"index:hospital"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
