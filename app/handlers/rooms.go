package handlers

import (
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowRooms(c echo.Context) error {
	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}
	ctx := c.Request().Context()
	conversations, err := h.Queries.GetConversationsByUserId(ctx, user.Email)
	if err != nil {
		return HandleError(c, err)
	}

	firstConversation := goqueries.Conversation{
		ID:   0,
		Name: "No Conversations",
	}

	if len(conversations) > 0 {
		firstConversation = conversations[0]
	}

	messages, err := h.Queries.GetMessagesByConversationId(ctx, firstConversation.ID)
	if err != nil {
		return HandleError(c, err)
	}

	for i, msg := range messages {
		if msg.Email == user.Email {
			messages[i].Email = "current"
		}
	}

	component := pages.Rooms(conversations, messages, firstConversation)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}
