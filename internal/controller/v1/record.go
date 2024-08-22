package v1

import (
	"net/http"
	"service-template/internal/service"

	"github.com/labstack/echo/v4"
)

type RecordRoutes struct {
	serviceRecord service.Record
}

func newRecordRoutes(g *echo.Group, serviceRecord service.Record) {
	r := &RecordRoutes{
		serviceRecord: serviceRecord,
	}

	g.GET("/get", r.getById)
	g.GET("/list", r.getList)
}

type getByIdInput struct {
	Id int `json:"id" validate:"required"`
}

func (r *RecordRoutes) getById(c echo.Context) error {
	var input getByIdInput

	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	record, err := r.serviceRecord.GetById(c.Request().Context(), input.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, record)
}

func (r *RecordRoutes) getList(c echo.Context) error {
	records, err := r.serviceRecord.GetList(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, records)
}
