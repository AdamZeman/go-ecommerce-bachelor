package pages

import (
	"database/sql"
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/app/tests/testUtils"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
	"testing"
)

func TestShowCategory(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show category",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProductsByCategoriesAndName", mock.Anything, goqueries.GetProductsByCategoriesAndNameParams{
					UserID:      2,
					CategoryIds: []int32{},
					NameFilter:  sql.NullString{String: "", Valid: true},
					SpecialIds:  []int32{},
					PriceFrom:   0,
					PriceTo:     999999,
					Offsetvar:   0,
					Limitvar:    30,
				}).Return(
					[]goqueries.GetProductsByCategoriesAndNameRow{{ID: 1}}, nil)
				m.On("GetCategories", mock.Anything).Return(
					[]goqueries.Category{{ID: 1}}, nil)
				m.On("GetSpecial", mock.Anything).Return(
					[]goqueries.Special{{ID: 1}}, nil)

			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/category",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowProducts(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 2},
	)
}
