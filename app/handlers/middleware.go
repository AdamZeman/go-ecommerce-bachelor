package handlers

import (
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (h DataHandler) AuthMiddleware(requiredRole int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := GetUserFromSession(h, c)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			if user.Role < requiredRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Forbidden: Insufficient permissions"})
			}
			return next(c)
		}
	}
}

func (h DataHandler) MessengerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			user, err := GetUserFromSession(h, c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			strIdParam := c.Param("id")
			numIdParam, _ := ParseToInt32(strIdParam)

			isPermitted, err := h.Queries.CountConvByUserConv(ctx, goqueries.CountConvByUserConvParams{
				UserID:         user.ID,
				ConversationID: numIdParam,
			})

			if isPermitted < 1 {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Forbidden: Insufficient permissions"})
			}

			return next(c)
		}
	}
}

func CORSConfigMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderAuthorization, echo.HeaderContentType},
		AllowCredentials: true,
		MaxAge:           3600,
	})
}
