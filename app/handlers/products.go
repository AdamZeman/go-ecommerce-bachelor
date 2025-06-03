package handlers

import (
	"database/sql"
	"ecomm-go/app/views/components"
	"ecomm-go/app/views/layouts"
	"ecomm-go/app/views/pages"
	"ecomm-go/db/goqueries"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var MaxPerPage int32 = 30

func getDistinctValues(items []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, item := range items {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func mapFunc[T any, U any](list []T, fn func(T) U) []U {
	var result []U
	for _, item := range list {
		result = append(result, fn(item))
	}
	return result
}

func (h DataHandler) ShowProducts(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := GetUserFromSession(h, c)

	filteredProducts, err := h.Queries.GetProductsByCategoriesAndName(ctx,
		goqueries.GetProductsByCategoriesAndNameParams{
			UserID:      user.ID,
			CategoryIds: []int32{},
			NameFilter:  sql.NullString{String: "", Valid: true},
			SpecialIds:  []int32{},
			PriceFrom:   0,
			PriceTo:     999999,
			Offsetvar:   0,
			Limitvar:    MaxPerPage,
		},
	)

	itemsCategories, err := h.Queries.GetCategories(ctx)
	if err != nil {
		return HandleError(c, err)
	}

	itemsSpecial, err := h.Queries.GetSpecial(ctx)
	if err != nil {
		return HandleError(c, err)
	}

	component := pages.Products(filteredProducts, itemsCategories, itemsSpecial, user.IsSigned)
	csrfToken := c.Get("csrf").(string)
	return Render(c, layouts.Base(user, component, csrfToken))
}

func (h DataHandler) ShowSingleProduct(c echo.Context) error {

	strIdParam := c.Param("id")
	numIdParam, _ := ParseToInt32(strIdParam)
	ctx := c.Request().Context()
	user, _ := GetUserFromSession(h, c)

	productVariants, err := h.Queries.ProductVariantsByProductId(ctx, numIdParam)
	option1Values := getDistinctValues(
		mapFunc(productVariants, func(item goqueries.ProductVariant) string {
			return item.Option1Value
		}),
	)
	option2Values := getDistinctValues(
		mapFunc(productVariants, func(item goqueries.ProductVariant) string {
			return item.Option2Value
		}),
	)

	itemProduct, err := h.Queries.GetProductById(ctx, numIdParam)
	if err != nil {
		return HandleError(c, err)
	}

	favouriteItems, err := h.Queries.GetFavouriteByProductId(ctx, goqueries.GetFavouriteByProductIdParams{
		ProductID: numIdParam,
		UserID:    user.ID,
	})
	if err != nil {
		return HandleError(c, err)
	}

	reviews, err := h.Queries.GetReviewsFillUserByProductID(ctx, numIdParam)
	if err != nil {
		return HandleError(c, err)
	}

	var recomendedSpecialsId int32
	if itemProduct.Special.Valid {
		recomendedSpecialsId = itemProduct.Special.Int32
	} else {
		recomendedSpecialsId = 1
	}

	itemsProducts, err := h.Queries.GetProductsByCategoriesAndName(ctx,
		goqueries.GetProductsByCategoriesAndNameParams{
			UserID:      user.ID,
			CategoryIds: []int32{},
			NameFilter:  sql.NullString{String: "", Valid: true},
			SpecialIds:  []int32{recomendedSpecialsId},
			PriceFrom:   0,
			PriceTo:     999999,
			Offsetvar:   0,
			Limitvar:    10,
		},
	)

	itemCategories, err := h.Queries.GetCategoriesByProductID(ctx, numIdParam)
	if err != nil {
		return HandleError(c, err)
	}

	isFavourite := len(favouriteItems) != 0

	component := pages.SingleProduct(user, itemProduct, option1Values, option2Values, isFavourite, reviews, itemsProducts, itemCategories)
	csrfToken := c.Get("csrf").(string)

	return Render(c, layouts.Base(user, component, csrfToken))
}

type ProductFilterForm struct {
	Name       string  `form:"search"`
	PriceFrom  int32   `form:"price-from"  validate:"omitempty,gt=-1"`
	PriceTo    int32   `form:"price-to"    validate:"omitempty,gt=-1"`
	Categories []int32 `form:"category"    validate:"dive,gt=0"`
	Specials   []int32 `form:"special"     validate:"dive,gt=0"`
	Page       int32   `form:"page"        validate:"omitempty,gt=-1"`
}

func (h DataHandler) handleFilteredProducts(c echo.Context, reset bool) error {
	ctx := c.Request().Context()
	user, _ := GetUserFromSession(h, c)
	var form ProductFilterForm
	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}

	if reset {
		form.Page = 0
	}
	if form.PriceTo == 0 {
		form.PriceTo = int32(999999)
	}
	filteredProducts, err := h.Queries.GetProductsByCategoriesAndName(ctx,
		goqueries.GetProductsByCategoriesAndNameParams{
			UserID:      user.ID,
			CategoryIds: form.Categories,
			NameFilter:  sql.NullString{String: form.Name, Valid: true},
			SpecialIds:  form.Specials,
			PriceFrom:   form.PriceFrom * 100,
			PriceTo:     form.PriceTo * 100,
			Offsetvar:   form.Page * MaxPerPage,
			Limitvar:    MaxPerPage,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	if len(filteredProducts) == 0 {
		return c.HTML(http.StatusOK, `
			<div id="more-products-loader" hx-swap-oob="delete"></div>
		`)
	}
	component := components.ProductList(filteredProducts, form.Page+1, true)
	return Render(c, component)
}

func (h DataHandler) ShowFilteredProducts(c echo.Context) error {
	return h.handleFilteredProducts(c, true)
}

func (h DataHandler) ShowMoreProducts(c echo.Context) error {
	return h.handleFilteredProducts(c, false)
}

func (h DataHandler) UpdateProductInfo(c echo.Context) error {

	option1value := c.FormValue("option1value")
	option2value := c.FormValue("option2value")
	amount := c.FormValue("amount")
	amountFloat64, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		return HandleError(c, err)
	}
	amountFloat32 := int32(amountFloat64)
	strIdParam := c.FormValue("product_id")
	numIdParam, _ := ParseToInt32(strIdParam)
	ctx := c.Request().Context()

	arg := goqueries.GetVariantByOptionsParams{
		ProductID:    numIdParam,
		Option1Value: option1value,
		Option2Value: option2value,
	}

	variant, err := h.Queries.GetVariantByOptions(ctx, arg)
	if err != nil {
		return HandleError(c, err)
	}

	price := variant.Price * amountFloat32
	stock := variant.StockQuantity

	responseHTML := fmt.Sprintf(`
		<h1 id="total-price" hx-swap-oob="true">$%.2f</h1>
		<h1 id="stock-info" hx-swap-oob="true">%d left</h1>
		
	`, float64(price)/100, stock)

	return c.HTML(http.StatusOK, responseHTML)
}
