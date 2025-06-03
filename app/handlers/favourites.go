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

func (h DataHandler) ShowFavourites(c echo.Context) error {
	user, err := GetUserFromSession(h, c)
	if err != nil {
		return HandleError(c, err)
	}

	ctx := c.Request().Context()
	favouritesItems, err := h.Queries.GetFavouritesById(ctx, user.Email)
	if err != nil {
		return HandleError(c, err)
	}

	component := pages.Favourites(favouritesItems)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func (h DataHandler) AddToFavourites(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	strIdParam := c.FormValue("product_id")
	numIdParam, err := ParseToInt32(strIdParam)
	if err != nil {
		return HandleError(c, err)
	}
	ctx := c.Request().Context()

	favouriteItems, err := h.Queries.GetFavouriteByProductId(ctx, goqueries.GetFavouriteByProductIdParams{
		ProductID: numIdParam,
		UserID:    user.ID,
	})
	if err != nil {
		return HandleError(c, err)
	}

	var favouriteClasses string
	if len(favouriteItems) == 0 {
		favouriteClasses = "fill-timber-green-700 group-hover/love:fill-transparent "
		insertItem := goqueries.InsertItemToFavouritesParams{
			UserID:    user.ID,
			ProductID: numIdParam,
		}
		_, err = h.Queries.InsertItemToFavourites(ctx, insertItem)
		if err != nil {
			return HandleError(c, err)
		}
	} else {
		favouriteClasses = "group-hover/love:fill-timber-green-700 fill-transparent"
		err = h.Queries.RemoveItemFromFavourites(ctx, goqueries.RemoveItemFromFavouritesParams{
			ProductID: numIdParam,
			UserID:    user.ID,
		})
		if err != nil {
			return HandleError(c, err)
		}
	}

	return Render(c, components.LoveIcon(favouriteClasses, numIdParam))
}

func (h DataHandler) RemoveFromFavourites(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)

	favouritesItemProductId, err := ParseToInt32(c.FormValue("favouritesItemProductId"))
	if err != nil {
		return HandleError(c, err)
	}
	ctx := c.Request().Context()
	err = h.Queries.RemoveItemFromFavourites(ctx, goqueries.RemoveItemFromFavouritesParams{
		ProductID: favouritesItemProductId,
		UserID:    user.ID,
	})
	if err != nil {
		return HandleError(c, err)
	}

	divItemId := fmt.Sprintf("favouriteItem-%d", favouritesItemProductId)
	responseHTML := fmt.Sprintf(`
		<div id="%s" hx-swap-oob="delete"></div>
	`, divItemId)

	return c.HTML(http.StatusOK, responseHTML)
}
