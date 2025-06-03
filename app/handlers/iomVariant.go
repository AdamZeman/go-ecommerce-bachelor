package handlers

import (
	"ecomm-go/app/views/components/iomVariant"
	"ecomm-go/db/goqueries"

	"github.com/labstack/echo/v4"
)

func (h DataHandler) ShowVariantManagement(c echo.Context) error {
	ctx := c.Request().Context()
	variants, err := h.Queries.GetVariantsFilledProducts(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomVariant.VariantManagement(variants))
}

func (h DataHandler) ShowEditVariant(c echo.Context) error {
	ctx := c.Request().Context()

	variantID, _ := ParseToInt32(c.FormValue("id"))
	variant, err := h.Queries.GetVariantById(ctx,
		variantID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return Render(c, iomVariant.EditVariant(variant))
}

type EditVariantForm struct {
	ID            int32  `form:"variant-id"      validate:"required"`
	Sku           string `form:"variant-sku"      validate:"required,min=2,max=100"`
	StockQuantity int32  `form:"variant-stock-quantity"      validate:"required,gt=0"`
	Price         int32  `form:"variant-price"     validate:"required,gt=0"`
	Option1Value  string `form:"variant-option1-value" validate:"omitempty,max=50"`
	Option2Value  string `form:"variant-option2-value" validate:"omitempty,max=50"`
}

func (h DataHandler) EditVariant(c echo.Context) error {
	ctx := c.Request().Context()

	var form EditVariantForm
	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}

	err := h.Queries.UpdateVariant(ctx,
		goqueries.UpdateVariantParams{
			ID:            form.ID,
			Sku:           form.Sku,
			StockQuantity: form.StockQuantity,
			Price:         form.Price,
			Option1Value:  form.Option1Value,
			Option2Value:  form.Option2Value,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowVariantManagement(c)
}

func (h DataHandler) ShowAddVariant(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := h.Queries.GetProducts(ctx)
	if err != nil {
		return HandleError(c, err)
	}

	return Render(c, iomVariant.AddVariant(products))
}

type AddVariantForm struct {
	ProductID     int32  `form:"product-id"      validate:"required"`
	Sku           string `form:"variant-sku"      validate:"required,min=2,max=100"`
	StockQuantity int32  `form:"variant-stock-quantity"      validate:"required,gt=0"`
	Price         int32  `form:"variant-price"     validate:"required,gt=0"`
	Option1Value  string `form:"variant-option1-value" validate:"omitempty,max=50"`
	Option2Value  string `form:"variant-option2-value" validate:"omitempty,max=50"`
}

func (h DataHandler) AddVariant(c echo.Context) error {
	ctx := c.Request().Context()
	var form AddVariantForm
	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}

	err := h.Queries.InsertVariant(ctx,
		goqueries.InsertVariantParams{
			ProductID:     form.ProductID,
			Sku:           form.Sku,
			Price:         form.Price,
			Option1Value:  form.Option1Value,
			StockQuantity: form.StockQuantity,
			Option2Value:  form.Option2Value,
		},
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowVariantManagement(c)
}

func (h DataHandler) DeleteVariant(c echo.Context) error {
	ctx := c.Request().Context()
	productID, _ := ParseToInt32(c.FormValue("variant-id"))

	err := h.Queries.DeleteVariantByID(ctx,
		productID,
	)
	if err != nil {
		return HandleError(c, err)
	}
	return h.ShowVariantManagement(c)
}
