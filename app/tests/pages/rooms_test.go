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

func TestShowRooms(t *testing.T) {
	user := types.CookieUser{Email: "test@example.com", ID: 2}
	firstMessage := goqueries.Conversation{ID: 8}

	testCases := []testUtils.TestCase{
		{
			Name:     "success show rooms",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetConversationsByUserId", mock.Anything, user.Email).Return(
					[]goqueries.Conversation{firstMessage}, nil)
				m.On("GetMessagesByConversationId", mock.Anything, firstMessage.ID).Return(
					[]goqueries.GetMessagesByConversationIdRow{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/rooms",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowRooms(c)
		},
		user,
	)
}
