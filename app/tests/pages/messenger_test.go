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

func TestShowMessenger(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show messenger",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetMessagesByConversationId", mock.Anything, int32(1)).Return(
					[]goqueries.GetMessagesByConversationIdRow{{Email: "hello"}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}
	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/messenger/1",
		func(m *queriesTest.Querier, c echo.Context) error {
			c.SetParamNames("id")
			c.SetParamValues("1")
			return testUtils.CreateMockHandler(m).ShowMessengerAlone(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 1},
	)
}
