package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pds-backend/apperror"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/service"
	"pds-backend/usecase"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type authUseCaseMock struct {
	mock.Mock
}

func (a *authUseCaseMock) Login(credential service.Credential) (*model.User, *service.TokenDetails, error) {
	args := a.Called(credential)
	if args.Get(0) == nil {
		return nil, nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Get(1).(*service.TokenDetails), args.Error(2)
}

func (a *authUseCaseMock) Logout(accessUuid string) error {
	args := a.Called(accessUuid)

	return args.Error(0)
}

func (a *authUseCaseMock) Register(user model.User) (*jsonmodels.RegisterReponseBody, error) {
	args := a.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jsonmodels.RegisterReponseBody), args.Error(1)
}

func (a *authUseCaseMock) GetDivisionName(divisionId uint) (*string, error) {
	args := a.Called(divisionId)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*string), args.Error(1)
}

func (a *authUseCaseMock) GetRoleName(roleId uint) (*string, error) {
	args := a.Called(roleId)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*string), args.Error(1)
}

type AuthApiTestSuite struct {
	suite.Suite
	authUseCaseTest usecase.AuthUsecaseEntity
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *AuthApiTestSuite) SetupTest() {
	suite.authUseCaseTest = new(authUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func (suite *AuthApiTestSuite) TestAuthApi_NewAuthRoute() {
	// NewAuthRoute(suite.routerGroupTest, suite.authUseCaseTest)
}

func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Success() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	loginRequestBodyDummy.Password = "mypassword"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, nil)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, nil)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, nil)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusOK)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Success_Bio_Pic_Empty() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	loginRequestBodyDummy.Password = "mypassword"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = nil
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = nil
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, nil)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, nil)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, nil)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusOK)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Error_Wrong_Credential() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	loginRequestBodyDummy.Password = "mypassword"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, apperror.ErrCredential)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, nil)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, nil)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusUnauthorized)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Error_RequiredField() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, nil)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, nil)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, nil)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusBadRequest)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Error_GetRoleName() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	loginRequestBodyDummy.Password = "mypassword"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, nil)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, apperror.ErrRoleNotFound)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, nil)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusBadRequest)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLoginHandler_Error_GetDivisionName() {
	var loginRequestBodyDummy jsonmodels.LoginRequestBody
	var credentialDummy service.Credential
	var userDummy model.User

	loginRequestBodyDummy.Email = "jeremypanjaitan@gmail.com"
	loginRequestBodyDummy.Password = "mypassword"
	credentialDummy.Email = loginRequestBodyDummy.Email
	credentialDummy.Password = loginRequestBodyDummy.Password

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	stringId := "23"
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	suite.authUseCaseTest.(*authUseCaseMock).On("Login", credentialDummy).Return(&userDummy, &service.TokenDetails{}, nil)
	suite.authUseCaseTest.(*authUseCaseMock).On("GetRoleName", roleIdUint).Return(&stringId, nil)

	suite.authUseCaseTest.(*authUseCaseMock).On("GetDivisionName", divisionIdUint).Return(&stringId, apperror.ErrDivisionNotFound)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLoginHandler

	loginRoute := suite.routerGroupTest.Group("/login")
	loginRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(loginRequestBodyDummy)

	request, _ := http.NewRequest(http.MethodPost, "/apimock/login", bytes.NewBuffer(reqBody))

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, http.StatusBadRequest)

}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLogoutHandler_Success() {
	mockAccessUuid := "4ff91736-9851-11ec-b909-0242ac120002"
	suite.authUseCaseTest.(*authUseCaseMock).On("Logout", mockAccessUuid).Return(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(constant.UUID, mockAccessUuid)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLogoutHandler
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusOK)
}
func (suite *AuthApiTestSuite) TestAuthApi_AuthLogoutHandler_Failed() {
	mockAccessUuid := "4ff91736-9851-11ec-b909-0242ac120002"
	suite.authUseCaseTest.(*authUseCaseMock).On("Logout", mockAccessUuid).Return(apperror.ErrCredential)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(constant.UUID, mockAccessUuid)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authLogoutHandler
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusBadRequest)
}
func (suite *AuthApiTestSuite) TestAuthApi_AuthRegisterHandler_Success() {
	var registerResponseBodyDummy jsonmodels.RegisterReponseBody
	var userDummy model.User

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)

	id, _ := strconv.ParseUint("20", 10, 32)
	idUint := uint(id)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	registerResponseBodyDummy.FullName = userDummy.FullName
	registerResponseBodyDummy.Gender = userDummy.Gender
	registerResponseBodyDummy.Email = userDummy.Email
	registerResponseBodyDummy.DivisionID = userDummy.DivisionID
	registerResponseBodyDummy.RoleID = userDummy.RoleID
	registerResponseBodyDummy.ID = &idUint

	suite.authUseCaseTest.(*authUseCaseMock).On("Register", userDummy).Return(&registerResponseBodyDummy, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authRegisterHandler
	reqBody, _ := json.Marshal(userDummy)
	c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusOK)
}
func (suite *AuthApiTestSuite) TestAuthApi_AuthRegisterHandler_Error_Required_Field() {
	var registerResponseBodyDummy jsonmodels.RegisterReponseBody
	var userDummy model.User

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)

	id, _ := strconv.ParseUint("20", 10, 32)
	idUint := uint(id)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	registerResponseBodyDummy.FullName = userDummy.FullName
	registerResponseBodyDummy.Gender = userDummy.Gender
	registerResponseBodyDummy.Email = userDummy.Email
	registerResponseBodyDummy.DivisionID = userDummy.DivisionID
	registerResponseBodyDummy.RoleID = userDummy.RoleID
	registerResponseBodyDummy.ID = &idUint

	suite.authUseCaseTest.(*authUseCaseMock).On("Register", userDummy).Return(&registerResponseBodyDummy, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authRegisterHandler
	reqBody, _ := json.Marshal(userDummy)
	c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusBadRequest)
}

