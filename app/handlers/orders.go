package handlers

import (
	"ecomm-go/app/views/components"
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func FilterByOrderId(list []goqueries.GetOrderItemsByUserIdRow, id int32) []goqueries.GetOrderItemsByUserIdRow {
	newList := make([]goqueries.GetOrderItemsByUserIdRow, 0)
	for _, item := range list {
		if item.OrderID == id {
			newList = append(newList, item)
		}
	}
	return newList
}

func (h DataHandler) ShowOrders(c echo.Context) error {
	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}
	ctx := c.Request().Context()
	orderItems, err := h.Queries.GetOrderItemsByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	ordersDB, err := h.Queries.GetOrdersByUserId(ctx, user.ID)
	if err != nil {
		return HandleError(c, err)
	}
	var ordersFilled []types.OrderFilled
	for _, orderDB := range ordersDB {
		orderFilled := types.OrderFilled{
			Order:      orderDB,
			OrderItems: FilterByOrderId(orderItems, orderDB.ID),
		}
		ordersFilled = append(ordersFilled, orderFilled)
	}
	component := pages.Orders(ordersFilled)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func (h DataHandler) OpenIssue(c echo.Context) error {
	orderId := c.FormValue("id")
	orderIdNum, _ := ParseToInt32(orderId)
	ctx := c.Request().Context()

	conversations, err := h.Queries.GetOpenConversationById(ctx, goqueries.GetOpenConversationByIdParams{
		OrderID: orderIdNum,
		Status:  "issue",
	})

	responseHTML := fmt.Sprintf(`
		<button
			disabled
			type="button"
		   	class="rounded-full px-2 py-1 bg-yellow-600 text-white hover:opacity-80 cursor-pointer"
		   	id="orderDiv-%d"}
			hx-swap-oob=true
		>
			Issue Opened
		</button>
		
	`, orderIdNum)

	order, err := h.Queries.UpdateOrderStatus(ctx,
		goqueries.UpdateOrderStatusParams{
			ID:     orderIdNum,
			Status: "issue",
		},
	)
	if err != nil {
		return HandleError(c, err)
	}

	if len(conversations) > 0 {
		if err != nil {
			return HandleError(c, err)
		}
		return c.HTML(http.StatusOK, responseHTML)
	}

	insertedConversation, err := h.Queries.InsertConversation(ctx, goqueries.InsertConversationParams{
		OrderID: orderIdNum,
		Name:    fmt.Sprint("#", orderIdNum),
	})
	if err != nil {
		return HandleError(c, err)
	}

	_, err = h.Queries.InsertConversationUser(ctx, goqueries.InsertConversationUserParams{
		ConversationID: insertedConversation.ID,
		UserID:         order.UserID,
	})
	if err != nil {
		return HandleError(c, err)
	}

	_, err = h.Queries.InsertConversationUser(ctx, goqueries.InsertConversationUserParams{
		ConversationID: insertedConversation.ID,
		UserID:         2,
	})
	if err != nil {
		return HandleError(c, err)
	}

	_, err = h.Queries.InsertMessage(ctx, goqueries.InsertMessageParams{
		ConversationID: insertedConversation.ID,
		Email:          "zeman5698@gmail.com",
		Content:        fmt.Sprint("Initiated new issue regarding order number #", orderIdNum),
	})
	if err != nil {
		return HandleError(c, err)
	}

	_, err = h.Queries.UpdateOrderStatus(ctx, goqueries.UpdateOrderStatusParams{
		ID:     orderIdNum,
		Status: "issue",
	})
	if err != nil {
		return HandleError(c, err)
	}

	return c.HTML(http.StatusOK, responseHTML)

}

func (h DataHandler) AddReview(c echo.Context) error {

	ctx := c.Request().Context()
	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}

	ProductID, _ := ParseToInt32(c.FormValue("product-id"))
	Rating, _ := ParseToInt32(c.FormValue("rating"))
	Content := c.FormValue("content")

	_, err = h.Queries.InsertReview(ctx,
		goqueries.InsertReviewParams{
			SenderID:  user.ID,
			ProductID: ProductID,
			Rating:    Rating,
			Content:   Content,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return c.HTML(http.StatusOK, "")
}

func (h DataHandler) ShowAddReview(c echo.Context) error {
	ProductID, _ := ParseToInt32(c.FormValue("product-id"))
	return Render(c, components.ReviewModal(ProductID))
}
