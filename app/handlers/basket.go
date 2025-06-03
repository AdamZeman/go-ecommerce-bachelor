package handlers

import (
	"ecomm-go/app/views/components"
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TotalPriceCalcBasket(basketItems []goqueries.GetBasketItemsByUserIdRow) int32 {
	count := int32(0)
	for _, basketItem := range basketItems {
		count += basketItem.Price * basketItem.Quantity
	}
	return count
}

func (h DataHandler) ShowBasket(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()
	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	totalPrice := TotalPriceCalcBasket(itemProductVariants)
	component := pages.Basket(itemProductVariants, totalPrice, user)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

type AddToBasketForm struct {
	Option1value string `form:"option1value" validate:"required"`
	Option2value string `form:"option2value" validate:"required" `
	Amount       int32  `form:"amount" validate:"required,gt=0"`
	ProductID    int32  `form:"product_id" validate:"required,gt=0"`
}

func (h DataHandler) AddToBasket(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()

	var addF AddToBasketForm
	if err := c.Bind(&addF); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&addF); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}
	variant, err := h.Queries.GetVariantByOptions(ctx, goqueries.GetVariantByOptionsParams{
		ProductID:    addF.ProductID,
		Option1Value: addF.Option1value,
		Option2Value: addF.Option2value,
	})
	if err != nil {
		return HandleError(c, err)
	}
	_, err = h.Queries.InsertBasketItem(ctx, goqueries.InsertBasketItemParams{
		UserID:           user.ID,
		ProductVariantID: variant.ID,
		Quantity:         addF.Amount,
	})
	if err != nil {
		return HandleError(c, err)
	}

	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	responseHTML := fmt.Sprintf(`
		<button
            class="w-1/2 text-center bg-hunter-green-700 text-white hover:opacity-80 transform transition-all duration-300 rounded-full p-2 border-hunter-green-700 border"
			id="addToBasketBtn"
			hx-swap-oob="true"	
		>
			Added
		</button>

		<span
			id="basket-counter"
			class="text-white absolute text-xs text-center top-0"
			hx-get="/api/update-basket-size"
			hx-target="#basket-counter"
			hx-trigger="load"
			hx-swap-oob="true"
		>
			%d
		</span>

		<div
			id="basket-reloader"
			class="hidden"
			hx-get="/api/get-quick-basket"
			hx-target="#quick-basket"
			hx-trigger="load"
			hx-swap-oob="true"
		>
		</div>
		
	`, len(itemProductVariants))

	return c.HTML(http.StatusOK, responseHTML)
}

func (h DataHandler) RemoveFromBasket(c echo.Context) error {
	basketItemId, err := ParseToInt32(c.FormValue("basketItemId"))
	if err != nil {
		return HandleError(c, err)
	}
	ctx := c.Request().Context()
	removedItem, err := h.Queries.RemoveFromBasket(ctx, basketItemId)
	if err != nil {
		return HandleError(c, err)
	}
	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, removedItem.UserID)
	if err != nil {
		return HandleError(c, err)
	}

	divItemId := fmt.Sprintf("basketItem-%d", basketItemId)
	totalPrice := TotalPriceCalcBasket(itemProductVariants)
	responseHTML := fmt.Sprintf(`
		<h1 id="total-price" hx-swap-oob="true">$%.2f</h1>
		<div id="%s" hx-swap-oob="delete"></div>
		<span
			id="basket-counter"
			class="text-white absolute text-xs text-center top-0"
			hx-get="/api/update-basket-size"
			hx-target="#basket-counter"
			hx-trigger="load"
		>
			%d
		</span>
	`, float64(totalPrice/100), divItemId, len(itemProductVariants))

	return c.HTML(http.StatusOK, responseHTML)
}

func (h DataHandler) UpdateBasketSize(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()
	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	return c.String(http.StatusOK, fmt.Sprintf("%d", len(itemProductVariants)))
}

func (h DataHandler) ShowQuickBasket(c echo.Context) error {

	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()
	itemProductVariants, err := h.Queries.GetBasketItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	totalPrice := TotalPriceCalcBasket(itemProductVariants)
	component := components.QuickBasket(itemProductVariants, totalPrice)

	return Render(c, component)
}
