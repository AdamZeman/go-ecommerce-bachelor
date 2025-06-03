package types

import "ecomm-go/db/goqueries"

type Product struct {
	Name string
	Id   int32
}

type Category struct {
	Name string
}

type CookieUser struct {
	Provider    string
	Email       string
	Name        string
	UserID      string
	AvatarURL   string
	AccessToken string
	Role        int
	ID          int32
	IsSigned    bool
}

type OrderFilled struct {
	Order      goqueries.Order
	OrderItems []goqueries.GetOrderItemsByUserIdRow
}

type CalculateReviewData struct {
	One     int
	Two     int
	Three   int
	Four    int
	Five    int
	Average float64
	All     int
}
