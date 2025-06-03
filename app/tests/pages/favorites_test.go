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

func TestShowFavourites(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show favourites",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetFavouritesById", mock.Anything, "test@example.com").
					Return([]goqueries.GetFavouritesByIdRow{{ID: 1, UserID: 3}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/favourites",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowFavourites(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 10},
	)
}
