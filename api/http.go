package api

import (
	"github.com/go-fuego/fuego"
	"github.com/joaquincorimayo/poc-fuego/domain/model"
	"github.com/joaquincorimayo/poc-fuego/domain/usecase"
)

type Handler struct {
	reportService usecase.Service
}

func NewReportHandler(reportService usecase.Service) *Handler {
	return &Handler{reportService: reportService}
}

func (h *Handler) Get(c *fuego.ContextNoBody) (*model.Report, error) {
	code := c.PathParam("code")

	report, err := h.reportService.FindByCode(code)
	if err != nil {
		return &model.Report{}, err
	}

	return report, nil
}

func (h *Handler) Post(c *fuego.ContextWithBody[model.Report]) (model.ReportCreatedResponse, error) {
	body, err := c.Body()
	if err != nil {
		return model.ReportCreatedResponse{}, err
	}

	err = h.reportService.Save(&body)
	if err != nil {
		return model.ReportCreatedResponse{}, err
	}

	return model.ReportCreatedResponse{
		Code:        body.Code,
		DateCreated: body.DateCreated,
	}, nil
}
