package handlers

import (
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h DataHandler) ShowIOM(c echo.Context) error {

	ctx := c.Request().Context()
	products, err := h.Queries.GetProducts(ctx)
	if err != nil {
		return HandleError(c, err)
	}

	user, _ := GetUserFromSession(h, c)
	component := pages.IOM(products)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func (h DataHandler) RemovePopup(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
