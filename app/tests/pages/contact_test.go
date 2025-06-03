package pages

import (
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/app/tests/testUtils"
	"ecomm-go/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"testing"
)

func TestShowContact(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show contact",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
			},
			ExpectedCode: http.StatusOK,
		},
	}
	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/contact",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowContact(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 10},
	)
}
