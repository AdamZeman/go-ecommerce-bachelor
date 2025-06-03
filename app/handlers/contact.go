package handlers

import (
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowContact(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	component := pages.Contact()
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}
