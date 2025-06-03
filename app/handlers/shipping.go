package handlers

import (
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
)

func (h DataHandler) ShowShipping(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()
	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	totalPrice := TotalPriceCalcBasket(itemProductVariants)

	component := pages.Shipping(totalPrice, h.StripePublicKey, user.Email, user.Name)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func (h DataHandler) ValidatePhone(c echo.Context) error {
	phone := c.FormValue("phone-number")
	pattern := `^\+?[0-9]{7,15}$`

	matched, err := regexp.MatchString(pattern, phone)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "")
	}
	if matched {
		return c.HTML(http.StatusOK, `<div style="color: green;">✅ Valid phone number</div>`)
	} else {
		return c.HTML(http.StatusOK, `<div style="color: red;">❌ Invalid phone number</div>`)
	}
}
