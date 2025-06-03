package handlers

import (
	"ecomm-go/app/views/components/iomPendingOrder"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowPendingOrderManagement(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := h.Queries.GetOrdersFillUserByStatus(ctx, "pending")
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomPendingOrder.PendingOrderManagement(orders))
}

func (h DataHandler) EditPendingOrder(c echo.Context) error {
	ctx := c.Request().Context()
	orderID, _ := ParseToInt32(c.FormValue("id"))

	_, err := h.Queries.UpdateOrderStatus(ctx,
		goqueries.UpdateOrderStatusParams{
			ID:     orderID,
			Status: "done",
		},
	)
	if err != nil {
		return HandleError(c, err)
	}

	return h.ShowPendingOrderManagement(c)
}
