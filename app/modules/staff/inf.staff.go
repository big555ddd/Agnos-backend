package staff

import (
	staffdto "app/app/modules/staff/dto"
	"context"
)

// ServiceInterface defines the interface for staff service operations
type ServiceInterface interface {
	Create(ctx context.Context, req *staffdto.CreateStaffRequest) error
	Login(ctx context.Context, req *staffdto.LoginStaffRequest) (string, error)
	ExistUsername(ctx context.Context, username string) (bool, error)
}

var _ ServiceInterface = (*Service)(nil)
