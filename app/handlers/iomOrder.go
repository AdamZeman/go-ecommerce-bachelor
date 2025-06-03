package handlers

import (
	"ecomm-go/app/views/components/iomOrder"
	"ecomm-go/db/goqueries"

	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowOrderManagement(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := h.Queries.GetOrdersFillUser(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomOrder.OrderManagement(orders))
}

func (h DataHandler) ShowEditOrder(c echo.Context) error {
	ctx := c.Request().Context()

	orderID, _ := ParseToInt32(c.FormValue("id"))
	orderItemsFilled, err := h.Queries.GetOrderItemsByOrderID(ctx,
		orderID,
	)
	if err != nil {
		return HandleError(c, err)
	}

	order, err := h.Queries.GetOrderByID(ctx, orderID)
	if err != nil {
		return HandleError(c, err)
	}

	return Render(c, iomOrder.EditOrder(order, orderItemsFilled))
}

func (h DataHandler) EditOrder(c echo.Context) error {
	ctx := c.Request().Context()
	orderID, _ := ParseToInt32(c.FormValue("order-id"))
	radioStatus := c.FormValue("radio-status")

	_, err := h.Queries.UpdateOrderStatus(ctx,
		goqueries.UpdateOrderStatusParams{
			ID:     orderID,
			Status: radioStatus,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowOrderManagement(c)
}
