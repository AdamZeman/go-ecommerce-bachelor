package actions

import (
	"ecomm-go/app/handlers"
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/app/tests/testUtils"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
	"testing"
)

func TestTotalPriceCalcBasket(t *testing.T) {
	tests := []struct {
		name     string
		input    []goqueries.GetBasketItemsByUserIdRow
		expected int32
	}{
		{
			name:     "empty basket",
			input:    []goqueries.GetBasketItemsByUserIdRow{},
			expected: 0,
		},
		{
			name: "single item",
			input: []goqueries.GetBasketItemsByUserIdRow{
				{Price: 1050, Quantity: 2},
			},
			expected: 2100, // 10.50 * 2
		},
		{
			name: "multiple items",
			input: []goqueries.GetBasketItemsByUserIdRow{
				{Price: 525, Quantity: 3},
				{Price: 1200, Quantity: 1},
				{Price: 750, Quantity: 4},
			},
			expected: 5775,
		},
		{
			name: "zero quantity items",
			input: []goqueries.GetBasketItemsByUserIdRow{
				{Price: 1500, Quantity: 0},
				{Price: 2000, Quantity: 1},
			},
			expected: 2000, // (15.00*0) + (20.00*1)
		},
		{
			name: "zero price items",
			input: []goqueries.GetBasketItemsByUserIdRow{
				{Price: 0, Quantity: 5},
				{Price: 800, Quantity: 2},
			},
			expected: 1600, // (0*5) + (8.00*2)
		},
		{
			name: "large quantities",
			input: []goqueries.GetBasketItemsByUserIdRow{
				{Price: 199, Quantity: 1000},
			},
			expected: 199000, // 1.99 * 1000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.TotalPriceCalcBasket(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAddToBasket(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name: "success add to basket",
			FormData: url.Values{
				"option1value": {"Red"}, "option2value": {"Medium"}, "amount": {"10"}, "product_id": {"10"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("GetVariantByOptions", mock.Anything, mock.Anything).
					Return(goqueries.ProductVariant{ID: 10, Price: 1000}, nil)
				m.On("InsertBasketItem", mock.Anything, mock.Anything).
					Return(goqueries.BasketItem{ID: 10}, nil)
				m.On("GetBasketItemsByUserId", mock.Anything, mock.Anything).
					Return([]goqueries.GetBasketItemsByUserIdRow{{ID: 1, UserID: 3, Quantity: 10, Price: 1000}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "fail insufficient params",
			FormData:     url.Values{"option1value": {"Red"}, "option2value": {"Medium"}, "amount": {"10"}},
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "fail wrong type params",
			FormData: url.Values{
				"option1value": {"Red"}, "option2value": {"Medium"}, "amount": {"Medium"}, "product_id": {"1"}},
			MockSetup:    func(m *queriesTest.Querier) {},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/api/add-to-basket",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).AddToBasket(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 3},
	)
}

func TestRemoveFromBasket(t *testing.T) {
	testCases := []testUtils.TestCase{
		{
			Name:     "success remove from basket",
			FormData: url.Values{"basketItemId": {"3"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("RemoveFromBasket", mock.Anything, int32(3)).
					Return(goqueries.BasketItem{ID: 3, UserID: 3, Quantity: 31}, nil)
				m.On("GetBasketItemsByUserId", mock.Anything, int32(3)).
					Return([]goqueries.GetBasketItemsByUserIdRow{{ID: 3, UserID: 3, Quantity: 3}}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "invalid basket item ID",
			FormData:     url.Values{"basketItemId": {"invalid"}},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:     "error removing from basket",
			FormData: url.Values{"basketItemId": {"3"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("RemoveFromBasket", mock.Anything, int32(3)).
					Return(goqueries.BasketItem{}, assert.AnError)
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:     "error getting basket items",
			FormData: url.Values{"basketItemId": {"3"}},
			MockSetup: func(m *queriesTest.Querier) {
				m.On("RemoveFromBasket", mock.Anything, int32(3)).
					Return(goqueries.BasketItem{ID: 3, UserID: 3, Quantity: 1}, nil)
				m.On("GetBasketItemsByUserId", mock.Anything, int32(3)).
					Return(nil, assert.AnError)
			},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	testUtils.RunTestCases(t,
		testCases,
		http.MethodPost,
		"/remove-from-basket",
		func(m *queriesTest.Querier, c echo.Context) error {
			return testUtils.CreateMockHandler(m).RemoveFromBasket(c)
		},
		types.CookieUser{Email: "test@example.com", ID: 3},
	)
}
