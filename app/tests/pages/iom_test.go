package pages

import (
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

func TestShowIOM(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show iom",
			FormData: url.Values{},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetProducts", mock.Anything).Return(
					[]goqueries.Product{{ID: 1}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodGet,
		"/basket",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowIOM(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 2},
	)
}

func TestShowEditCategory(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success show edit form",
			FormData: url.Values{"id": {"1"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetCategoryByID", mock.Anything, int32(1)).
					Return(goqueries.Category{ID: 1, Name: "Electronics"}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "missing id parameter",
			FormData:     url.Values{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "invalid id format",
			FormData:     url.Values{"id": {"invalid"}},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:     "category not found",
			FormData: url.Values{"id": {"99"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetCategoryByID", mock.Anything, int32(99)).
					Return(goqueries.Category{}, errors.New("not found"))
			},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/api/show-edit-category",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).ShowEditCategory(c)
		},
		types.CookieUser{Email: "admin@example.com", ID: 1},
	)
}
