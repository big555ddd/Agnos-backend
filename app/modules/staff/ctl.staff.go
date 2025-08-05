package staff

import (
	"app/app/message"
	staffdto "app/app/modules/staff/dto"
	"app/app/response"
	"app/internal/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service ServiceInterface
}

func NewController(svc *Service) *Controller {
	return &Controller{
		Service: svc,
	}
}

func (c *Controller) Create(ctx *gin.Context) {
	req := new(staffdto.CreateStaffRequest)
	if err := ctx.Bind(req); err != nil {
		logger.Err(err)
		response.BadRequest(ctx, message.InvalidRequest, nil)
		return
	}
	err := c.Service.Create(ctx, req)
	if err != nil {
		logger.Err(err)
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil)
}

func (c *Controller) Login(ctx *gin.Context) {
	req := new(staffdto.LoginStaffRequest)
	if err := ctx.Bind(req); err != nil {
		logger.Err(err)
		response.BadRequest(ctx, message.InvalidRequest, nil)
		return
	}
	token, err := c.Service.Login(ctx, req)
	if err != nil {
		logger.Err(err)
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, token)
}
