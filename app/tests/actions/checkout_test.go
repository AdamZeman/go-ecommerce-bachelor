package actions

import (
	"ecomm-go/app/handlers"
	"ecomm-go/db/goqueries"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTotalPriceCalc(t *testing.T) {
	tests := []struct {
		name     string
		input    []goqueries.GetBasketItemsByUserEmailRow
		expected int32
	}{
		{
			name:     "empty basket",
			input:    []goqueries.GetBasketItemsByUserEmailRow{},
			expected: 0,
		},
		{
			name: "single item",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 1050, Quantity: 2},
			},
			expected: 2100,
		},
		{
			name: "multiple items",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 525, Quantity: 3},
				{Price: 1200, Quantity: 1},
				{Price: 750, Quantity: 4},
			},
			expected: 5775,
		},
		{
			name: "zero quantity items",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 1500, Quantity: 0},
				{Price: 2000, Quantity: 1},
			},
			expected: 2000,
		},
		{
			name: "zero price items",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 0, Quantity: 5},
				{Price: 800, Quantity: 2},
			},
			expected: 1600,
		},
		{
			name: "large quantities",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 199, Quantity: 1000},
			},
			expected: 199000,
		},
		{
			name: "negative price (should handle but result may be negative)",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: -500, Quantity: 2},
			},
			expected: -1000,
		},
		{
			name: "negative quantity (should handle but result may be negative)",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: 300, Quantity: -3},
			},
			expected: -900,
		},
		{
			name: "maximum int32 values",
			input: []goqueries.GetBasketItemsByUserEmailRow{
				{Price: math.MaxInt32, Quantity: 1},
			},
			expected: math.MaxInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.TotalPriceCalc(tt.input)
			assert.Equal(t, tt.expected, result,
				"Test case '%s' failed: expected %d, got %d",
				tt.name, tt.expected, result)
		})
	}
}
