package patient

import (
	"app/app/helper"
	"app/app/message"
	patientdto "app/app/modules/patient/dto"
	"app/app/response"
	"app/internal/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service ServiceInterface
}

func NewController(svc ServiceInterface) *Controller {
	return &Controller{
		Service: svc,
	}
}

func (c *Controller) GetPatient(ctx *gin.Context) {
	id := new(patientdto.GetPatientByIdRequest)
	if err := ctx.BindUri(id); err != nil {
		logger.Err(err)
		response.BadRequest(ctx, message.InvalidRequest, nil)
		return
	}
	resp, err := c.Service.GetPatient(ctx, id.ID)
	if err != nil {
		logger.Err(err)
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	defer resp.Body.Close()
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

func (c *Controller) List(ctx *gin.Context) {
	req := patientdto.ListPatientRequest{
		Page:    1,
		Size:    10,
		OrderBy: "asc",
		SortBy:  "created_at",
	}
	if err := ctx.BindQuery(&req); err != nil {
		logger.Err(err)
		response.BadRequest(ctx, message.InvalidRequest, nil)
		return
	}
	user, _ := helper.GetUserByToken(ctx)
	data, total, err := c.Service.List(ctx, &req, user.Data.Hospital)
	if err != nil {
		logger.Err(err)
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	response.SuccessWithPaginate(ctx, data, req.Page, req.Size, total)
}
