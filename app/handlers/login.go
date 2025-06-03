package handlers

import (
	"context"
	"database/sql"
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func (h DataHandler) ShowLogin(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	component := pages.Login()
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func toCookieUser(userGoth goth.User, databaseUser goqueries.User) types.CookieUser {
	return types.CookieUser{
		Provider:    userGoth.Provider,
		Email:       userGoth.Email,
		Name:        userGoth.Name,
		UserID:      userGoth.UserID,
		AvatarURL:   userGoth.AvatarURL,
		AccessToken: userGoth.AccessToken,
		Role:        int(databaseUser.Role),
		ID:          databaseUser.ID,
		IsSigned:    true,
	}
}

func (h DataHandler) AuthCallback(c echo.Context) error {

	provider := c.Param("provider")
	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", provider))
	res := c.Response().Writer

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userFromDatabase, err := h.Queries.UpsertUser(ctx, goqueries.UpsertUserParams{
		GoogleID:  user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: sql.NullString{String: user.AvatarURL, Valid: true},
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	cookieUser := toCookieUser(user, userFromDatabase)
	session, _ := h.Store.Get(req, "session-name")
	session.Values["user"] = cookieUser

	if err := session.Save(req, res); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save session")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h DataHandler) Logout(c echo.Context) error {
	provider := c.Param("provider")
	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", provider))

	session, _ := h.Store.Get(req, "session-name")
	session.Values["user"] = nil
	session.Options.MaxAge = -1

	if err := session.Save(req, c.Response().Writer); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete session")
	}

	if err := gothic.Logout(c.Response().Writer, req); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to logout")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h DataHandler) Auth(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return c.String(http.StatusBadRequest, "Provider not specified")
	}

	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", provider))
	res := c.Response().Writer

	gothic.BeginAuthHandler(res, req)

	return nil
}