func (suite *AuthApiTestSuite) TestAuthApi_AuthRegisterHandler_Error_Duplicate() {
	var registerResponseBodyDummy jsonmodels.RegisterReponseBody
	var userDummy model.User

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)

	id, _ := strconv.ParseUint("20", 10, 32)
	idUint := uint(id)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Picture = &picture
	userDummy.Email = &userEmail
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	registerResponseBodyDummy.FullName = userDummy.FullName
	registerResponseBodyDummy.Gender = userDummy.Gender
	registerResponseBodyDummy.Email = userDummy.Email
	registerResponseBodyDummy.DivisionID = userDummy.DivisionID
	registerResponseBodyDummy.RoleID = userDummy.RoleID
	registerResponseBodyDummy.ID = &idUint

	suite.authUseCaseTest.(*authUseCaseMock).On("Register", userDummy).Return(nil, apperror.ErrDuplicateValue)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authRegisterHandler
	reqBody, _ := json.Marshal(userDummy)
	c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusConflict)
}
func (suite *AuthApiTestSuite) TestAuthApi_AuthRegisterHandler_Error_Internal_Server_Error() {
	var registerResponseBodyDummy jsonmodels.RegisterReponseBody
	var userDummy model.User

	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)

	id, _ := strconv.ParseUint("20", 10, 32)
	idUint := uint(id)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Picture = &picture
	userDummy.Email = &userEmail
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	registerResponseBodyDummy.FullName = userDummy.FullName
	registerResponseBodyDummy.Gender = userDummy.Gender
	registerResponseBodyDummy.Email = userDummy.Email
	registerResponseBodyDummy.DivisionID = userDummy.DivisionID
	registerResponseBodyDummy.RoleID = userDummy.RoleID
	registerResponseBodyDummy.ID = &idUint

	suite.authUseCaseTest.(*authUseCaseMock).On("Register", userDummy).Return(nil, errors.New("another error"))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	delivery := AuthRoute{authUsecase: suite.authUseCaseTest}
	handler := delivery.authRegisterHandler
	reqBody, _ := json.Marshal(userDummy)
	c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
	handler(c)
	assert.Equal(suite.T(), w.Code, http.StatusInternalServerError)
}
func TestAuthApiSuite(t *testing.T) {
	suite.Run(t, new(AuthApiTestSuite))
}
