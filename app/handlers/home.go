package handlers

import (
	"database/sql"
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowHome(c echo.Context) error {
	user, err := GetUserFromSession(h, c)

	ctx := c.Request().Context()
	filteredProducts, err := h.Queries.GetProductsByCategoriesAndName(ctx,
		goqueries.GetProductsByCategoriesAndNameParams{
			CategoryIds: []int32{},
			NameFilter:  sql.NullString{String: "", Valid: true},
			SpecialIds:  []int32{2},
			PriceFrom:   0,
			PriceTo:     999999,
			Offsetvar:   0,
			Limitvar:    10,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}

	component := pages.Home(filteredProducts)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}
