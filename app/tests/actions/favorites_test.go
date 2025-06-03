package actions

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

func TestAddToFavourites(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success remove existing",
			FormData: url.Values{"product_id": {"10"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetFavouriteByProductId", mock.Anything, goqueries.GetFavouriteByProductIdParams{
					ProductID: int32(10),
					UserID:    int32(10),
				}).
					Return([]goqueries.FavouriteItem{{ID: 1, UserID: 10, ProductID: 10}}, nil)
				m.On("RemoveItemFromFavourites", mock.Anything, goqueries.RemoveItemFromFavouritesParams{
					ProductID: int32(10),
					UserID:    int32(10),
				}).
					Return(nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:     "success insert new",
			FormData: url.Values{"product_id": {"10"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetFavouriteByProductId", mock.Anything, goqueries.GetFavouriteByProductIdParams{
					ProductID: int32(10),
					UserID:    int32(10),
				}).
					Return([]goqueries.FavouriteItem{}, nil)
				m.On("InsertItemToFavourites", mock.Anything, goqueries.InsertItemToFavouritesParams{
					UserID:    10,
					ProductID: 10,
				}).Return(goqueries.FavouriteItem{ID: 1, UserID: 10, ProductID: 10}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "fail empty form",
			FormData:     nil,
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "fail invalid product_id",
			FormData:     url.Values{"product_id": {"hello"}},
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/add-to-favourites",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).AddToFavourites(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 10},
	)
}

func TestRemoveFromFavourites(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success remove favourite",
			FormData: url.Values{"favouritesItemProductId": {"1"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("RemoveItemFromFavourites", mock.Anything, goqueries.RemoveItemFromFavouritesParams{
					ProductID: int32(1),
					UserID:    int32(10),
				}).Return(nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}
	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/remove-from-favourites",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).RemoveFromFavourites(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 10},
	)
}
