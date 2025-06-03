package actions

import (
	"ecomm-go/app/handlers"
	"ecomm-go/db/goqueries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterByOrderId(t *testing.T) {
	tests := []struct {
		name     string
		input    []goqueries.GetOrderItemsByUserIdRow
		filterID int32
		expected []goqueries.GetOrderItemsByUserIdRow
	}{
		{
			name:     "empty list",
			input:    []goqueries.GetOrderItemsByUserIdRow{},
			filterID: 1,
			expected: []goqueries.GetOrderItemsByUserIdRow{},
		},
		{
			name: "no matching items",
			input: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 2},
				{OrderID: 3},
				{OrderID: 4},
			},
			filterID: 1,
			expected: []goqueries.GetOrderItemsByUserIdRow{},
		},
		{
			name: "some matching items",
			input: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 1, ProductID: 100},
				{OrderID: 2, ProductID: 101},
				{OrderID: 1, ProductID: 102},
				{OrderID: 3, ProductID: 103},
			},
			filterID: 1,
			expected: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 1, ProductID: 100},
				{OrderID: 1, ProductID: 102},
			},
		},
		{
			name: "all items match",
			input: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 5, ProductID: 200},
				{OrderID: 5, ProductID: 201},
				{OrderID: 5, ProductID: 202},
			},
			filterID: 5,
			expected: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 5, ProductID: 200},
				{OrderID: 5, ProductID: 201},
				{OrderID: 5, ProductID: 202},
			},
		},
		{
			name: "multiple fields in struct",
			input: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 1, ProductID: 100, Quantity: 2, Price: 1099},
				{OrderID: 2, ProductID: 101, Quantity: 1, Price: 599},
				{OrderID: 1, ProductID: 102, Quantity: 3, Price: 750},
			},
			filterID: 1,
			expected: []goqueries.GetOrderItemsByUserIdRow{
				{OrderID: 1, ProductID: 100, Quantity: 2, Price: 1099},
				{OrderID: 1, ProductID: 102, Quantity: 3, Price: 750},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.FilterByOrderId(tt.input, tt.filterID)
			assert.Equal(t, tt.expected, result)
		})
	}
}
