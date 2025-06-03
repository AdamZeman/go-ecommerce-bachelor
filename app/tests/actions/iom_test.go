package actions

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

func TestIOMAddProduct(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name: "success add product",
			FormData: url.Values{
				"product-name": {"T-Shirt"}, "product-desc": {"desc"}, "product-price": {"100"}, "product-option1-name": {"Color"}, "product-option2-name": {"Size"}, "product-special-id": {"2"},
			},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("InsertProduct", mock.Anything, goqueries.InsertProductParams{
					Name:        "T-Shirt",
					Description: "desc",
					Price:       100,
					Option1Name: "Color",
					Option2Name: "Size",
					Special:     sql.NullInt32{Valid: true, Int32: 2},
				}).
					Return(nil)

				m.On("GetProducts", mock.Anything).
					Return([]goqueries.Product{}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "fail Wrong Form Values",
			FormData:     url.Values{"option1value": {"Red"}, "option2value": {"Medium"}, "amount": {"10"}},
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/api/iom/add-product",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).AddProduct(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 3},
	)
}

func TestIOMEditProduct(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name: "success edit product",
			FormData: url.Values{
				"product-id": {"1"}, "product-name": {"T-Shirt"}, "product-desc": {"desc"}, "product-price": {"100"}, "product-option1-name": {"Color"}, "product-option2-name": {"Size"}, "product-special-id": {"2"},
			},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("UpdateProduct", mock.Anything, goqueries.UpdateProductParams{
					ID:          int32(1),
					Name:        "T-Shirt",
					Description: "desc",
					Price:       100,
					Option1Name: "Color",
					Option2Name: "Size",
					Special:     sql.NullInt32{Valid: true, Int32: 2},
				}).
					Return(nil)

				m.On("GetProducts", mock.Anything).
					Return([]goqueries.Product{}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "fail Wrong Form Values",
			FormData:     url.Values{"option1value": {"Red"}, "option2value": {"Medium"}, "amount": {"10"}},
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/api/iom/edit-product",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).EditProduct(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 3},
	)
}
