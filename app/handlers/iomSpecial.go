package handlers

import (
	"ecomm-go/app/views/components/iomSpecial"
	"ecomm-go/db/goqueries"

	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowSpecialManagement(c echo.Context) error {
	ctx := c.Request().Context()
	specialTags, err := h.Queries.GetSpecial(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomSpecial.SpecialManagement(specialTags))
}

func (h DataHandler) ShowEditSpecial(c echo.Context) error {
	ctx := c.Request().Context()

	specialID, _ := ParseToInt32(c.FormValue("id"))
	specialTags, err := h.Queries.GetSpecialByID(ctx,
		specialID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomSpecial.EditSpecial(specialTags))
}

func (h DataHandler) EditSpecial(c echo.Context) error {
	ctx := c.Request().Context()
	specialID, _ := ParseToInt32(c.FormValue("special-id"))
	specialName := c.FormValue("special-name")

	err := h.Queries.UpdateSpecial(ctx,
		goqueries.UpdateSpecialParams{
			ID:   specialID,
			Name: specialName,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowSpecialManagement(c)
}

func (h DataHandler) ShowAddSpecial(c echo.Context) error {
	return Render(c, iomSpecial.AddSpecial())
}

func (h DataHandler) AddSpecial(c echo.Context) error {
	ctx := c.Request().Context()
	specialName := c.FormValue("special-name")

	err := h.Queries.InsertSpecial(ctx, specialName)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowSpecialManagement(c)
}

func (h DataHandler) DeleteSpecial(c echo.Context) error {
	ctx := c.Request().Context()
	specialID, _ := ParseToInt32(c.FormValue("special-id"))

	err := h.Queries.DeleteSpecialByID(ctx,
		specialID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowSpecialManagement(c)
}
