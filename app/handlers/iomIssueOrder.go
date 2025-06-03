package handlers

import (
	"ecomm-go/app/views/components/iomIssueOrder"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowIssueOrderManagement(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := h.Queries.GetOrdersFillUserConvByStatus(ctx, "issue")
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomIssueOrder.IssueOrderManagement(orders))
}

func (h DataHandler) EditIssueOrder(c echo.Context) error {
	ctx := c.Request().Context()
	orderID, err := ParseToInt32(c.FormValue("id"))
	if err != nil {
		return HandleError(c, err)
	}
	_, err = h.Queries.UpdateOrderStatus(ctx,
		goqueries.UpdateOrderStatusParams{
			ID:     orderID,
			Status: "done",
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowIssueOrderManagement(c)
}
