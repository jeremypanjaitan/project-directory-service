package routes

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"pds-backend/constant"
// 	"pds-backend/delivery/httpdelivery/jsonmodels"
// 	"pds-backend/orm/gorm/model"
// 	"pds-backend/usecase"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// var name = "fahmi uhuy"
// var name2 = "diva ihiy"
// var email = "fahmi@gmail.com"
// var email2 = "diva@gmail.com"
// var password = "test123"
// var password2 = "asd123f"
// var gender = "M"
// var gender2 = "F"
// var bio = "ini bio loh"
// var bio2 = "itu bio loh"
// var div = uint(2)
// var div2 = uint(3)
// var role = uint(1)
// var role2 = uint(2)
// var dummyUsers = []model.User{
// 	{
// 		FullName:   &name,
// 		Email:      &email,
// 		Password:   &password,
// 		Picture:    nil,
// 		Gender:     &gender,
// 		Biography:  &bio,
// 		DivisionID: &div,
// 		RoleID:     &role,
// 	},
// 	{
// 		FullName:   &name2,
// 		Email:      &email2,
// 		Password:   &password2,
// 		Picture:    nil,
// 		Gender:     &gender2,
// 		Biography:  &bio2,
// 		DivisionID: &div2,
// 		RoleID:     &role2,
// 	},
// }

// var dummyPass = jsonmodels.UserChangePassword{
// 	OldPassword: "hehe",
// 	NewPassword: "haha",
// }

// type UserMockResponse struct {
// 	StatusCode  int    `json:"status"`
// 	Description string `json:"description"`
// 	Data        interface{}
// }

// type UserMockErrorResponse struct {
// 	Message string
// }

// type userUsecaseMock struct {
// 	mock.Mock
// }

// type UserApiTestSuite struct {
// 	suite.Suite
// 	usecaseTest     usecase.UserUsecaseEntity
// 	routerTest      *gin.Engine
// 	routerGroupTest *gin.RouterGroup
// }

// func (suite *UserApiTestSuite) SetupTest() {
// 	suite.usecaseTest = new(userUsecaseMock)
// 	suite.routerTest = gin.Default()
// 	suite.routerGroupTest = suite.routerTest.Group("/apimock")
// }

// func TestUserApiTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserApiTestSuite))
// }

// func (u *userUsecaseMock) GetProfileUser(email string) (*model.User, *string, *string, error) {
// 	args := u.Called(email)
// 	log.Println(args)
// 	if args.Get(0) == nil {
// 		return nil, nil, nil, args.Error(3)
// 	}
// 	return args.Get(0).(*model.User), args.Get(1).(*string), args.Get(2).(*string), nil
// }

// func (u *userUsecaseMock) GetProfileUserById(id uint) (*model.User, *string, *string, error) {
// 	args := u.Called(id)
// 	if args.Get(0) == nil || args.Get(1) == nil || args.Get(2) == nil {
// 		return nil, nil, nil, args.Error(3)
// 	}
// 	return args.Get(0).(*model.User), args.Get(1).(*string), args.Get(2).(*string), nil
// }

// func (u *userUsecaseMock) ChangePassword(email string, oldPassword string, newPassword string) error {
// 	args := u.Called(email, oldPassword, newPassword)
// 	if args.Get(0) != nil {
// 		return args.Error(0)
// 	}
// 	return nil
// }

// func (u *userUsecaseMock) ChangeProfile(email string, updatedProfile model.User) (*model.User, *string, *string, error) {
// 	args := u.Called(email, updatedProfile)
// 	if args.Get(0) == nil || args.Get(1) == nil || args.Get(2) == nil {
// 		return nil, nil, nil, args.Error(3)
// 	}
// 	return args.Get(0).(*model.User), args.Get(1).(*string), args.Get(2).(*string), nil
// }

// func (suite *UserApiTestSuite) TestUserApi_GetProfilerHandler_Success() {
// 	dummyUser := dummyUsers[0]
// 	// dummyEmail := dummyUser.Email
// 	dummyDivision := "electrical"
// 	dummyRole := "electrician"
// 	suite.usecaseTest.(*userUsecaseMock).On("GetProfileUser", "").Return(&dummyUser, &dummyDivision, &dummyRole, nil)
// 	delivery := UserRoute{userUsecase: suite.usecaseTest}
// 	handler := delivery.getProfileHandler
// 	profileRoute := suite.routerGroupTest.Group("/profile")
// 	profileRoute.GET("", handler)

// 	rr := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(rr)
// 	ctx.Set(constant.EMAIL, "gongon@gmail.com")
// 	req, _ := http.NewRequest(http.MethodGet, "/apimock/profile", nil)

// 	suite.routerTest.ServeHTTP(rr, req)
// 	assert.Equal(suite.T(), 200, rr.Code)

