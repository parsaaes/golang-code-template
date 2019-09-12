package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rashadansari/golang-code-template/model"
	"github.com/rashadansari/golang-code-template/request"
	"github.com/rashadansari/golang-code-template/response"
	"github.com/sirupsen/logrus"
)

type TemplateHandler struct {
	TemplateRepo model.TemplateRepo
}

func (t TemplateHandler) Create(c echo.Context) error {
	r := new(request.CreateTemplateRequest)

	if err := c.Bind(r); err != nil {
		logrus.Errorf("template handler bind: %s", err.Error())
		return echo.ErrBadRequest
	}

	if err := r.Validate(); err != nil {
		logrus.Errorf("template handler validate: %s", err.Error())
		return echo.ErrBadRequest
	}

	template := &model.Template{
		Name: r.Name,
	}

	if err := t.TemplateRepo.Create(template); err != nil {
		logrus.Errorf("template handler failed to create template: %s", err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, response.CreateTemplateResponse{
		ID:   template.ID,
		Name: template.Name,
	})
}
