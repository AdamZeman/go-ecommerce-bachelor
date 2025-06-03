package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(app *echo.Echo, h DataHandler) {

	// Middlewares
	adminMiddleware := h.AuthMiddleware(2)
	customerMiddleware := h.AuthMiddleware(1)
	messengerMiddleware := h.MessengerMiddleware()
	CORSMiddleware := CORSConfigMiddleware()

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() == "/api/stripe/event" {
				return next(c)
			}
			return middleware.CSRFWithConfig(middleware.CSRFConfig{
				TokenLookup:    "cookie:csrf",
				CookieName:     "csrf",
				CookiePath:     "/",
				CookieHTTPOnly: true,
			})(next)(c)
		}
	})

	// ------ Public

	// Navigation
	app.GET("/contact", h.ShowContact)
	app.GET("/login", h.ShowLogin)
	app.GET("/category", h.ShowProducts)
	app.GET("/products/:id", h.ShowSingleProduct)
	app.GET("/", h.ShowHome)

	// Auth
	app.GET("/auth/:provider/callback", h.AuthCallback)
	app.GET("/auth/:provider", h.Auth)
	app.GET("/logout/:provider", h.Logout)

	// other api
	app.POST("/api/show-filtered-products", h.ShowFilteredProducts)
	app.POST("/api/show-more-products", h.ShowMoreProducts)
	app.GET("/api/update-product-info", h.UpdateProductInfo)

	// stripe
	app.POST("/api/stripe/checkout", h.CheckoutCreator, CORSMiddleware, customerMiddleware)
	app.POST("/api/stripe/event", h.HandleEvent)

	// UI

	// Static
	app.Static("/css", "./css")
	app.Static("/images", "./images")

	// ------ Customer

	// Navigation
	app.GET("/basket", h.ShowBasket, customerMiddleware)
	app.GET("/rooms", h.ShowRooms, customerMiddleware)
	app.GET("/orders", h.ShowOrders, customerMiddleware)
	app.GET("/api/messenger/:id", h.ShowMessengerAlone, messengerMiddleware, customerMiddleware)
	app.GET("/messenger/:id", h.ShowMessengerPage, messengerMiddleware, customerMiddleware)
	app.GET("/favourites", h.ShowFavourites, customerMiddleware)
	app.GET("/shipping", h.ShowShipping, customerMiddleware)

	// WebSocket
	app.GET("/ws", h.HandleWebSocket, customerMiddleware)

	// Basket
	app.POST("/api/add-to-basket", h.AddToBasket, customerMiddleware)
	app.POST("/api/remove-from-basket", h.RemoveFromBasket, customerMiddleware)
	app.GET("/api/update-basket-size", h.UpdateBasketSize, customerMiddleware)
	app.GET("/api/get-quick-basket", h.ShowQuickBasket, customerMiddleware)

	// Favourites
	app.POST("/api/add-to-favourites", h.AddToFavourites, customerMiddleware)
	app.POST("/api/remove-from-favourites", h.RemoveFromFavourites, customerMiddleware)

	// Rooms
	app.POST("/api/open-issue", h.OpenIssue, customerMiddleware)

	// Review
	app.POST("/api/orders/show-add-review", h.ShowAddReview, customerMiddleware)
	app.POST("/api/orders/add-review", h.AddReview, customerMiddleware)

	// Shipping
	app.POST("/api/shipping/validate-phone", h.ValidatePhone, customerMiddleware)

	// ------ Admin

	// IOM
	app.GET("/iom", h.ShowIOM, adminMiddleware)

	// product
	app.GET("/api/show-product-management", h.ShowProductManagement, adminMiddleware)
	app.GET("/api/show-add-product", h.ShowAddProduct, adminMiddleware)
	app.GET("/api/show-edit-product", h.ShowEditProduct, adminMiddleware)

	app.POST("/api/iom/add-product", h.AddProduct, adminMiddleware)
	app.POST("/api/iom/edit-product", h.EditProduct, adminMiddleware)
	app.POST("/api/iom/delete-product", h.DeleteProduct, adminMiddleware)

	app.GET("/api/show-product-categories", h.ShowProductCategories, adminMiddleware)
	app.POST("/api/product-category-switch", h.CategorySwitch, adminMiddleware)

	// variant
	app.GET("/api/show-variant-management", h.ShowVariantManagement, adminMiddleware)
	app.GET("/api/show-edit-variants", h.ShowEditVariant, adminMiddleware)
	app.GET("/api/show-add-variant", h.ShowAddVariant, adminMiddleware)

	app.POST("/api/iom/edit-variant", h.EditVariant, adminMiddleware)
	app.POST("/api/iom/add-variant", h.AddVariant, adminMiddleware)
	app.POST("/api/iom/delete-variant", h.DeleteVariant, adminMiddleware)

	// category
	app.GET("/api/show-category-management", h.ShowCategoryManagement, adminMiddleware)
	app.GET("/api/show-edit-category", h.ShowEditCategory, adminMiddleware)
	app.GET("/api/show-add-category", h.ShowAddCategory, adminMiddleware)

	app.POST("/api/iom/edit-category", h.EditCategory, adminMiddleware)
	app.POST("/api/iom/add-category", h.AddCategory, adminMiddleware)
	app.POST("/api/iom/delete-category", h.DeleteCategory, adminMiddleware)

	// special
	app.GET("/api/show-special-management", h.ShowSpecialManagement, adminMiddleware)
	app.GET("/api/show-edit-special", h.ShowEditSpecial, adminMiddleware)
	app.GET("/api/show-add-special", h.ShowAddSpecial, adminMiddleware)

	app.POST("/api/iom/edit-special", h.EditSpecial, adminMiddleware)
	app.POST("/api/iom/add-special", h.AddSpecial, adminMiddleware)
	app.POST("/api/iom/delete-special", h.DeleteSpecial, adminMiddleware)

	//order
	app.GET("/api/show-order-management", h.ShowOrderManagement, adminMiddleware)
	app.GET("/api/show-edit-order", h.ShowEditOrder, adminMiddleware)
	app.POST("/api/iom/edit-order", h.EditOrder, adminMiddleware)

	//pending order
	app.GET("/api/show-pending-order-management", h.ShowPendingOrderManagement, adminMiddleware)
	app.POST("/api/iom/edit-pending-order", h.EditPendingOrder, adminMiddleware)

	//issue order
	app.GET("/api/show-issue-order-management", h.ShowIssueOrderManagement, adminMiddleware)
	app.POST("/api/iom/edit-issue-order", h.EditIssueOrder, adminMiddleware)

	// shared
	app.GET("/api/remove-popup", h.RemovePopup, adminMiddleware)
}