// 	// a := rr.Body.String()
// 	// actualProfile := new(UserMockResponse)
// 	// json.Unmarshal([]byte(a), actualProfile)
// 	// assert.Equal(suite.T(), dummyUser, actualProfile.Data)
// }

// // func (suite *UserApiTestSuite) TestUserApi_GetProfilerHandler_Failed() {
// // 	dummyUser := dummyUsers[0]
// // 	dummyEmail := dummyUser.Email
// // 	suite.usecaseTest.(*userUsecaseMock).On("GetProfileUser", dummyEmail).Return(nil, nil, nil, errors.New("failed"))
// // 	delivery := UserRoute{UserUsecase: suite.usecaseTest}
// // 	handler := delivery.GetProfileHandler
// // 	profileRoute := suite.routerGroupTest.Group("/profile")
// // 	profileRoute.GET("", handler)

// // 	rr := httptest.NewRecorder()
// // 	req, _ := http.NewRequest(http.MethodGet, "/apiMock/profile", nil)

// // 	suite.routerTest.ServeHTTP(rr, req)
// // 	assert.Equal(suite.T(), 500, rr.Code)

// // 	// a := rr.Body.String()
// // 	// errorMessage := new(UserMockErrorResponse)
// // 	// json.Unmarshal([]byte(a), errorMessage)
// // 	// log.Println("test",a)
// // 	// assert.Equal(suite.T(), nil, actualProfile.Data)
// // }

// func (suite *UserApiTestSuite) TestUserApi_ChangePassword_Success() {
// 	dummyEmail := ""
// 	suite.usecaseTest.(*userUsecaseMock).On("ChangePassword", dummyEmail, dummyPass.OldPassword, dummyPass.NewPassword).Return(nil)
// 	delivery := UserRoute{userUsecase: suite.usecaseTest}
// 	handler := delivery.changePasswordHandler
// 	profileRoute := suite.routerGroupTest.Group("/profile")
// 	profileRoute.PUT("/password", handler)

// 	rr := httptest.NewRecorder()
// 	reqBody, _ := json.Marshal(dummyPass)
// 	req, _ := http.NewRequest(http.MethodPut, "/apimock/profile/password", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	suite.routerTest.ServeHTTP(rr, req)
// 	assert.Equal(suite.T(), 200, rr.Code)

// 	a := rr.Body.String()
// 	actual := new(UserMockResponse)
// 	json.Unmarshal([]byte(a), actual)
// 	log.Println("after parse", actual)
// 	assert.Equal(suite.T(), "successfully update password", actual.Description)
// }

// func (suite *UserApiTestSuite) TestUserApi_ChangePassword_Failed() {
// 	dummyEmail := ""
// 	suite.usecaseTest.(*userUsecaseMock).On("ChangePassword", dummyEmail, dummyPass.OldPassword, dummyPass.NewPassword).Return(errors.New("failed"))
// 	delivery := UserRoute{userUsecase: suite.usecaseTest}
// 	handler := delivery.changePasswordHandler
// 	profileRoute := suite.routerGroupTest.Group("/profile")
// 	profileRoute.PUT("/password", handler)

// 	rr := httptest.NewRecorder()
// 	reqBody, _ := json.Marshal(dummyPass)
// 	req, _ := http.NewRequest(http.MethodPut, "/apimock/profile/password", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	suite.routerTest.ServeHTTP(rr, req)
// 	assert.Equal(suite.T(), 400, rr.Code)

// 	a := rr.Body.String()
// 	actual := new(UserMockResponse)
// 	json.Unmarshal([]byte(a), actual)
// 	log.Println("after parse", actual)
// 	assert.Equal(suite.T(), "error required field", actual.Description)
// }

// // func (suite *UserApiTestSuite) TestUserApi_ChangeProfileHandler_Success() {
// // 	// dummyUser := dummyUsers[0]
// // 	dummyDivi := "hehe"
// // 	dummyRole := "haha"
// // 	suite.usecaseTest.(*userUsecaseMock).On("ChangeProfile", "", dummyUsers[1]).Return(&dummyUsers[1], &dummyDivi, &dummyRole, nil)
// // 	delivery := UserRoute{userUsecase: suite.usecaseTest}
// // 	handler := delivery.changeProfileHandler
// // 	profileRoute := suite.routerGroupTest.Group("/profile")
// // 	profileRoute.PUT("", handler)

// // 	rr := httptest.NewRecorder()
// // 	reqBody, _ := json.Marshal(dummyUsers[1])
// // 	req, _ := http.NewRequest(http.MethodPut, "/apimock/profile", bytes.NewBuffer(reqBody))
// // 	req.Header.Set("Content-Type", "application/json")

// // 	suite.routerTest.ServeHTTP(rr, req)
// // 	assert.Equal(suite.T(), 200, rr.Code)
// // }
