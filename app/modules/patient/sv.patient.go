package patient

import (
	"app/app/model"
	patientdto "app/app/modules/patient/dto"
	"app/internal/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/net/context"
)

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetPatient(ctx context.Context, id string) (*http.Response, error) {
	url := fmt.Sprintf("https://hospital-a.api.co.th/patient/search/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (s *Service) List(ctx context.Context, req *patientdto.ListPatientRequest, hospital string) ([]*model.Patient, int, error) {
	resp := []*model.Patient{}
	var (
		offset = (req.Page - 1) * req.Size
		limit  = req.Size
	)
	logger.Info(req)

	query := s.db.NewSelect().
		Model(&resp).
		Where("hospital = ?", hospital)
	if req.NationalID != "" {
		nationalID := fmt.Sprint(strings.ToLower(req.NationalID) + "%")
		query.Where("LOWER(national_id) LIKE ?", nationalID)
	}

	if req.PassportID != "" {
		passportID := fmt.Sprint(strings.ToLower(req.PassportID) + "%")
		query.Where("LOWER(passport_id) LIKE ?", passportID)
	}

	if req.FirstName != "" {
		firstName := fmt.Sprint(strings.ToLower(req.FirstName) + "%")
		query.Where("LOWER(first_name_th) LIKE ? OR LOWER(first_name_en) LIKE ?", firstName, firstName)
	}

	if req.MiddleName != "" {
		middleName := fmt.Sprint(strings.ToLower(req.MiddleName) + "%")
		query.Where("LOWER(middle_name_th) LIKE ? OR LOWER(middle_name_en) LIKE ?", middleName, middleName)
	}

	if req.LastName != "" {
		lastName := fmt.Sprint(strings.ToLower(req.LastName) + "%")
		query.Where("LOWER(last_name_th) LIKE ? OR LOWER(last_name_en) LIKE ?", lastName, lastName)
	}

	if req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err == nil {
			query.Where("date_of_birth = ?", dob)
		}
	}
	if req.Email != "" {
		email := fmt.Sprint(strings.ToLower(req.Email) + "%")
		query.Where("LOWER(email) LIKE ?", email)
	}

	if req.PhoneNumber != "" {
		phoneNumber := fmt.Sprint(strings.ToLower(req.PhoneNumber) + "%")
		query.Where("LOWER(phone_number) LIKE ?", phoneNumber)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return resp, 0, err
	}
	if total == 0 {
		return resp, 0, nil
	}
	order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)

	err = query.
		Offset(offset).
		Limit(limit).
		Order(order).
		Scan(ctx, &resp)
	if err != nil {
		return resp, 0, err
	}

	return resp, total, nil
}
