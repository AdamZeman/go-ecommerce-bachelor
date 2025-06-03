package pages

import (
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

func TestShowBasket(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show basket",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetBasketItemsByUserId", mock.Anything, int32(2)).Return(
					[]goqueries.GetBasketItemsByUserIdRow{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/basket",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowBasket(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 2},
	)
}
