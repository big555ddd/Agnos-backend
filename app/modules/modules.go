package modules

import (
	"app/app/modules/patient"
	"app/app/modules/staff"
	"app/config"
)

type Module struct {
	Patient *patient.Module
	Staff   *staff.Module
}

func New() *Module {

	db := config.GetDB()
	patient := patient.NewModule(db)
	staff := staff.NewModule(db)

	return &Module{
		Patient: patient,
		Staff:   staff,
	}
}
