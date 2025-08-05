package staff

import (
	"app/app/message"
	"app/app/model"
	staffdto "app/app/modules/staff/dto"
	"app/app/util/hashing"
	"app/app/util/jwt"
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create(ctx context.Context, req *staffdto.CreateStaffRequest) error {
	// Check if username already exists
	exists, err := s.ExistUsername(ctx, req.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(message.StaffAlreadyExists)
	}
	//hashpassword
	hash, err := hashing.HashPassword(req.Password)
	if err != nil {
		return err
	}
	// Create new staff record
	data := &model.Staff{
		Username: req.Username,
		Password: string(hash),
		Hospital: req.Hospital,
	}
	_, err = s.db.NewInsert().
		Model(data).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ExistUsername(ctx context.Context, username string) (bool, error) {
	ex, err := s.db.NewSelect().
		Model((*model.Staff)(nil)).
		Where("username = ?", username).
		Exists(ctx)
	return ex, err
}

func (s *Service) GetStaffByUsername(ctx context.Context, username string) (*model.Staff, error) {
	staff := new(model.Staff)
	err := s.db.NewSelect().
		Model(staff).
		Where("username = ?", username).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(message.StaffNotFound)
		}
		return nil, err
	}
	return staff, nil
}

func (s *Service) Login(ctx context.Context, req *staffdto.LoginStaffRequest) (string, error) {
	// Find staff by username
	staff, err := s.GetStaffByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}
	// Verify password
	if !hashing.CheckPasswordHash(staff.Password, req.Password) {
		return "", errors.New(message.InvalidCredentials)
	}

	// Verify hospital
	if staff.Hospital != req.Hospital {
		return "", errors.New(message.InvalidCredentials)
	}
	claim := jwt.ClaimData{
		ID:       staff.ID,
		Username: staff.Username,
		Hospital: staff.Hospital,
	}
	//Create token
	token, _, err := jwt.CreateToken(claim)
	if err != nil {
		return "", err
	}

	return token, nil
}
