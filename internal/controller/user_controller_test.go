package controller

import (
	"GinEchoCrud/tests"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CreateUser_Success(t *testing.T) {
	controller := UserController{userService: &tests.MockUserService{}}

	userJSON := `{"name":"Jon Snow","age":45,"email":"jon@labstack.com"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	assert.NoError(t, controller.CreateUser(ctx))
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Test_CreateUser_Fail(t *testing.T) {
	controller := UserController{userService: &tests.MockUserService{}}

	userJSON := `{"name":Jon Snow","age":45,"email":"jon@labstack.com"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	controller.CreateUser(ctx)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func Test_GetAllUsers_Success(t *testing.T) {
	controller := UserController{userService: &tests.MockUserService{}}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	assert.NoError(t, controller.GetAllUsers(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func Test_GetAllUsers_Fail(t *testing.T) {
	controller := UserController{userService: &tests.MockUserServiceFail{}}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	controller.GetAllUsers(ctx)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
