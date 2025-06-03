package handlers

import (
	"ecomm-go/app/views/components/iomCategory"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowCategoryManagement(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := h.Queries.GetCategories(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomCategory.CategoryManagement(categories))
}

func (h DataHandler) ShowEditCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryID, err := ParseToInt32(c.FormValue("id"))
	if err != nil {
		return HandleError(c, err)
	}
	category, err := h.Queries.GetCategoryByID(ctx,
		categoryID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomCategory.EditCategory(category))
}

func (h DataHandler) EditCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryID, err := ParseToInt32(c.FormValue("category-id"))
	if err != nil {
		return HandleError(c, err)
	}
	categoryName := c.FormValue("category-name")
	err = h.Queries.UpdateCategory(ctx,
		goqueries.UpdateCategoryParams{
			ID:   categoryID,
			Name: categoryName,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowCategoryManagement(c)
}

func (h DataHandler) ShowAddCategory(c echo.Context) error {
	return Render(c, iomCategory.AddCategory())
}

func (h DataHandler) AddCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryName := c.FormValue("category-name")

	err := h.Queries.InsertCategory(ctx, categoryName)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowCategoryManagement(c)
}

func (h DataHandler) DeleteCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryID, err := ParseToInt32(c.FormValue("category-id"))
	if err != nil {
		return HandleError(c, err)
	}
	err = h.Queries.DeleteCategoryByID(ctx,
		categoryID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowCategoryManagement(c)
}
