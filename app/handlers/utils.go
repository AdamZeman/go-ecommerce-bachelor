package handlers

import (
	"context"
	"database/sql"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"fmt"
	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type Querier interface {
	WithTx(tx *sql.Tx) *goqueries.Queries
	AddProductCategoryByIDs(ctx context.Context, arg goqueries.AddProductCategoryByIDsParams) error
	CountConvByUserConv(ctx context.Context, arg goqueries.CountConvByUserConvParams) (int64, error)
	DeleteCategoryByID(ctx context.Context, id int32) error
	DeleteItemsFromBasket(ctx context.Context, id int32) error
	DeleteProductByID(ctx context.Context, id int32) error
	DeleteProductCategoryByIDs(ctx context.Context, arg goqueries.DeleteProductCategoryByIDsParams) error
	DeleteSpecialByID(ctx context.Context, id int32) error
	DeleteVariantByID(ctx context.Context, id int32) error
	GetBasketItemsByUserEmail(ctx context.Context, email string) ([]goqueries.GetBasketItemsByUserEmailRow, error)
	GetBasketItemsByUserId(ctx context.Context, userID int32) ([]goqueries.GetBasketItemsByUserIdRow, error)
	GetCategories(ctx context.Context) ([]goqueries.Category, error)
	GetCategoriesByProductID(ctx context.Context, productID int32) ([]goqueries.Category, error)
	GetCategoriesFilteredProductID(ctx context.Context, productID int32) ([]goqueries.GetCategoriesFilteredProductIDRow, error)
	GetCategoryByID(ctx context.Context, id int32) (goqueries.Category, error)
	GetConversations(ctx context.Context) ([]goqueries.Conversation, error)
	GetConversationsByUserId(ctx context.Context, email string) ([]goqueries.Conversation, error)
	GetFavouriteByProductId(ctx context.Context, arg goqueries.GetFavouriteByProductIdParams) ([]goqueries.FavouriteItem, error)
	GetFavouritesById(ctx context.Context, email string) ([]goqueries.GetFavouritesByIdRow, error)
	GetMessages(ctx context.Context) ([]goqueries.GetMessagesRow, error)
	GetMessagesByConversationId(ctx context.Context, id int32) ([]goqueries.GetMessagesByConversationIdRow, error)
	GetLatestShippingByUserId(ctx context.Context, email string) (goqueries.Shipping, error)
	GetOpenConversationById(ctx context.Context, arg goqueries.GetOpenConversationByIdParams) ([]goqueries.GetOpenConversationByIdRow, error)
	GetOrderByID(ctx context.Context, id int32) (goqueries.Order, error)
	GetOrderItemsByOrderID(ctx context.Context, id int32) ([]goqueries.GetOrderItemsByOrderIDRow, error)
	GetOrderItemsByUserId(ctx context.Context, userID int32) ([]goqueries.GetOrderItemsByUserIdRow, error)
	GetOrdersByUserId(ctx context.Context, userID int32) ([]goqueries.Order, error)
	GetOrdersFillUser(ctx context.Context) ([]goqueries.GetOrdersFillUserRow, error)
	GetOrdersFillUserByStatus(ctx context.Context, status string) ([]goqueries.GetOrdersFillUserByStatusRow, error)
	GetOrdersFillUserConvByStatus(ctx context.Context, status string) ([]goqueries.GetOrdersFillUserConvByStatusRow, error)
	GetProductById(ctx context.Context, id int32) (goqueries.Product, error)
	GetProductVariantsByProductId(ctx context.Context, id int32) ([]goqueries.ProductVariant, error)
	GetProducts(ctx context.Context) ([]goqueries.Product, error)
	GetProductsByCategoriesAndName(ctx context.Context, arg goqueries.GetProductsByCategoriesAndNameParams) ([]goqueries.GetProductsByCategoriesAndNameRow, error)
	GetProductsByCategoryId(ctx context.Context, categoryID int32) ([]goqueries.Product, error)
	GetReviewsFillUserByProductID(ctx context.Context, productID int32) ([]goqueries.GetReviewsFillUserByProductIDRow, error)
	GetSpecial(ctx context.Context) ([]goqueries.Special, error)
	GetSpecialByID(ctx context.Context, id int32) (goqueries.Special, error)
	GetUsers(ctx context.Context) ([]goqueries.User, error)
	GetVariantById(ctx context.Context, id int32) (goqueries.ProductVariant, error)
	GetVariantByOptions(ctx context.Context, arg goqueries.GetVariantByOptionsParams) (goqueries.ProductVariant, error)
	GetVariantsFilledProducts(ctx context.Context) ([]goqueries.GetVariantsFilledProductsRow, error)
	InsertBasketItem(ctx context.Context, arg goqueries.InsertBasketItemParams) (goqueries.BasketItem, error)
	InsertCategory(ctx context.Context, name string) error
	InsertConversation(ctx context.Context, arg goqueries.InsertConversationParams) (goqueries.Conversation, error)
	InsertConversationUser(ctx context.Context, arg goqueries.InsertConversationUserParams) (goqueries.ConversationsUser, error)
	InsertItemToFavourites(ctx context.Context, arg goqueries.InsertItemToFavouritesParams) (goqueries.FavouriteItem, error)
	InsertMessage(ctx context.Context, arg goqueries.InsertMessageParams) (goqueries.ConversationMessage, error)
	InsertOrder(ctx context.Context, arg goqueries.InsertOrderParams) (goqueries.Order, error)
	InsertOrderItem(ctx context.Context, arg goqueries.InsertOrderItemParams) (goqueries.OrderItem, error)
	InsertProduct(ctx context.Context, arg goqueries.InsertProductParams) error
	InsertReview(ctx context.Context, arg goqueries.InsertReviewParams) (goqueries.Review, error)
	InsertShipping(ctx context.Context, arg goqueries.InsertShippingParams) (goqueries.Shipping, error)
	InsertSpecial(ctx context.Context, name string) error
	InsertVariant(ctx context.Context, arg goqueries.InsertVariantParams) error
	ProductVariantsByProductId(ctx context.Context, productID int32) ([]goqueries.ProductVariant, error)
	RemoveFromBasket(ctx context.Context, id int32) (goqueries.BasketItem, error)
	RemoveItemFromFavourites(ctx context.Context, arg goqueries.RemoveItemFromFavouritesParams) error
	UpdateCategory(ctx context.Context, arg goqueries.UpdateCategoryParams) error
	UpdateOrderStatus(ctx context.Context, arg goqueries.UpdateOrderStatusParams) (goqueries.Order, error)
	UpdateProduct(ctx context.Context, arg goqueries.UpdateProductParams) error
	UpdateSpecial(ctx context.Context, arg goqueries.UpdateSpecialParams) error
	UpdateVariant(ctx context.Context, arg goqueries.UpdateVariantParams) error
	UpsertUser(ctx context.Context, arg goqueries.UpsertUserParams) (goqueries.User, error)
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func ParseToBool(str string) (bool, error) {
	if str == "" {
		return false, fmt.Errorf("empty string provided")
	}
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false, fmt.Errorf("failed to parse '%s' as bool: %w", str, err)
	}
	return val, nil
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

type DataHandler struct {
	Store           *sessions.CookieStore
	Queries         Querier
	DefaultURL      string
	StripePublicKey string
}

var GetUserFromSession = func(h DataHandler, c echo.Context) (types.CookieUser, error) {
	session, err := h.Store.Get(c.Request(), "session-name")
	if err != nil {
		return types.CookieUser{
				IsSigned: false,
			},
			fmt.Errorf("failed to get session")
	}
	user, ok := session.Values["user"].(types.CookieUser)
	if !ok {
		return types.CookieUser{
				IsSigned: false,
			},
			nil
	}
	return user, nil
}

func ParseToInt32(strID string) (int32, error) {
	if strID == "" {
		return 0, fmt.Errorf("empty string provided")
	}
	id64, err := strconv.ParseInt(strID, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s' as int32: %w", strID, err)
	}
	return int32(id64), nil
}
func HandleError(c echo.Context, err error, message ...string) error {
	msg := "HandleError, Bad Request"
	if len(message) > 0 {
		msg = message[0]
	}
	log.Printf("HandleError: %v, Message: %s", err, msg)
	return c.HTML(http.StatusBadRequest, fmt.Sprintf("<h1>%s</h1>", msg))
}
