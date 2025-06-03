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

func TestShowHome(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show home",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProductsByCategoriesAndName", mock.Anything, goqueries.GetProductsByCategoriesAndNameParams{
					CategoryIds: []int32{},
					NameFilter:  sql.NullString{String: "", Valid: true},
					SpecialIds:  []int32{2},
					PriceFrom:   0,
					PriceTo:     999999,
					Offsetvar:   0,
					Limitvar:    10,
				}).Return(
					[]goqueries.GetProductsByCategoriesAndNameRow{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowHome(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 2},
	)
}
