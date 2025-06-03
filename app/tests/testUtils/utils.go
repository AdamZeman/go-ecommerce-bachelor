package testUtils

import (
	"ecomm-go/app/handlers"
	"ecomm-go/app/tests/queriesTest"
	"ecomm-go/types"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type TestCase struct {
	Name         string
	FormData     url.Values
	MockSetup    func(*queriesTest.Querier)
	ExpectError  bool
	ExpectedCode int
}
type HandlerFunc func(*queriesTest.Querier, echo.Context) error

func RunTestCases(t *testing.T, testCases []TestCase, method string, path string, handlerFunc HandlerFunc, user types.CookieUser) {
	log.SetOutput(io.Discard)
	handlers.GetUserFromSession = func(h handlers.DataHandler, c echo.Context) (types.CookieUser, error) {
		return user, nil
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			c := createFormContext(method, path, tc.FormData)
			mockQueries := new(queriesTest.Querier)
			if tc.MockSetup != nil {
				tc.MockSetup(mockQueries)
			}
			err := handlerFunc(mockQueries, c)
			if tc.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedCode, c.Response().Status)
			}
			mockQueries.AssertExpectations(t)
		})
	}
}
func createFormContext(method, path string, formData url.Values) echo.Context {
	e := echo.New()
	e.Validator = &handlers.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("csrf", "test-csrf-token")
	return c
}
func CreateMockHandler(queries *queriesTest.Querier) *handlers.DataHandler {
	return &handlers.DataHandler{
		Queries: queries,
	}
}
