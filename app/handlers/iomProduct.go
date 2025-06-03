package handlers

import (
	"database/sql"
	"ecomm-go/app/views/components/iomProduct"
	"ecomm-go/db/goqueries"
	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowAddProduct(c echo.Context) error {
	ctx := c.Request().Context()
	special, err := h.Queries.GetSpecial(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomProduct.AddProduct(special))
}

func (h DataHandler) ShowEditProduct(c echo.Context) error {
	ctx := c.Request().Context()

	productID, _ := ParseToInt32(c.FormValue("id"))
	product, err := h.Queries.GetProductById(ctx,
		productID,
	)
	if err != nil {
		return HandleError(c, err)
	}

	special, err := h.Queries.GetSpecial(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomProduct.EditProduct(product, special))
}

func (h DataHandler) ShowProductManagement(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := h.Queries.GetProducts(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomProduct.ProductManagement(products))
}

type AddProductForm struct {
	Name             string `form:"product-name"      validate:"required,min=2,max=100"`
	Description      string `form:"product-desc"      validate:"omitempty,min=0"`
	Price            int32  `form:"product-price"     validate:"required,gt=0"`
	Option1Name      string `form:"product-option1-name" validate:"omitempty,max=50"`
	Option2Name      string `form:"product-option2-name" validate:"omitempty,max=50"`
	ProductSpecialID *int32 `form:"product-special-id"  validate:"omitempty"`
}

func (h DataHandler) AddProduct(c echo.Context) error {
	ctx := c.Request().Context()
	var form AddProductForm

	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}

	special := sql.NullInt32{Valid: false}
	if *form.ProductSpecialID != 0 {
		special = sql.NullInt32{Valid: true, Int32: *form.ProductSpecialID}
	}

	err := h.Queries.InsertProduct(ctx,
		goqueries.InsertProductParams{
			Name:        form.Name,
			Description: form.Description,
			Price:       form.Price,
			Option1Name: form.Option1Name,
			Option2Name: form.Option2Name,
			Special:     special,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowProductManagement(c)
}

type EditProductForm struct {
	ID               int32  `form:"product-id"      validate:"required"`
	Name             string `form:"product-name"      validate:"required,min=2,max=100"`
	Description      string `form:"product-desc"      validate:"omitempty,min=0"`
	Price            int32  `form:"product-price"     validate:"required,gt=0"`
	Option1Name      string `form:"product-option1-name" validate:"omitempty,max=50"`
	Option2Name      string `form:"product-option2-name" validate:"omitempty,max=50"`
	ProductSpecialID *int32 `form:"product-special-id"  validate:"omitempty"`
}

func (h DataHandler) EditProduct(c echo.Context) error {
	ctx := c.Request().Context()
	var form EditProductForm
	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}

	special := sql.NullInt32{Valid: false}
	if *form.ProductSpecialID != 0 {
		special = sql.NullInt32{Valid: true, Int32: *form.ProductSpecialID}
	}

	err := h.Queries.UpdateProduct(ctx,
		goqueries.UpdateProductParams{
			ID:          form.ID,
			Name:        form.Name,
			Description: form.Description,
			Price:       form.Price,
			Option1Name: form.Option1Name,
			Option2Name: form.Option2Name,
			Special:     special,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowProductManagement(c)
}

func (h DataHandler) DeleteProduct(c echo.Context) error {
	ctx := c.Request().Context()
	productID, _ := ParseToInt32(c.FormValue("product-id"))

	err := h.Queries.DeleteProductByID(ctx,
		productID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowProductManagement(c)
}

func (h DataHandler) ShowProductCategories(c echo.Context) error {
	ctx := c.Request().Context()
	productID, _ := ParseToInt32(c.FormValue("product-id"))

	productCategories, err := h.Queries.GetCategoriesFilteredProductID(ctx, productID)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomProduct.ProductCategory(productID, productCategories))

}

func (h DataHandler) CategorySwitch(c echo.Context) error {
	ctx := c.Request().Context()
	productID, _ := ParseToInt32(c.FormValue("product-id"))
	categoryID, _ := ParseToInt32(c.FormValue("category-id"))
	isTrue, _ := ParseToBool(c.FormValue("is-true"))

	if isTrue {
		err := h.Queries.DeleteProductCategoryByIDs(ctx,
			goqueries.DeleteProductCategoryByIDsParams{
				ProductID:  productID,
				CategoryID: categoryID,
			},
		)
		if err != nil {
			return HandleError(c, err)
		}
	} else {
		err := h.Queries.AddProductCategoryByIDs(ctx,
			goqueries.AddProductCategoryByIDsParams{
				ProductID:  productID,
				CategoryID: categoryID,
			},
		)
		if err != nil {
			return HandleError(c, err)
		}
	}

	productCategories, err := h.Queries.GetCategoriesFilteredProductID(ctx, productID)
	if err != nil {
		return HandleError(c, err)
	}

	return Render(c, iomProduct.ProductCategory(
		productID,
		productCategories,
	),
	)

}
