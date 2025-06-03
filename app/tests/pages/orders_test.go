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

func TestShowOrders(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show orders",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetOrderItemsByUserId", mock.Anything, int32(2)).Return(
					[]goqueries.GetOrderItemsByUserIdRow{{ID: 1}}, nil)
				m.On("GetOrdersByUserId", mock.Anything, int32(2)).Return(
					[]goqueries.Order{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/orders",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowOrders(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 2},
	)
}
