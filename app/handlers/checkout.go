package handlers

import (
	"ecomm-go/db/goqueries"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
	"io/ioutil"
	"log"
	"net/http"
)

type SessionOutput struct {
	Id              string `json:"id"`
	StripePublicKey string `json:"stripePublicKey"`
}

func TotalPriceCalc(basketItems []goqueries.GetBasketItemsByUserEmailRow) int32 {
	count := int32(0)
	for _, basketItem := range basketItems {
		count += basketItem.Price * basketItem.Quantity
	}
	return count
}

func Checkout(email string, c echo.Context, h DataHandler) (*stripe.CheckoutSession, error) {
	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("FinalEmail", email)
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, HandleError(c, err)
	}

	meta := map[string]string{
		"FinalEmail": email,
	}

	log.Println("Creating meta for user: ", meta)

	var stripeItems []*stripe.CheckoutSessionLineItemParams
	ctx := c.Request().Context()
	itemProductVariants, err := h.Queries.GetBasketItemsByUserEmail(ctx, email)
	if err != nil {
		return nil, HandleError(c, err)
	}

	for _, basketItem := range itemProductVariants {
		product := &stripe.CheckoutSessionLineItemParams{
			Name:        stripe.String(basketItem.Name),
			Description: stripe.String("Description of the product"),
			Amount:      stripe.Int64(int64(basketItem.Price)),
			Currency:    stripe.String(string(stripe.CurrencyUSD)),
			Quantity:    stripe.Int64(int64(basketItem.Quantity)),
		}
		stripeItems = append(stripeItems, product)
	}

	params := &stripe.CheckoutSessionParams{
		Customer:   &newCustomer.ID,
		SuccessURL: stripe.String(h.DefaultURL + "/orders"),
		CancelURL:  stripe.String(h.DefaultURL + "/orders"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode:      stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: stripeItems,
	}
	return session.New(params)
}

func GetEvent(c echo.Context) (*stripe.Event, error) {
	const MaxBodyBytes = int64(65536)
	body := http.MaxBytesReader(c.Response(), c.Request().Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, HandleError(c, err)
	}
	event := stripe.Event{}
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, HandleError(c, err)
	}
	return &event, nil
}

type CheckoutCreatorForm struct {
	PhoneNumber string `form:"phone-number"`
	Address     string `form:"address"`
	HouseNumber string `form:"house-number"`
	PostalCode  string `form:"postal-code"`
	City        string `form:"city"`
	Notes       string `form:"notes"`
}

func (h DataHandler) CheckoutCreator(c echo.Context) error {
	user, _ := GetUserFromSession(h, c)
	ctx := c.Request().Context()

	var form CheckoutCreatorForm
	if err := c.Bind(&form); err != nil {
		return HandleError(c, err)
	}
	if err := c.Validate(&form); err != nil {
		return HandleError(c, err, "Wrong Form Values")
	}
	_, err := h.Queries.InsertShipping(ctx, goqueries.InsertShippingParams{
		UserID:      user.ID,
		PhoneNumber: form.PhoneNumber,
		Address:     form.Address,
		HouseNumber: form.HouseNumber,
		PostalCode:  form.PostalCode,
		City:        form.City,
		Notes:       form.Notes,
	})
	if err != nil {
		return HandleError(c, err)
	}

	stripeSession, err := Checkout(user.Email, c, h)
	if err != nil {
		return HandleError(c, err)
	}
	return c.JSON(http.StatusOK, &SessionOutput{
		Id:              stripeSession.ID,
		StripePublicKey: h.StripePublicKey,
	})
}

func (h DataHandler) HandleEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := GetEvent(c)
	if err != nil {
		return HandleError(c, err)
	}
	if event.Type == "charge.succeeded" {
		customerID := event.Data.Object["customer"].(string)
		customerForMeta, err := customer.Get(customerID, nil)
		if err != nil {
			return HandleError(c, err)
		}
		email := customerForMeta.Metadata["FinalEmail"]

		itemProductVariants, err := h.Queries.GetBasketItemsByUserEmail(ctx, email)
		if err != nil {
			return HandleError(c, err)
		}

		latestShipping, err := h.Queries.GetLatestShippingByUserId(ctx, email)
		if err != nil {
			return HandleError(c, err)
		}

		newDatabaseOrder, err := h.Queries.InsertOrder(ctx, goqueries.InsertOrderParams{
			UserID:     itemProductVariants[0].UserID,
			TotalPrice: TotalPriceCalc(itemProductVariants),
			ShippingID: latestShipping.ID,
		})
		if err != nil {
			return HandleError(c, err)
		}

		for _, itemProductVariant := range itemProductVariants {
			_, err := h.Queries.InsertOrderItem(ctx, goqueries.InsertOrderItemParams{
				OrderID:          newDatabaseOrder.ID,
				ProductVariantID: itemProductVariant.ProductVariantID,
				Quantity:         itemProductVariant.Quantity,
			})
			if err != nil {
				return HandleError(c, err)
			}
		}
		for _, itemProductVariant := range itemProductVariants {
			err = h.Queries.DeleteItemsFromBasket(ctx, itemProductVariant.ID)
			if err != nil {
				return HandleError(c, err)
			}
		}
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "Event processed"})
}
