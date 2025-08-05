package patient

import (
	"app/app/model"
	patientdto "app/app/modules/patient/dto"
	"context"
	"net/http"
)

type ServiceInterface interface {
	GetPatient(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, req *patientdto.ListPatientRequest, hospital string) ([]*model.Patient, int, error)
}

var _ ServiceInterface = (*Service)(nil)
