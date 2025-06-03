package actions

import (
	"database/sql"
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/app/tests/testUtils"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
	"testing"
)

func TestFilterProducts(t *testing.T) {
	user := types.CookieUser{Email: "test@example.com", ID: 3}

	testCases := []testUtils.TestCase{
		{
			Name: "success filter products",
			FormData: url.Values{
				"search":     {"Red T-shirt"},
				"price-from": {"10"},
				"price-to":   {"100"},
				"category":   {"1", "2"},
				"special":    {"5"},
				"page":       {"0"},
			},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProductsByCategoriesAndName", mock.Anything, goqueries.GetProductsByCategoriesAndNameParams{
					UserID:      user.ID,
					CategoryIds: []int32{1, 2},
					NameFilter:  sql.NullString{String: "Red T-shirt", Valid: true},
					SpecialIds:  []int32{5},
					PriceFrom:   int32(10),
					PriceTo:     int32(100),
					Offsetvar:   0,
					Limitvar:    30,
				}).
					Return([]goqueries.GetProductsByCategoriesAndNameRow{
						{Name: "Red T-shirt1", Price: int32(11), Special: sql.NullInt32{Int32: 5, Valid: true}},
						{Name: "Red T-shirt2", Price: int32(15), Special: sql.NullInt32{Int32: 5, Valid: true}},
					}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:     "success filter products empty",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProductsByCategoriesAndName", mock.Anything, goqueries.GetProductsByCategoriesAndNameParams{
					UserID:      user.ID,
					CategoryIds: nil,
					NameFilter:  sql.NullString{String: "", Valid: true},
					SpecialIds:  nil,
					PriceFrom:   int32(0),
					PriceTo:     int32(999999),
					Offsetvar:   0,
					Limitvar:    30,
				}).
					Return([]goqueries.GetProductsByCategoriesAndNameRow{
						{Name: "Red T-shirt1", Price: int32(11), Special: sql.NullInt32{Int32: 5, Valid: true}},
						{Name: "Red T-shirt2", Price: int32(15), Special: sql.NullInt32{Int32: 5, Valid: true}},
					}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name: "error filter products ",
			FormData: url.Values{
				"search":     {"None"},
				"price-from": {"10"},
				"price-to":   {"100"},
				"category":   {"1", "2"},
				"special":    {"6"},
				"page":       {"0"},
			},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProductsByCategoriesAndName",
					mock.Anything,
					mock.AnythingOfType("goqueries.GetProductsByCategoriesAndNameParams"),
				).Return(nil, errors.New("DB exploded"))
			},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/api/show-filtered-products",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowFilteredProducts(c)
		},
		user,
	)
}
