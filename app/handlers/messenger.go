package handlers

import (
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowMessengerAlone(c echo.Context) error {
	strIdParam := c.Param("id")
	numIdParam, _ := ParseToInt32(strIdParam)
	convName := c.FormValue("conversation-name")

	firstConversation := goqueries.Conversation{
		ID:   numIdParam,
		Name: convName,
	}
	ctx := c.Request().Context()
	messages, err := h.Queries.GetMessagesByConversationId(ctx, numIdParam)
	if err != nil {
		return HandleError(c, err)
	}

	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}

	for i, msg := range messages {
		if msg.Email == user.Email {
			messages[i].Email = "current"
		}
	}
	user, _ = GetUserFromSession(h, c)
	component := pages.Messenger(messages, firstConversation)
	return Render(c, component)
}

func (h DataHandler) ShowMessengerPage(c echo.Context) error {
	strIdParam := c.Param("id")
	numIdParam, _ := ParseToInt32(strIdParam)
	convName := c.FormValue("conversation-name")

	firstConversation := goqueries.Conversation{
		ID:   numIdParam,
		Name: convName,
	}
	ctx := c.Request().Context()
	messages, err := h.Queries.GetMessagesByConversationId(ctx, numIdParam)
	if err != nil {
		return HandleError(c, err)
	}

	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}

	for i, msg := range messages {
		if msg.Email == user.Email {
			messages[i].Email = "current"
		}
	}

	user, _ = GetUserFromSession(h, c)
	component := pages.Messenger(messages, firstConversation)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}
