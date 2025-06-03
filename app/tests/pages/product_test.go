package pages

import (
	"database/sql"
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/app/tests/testUtils"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
	"testing"
)

func TestShowProduct(t *testing.T) {
	user := types.CookieUser{Email: "test@example.com", ID: 2}
	productID := int32(5)

	testCases := []testUtils.TestCase{
		{
			Name:     "success show product",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("ProductVariantsByProductId", mock.Anything, productID).Return(
					[]goqueries.ProductVariant{{ID: productID, Option1Value: "Red", Option2Value: "Small"}}, nil)
				m.On("GetProductById", mock.Anything, int32(5)).Return(
					goqueries.Product{ID: productID}, nil)
				m.On("GetFavouriteByProductId", mock.Anything, goqueries.GetFavouriteByProductIdParams{
					ProductID: productID,
					UserID:    user.ID,
				}).Return(
					[]goqueries.FavouriteItem{{ID: 1}}, nil)
				m.On("GetReviewsFillUserByProductID", mock.Anything, productID).Return(
					[]goqueries.GetReviewsFillUserByProductIDRow{{Rating: 5}}, nil)
				m.On("GetProductsByCategoriesAndName", mock.Anything, goqueries.GetProductsByCategoriesAndNameParams{
					UserID:      user.ID,
					CategoryIds: []int32{},
					NameFilter:  sql.NullString{String: "", Valid: true},
					SpecialIds:  []int32{1},
					PriceFrom:   0,
					PriceTo:     999999,
					Offsetvar:   0,
					Limitvar:    10,
				}).Return(
					[]goqueries.GetProductsByCategoriesAndNameRow{{ID: 1}}, nil)
				m.On("GetCategoriesByProductID", mock.Anything, productID).Return(
					[]goqueries.Category{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		fmt.Sprint("/product/", productID),
		func(m *queriesTest.Querier, c echo.Context) error {
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprint(productID))
			return testUtils.CreateMockHandler(m).ShowSingleProduct(c)
		},
		user,
	)
}
