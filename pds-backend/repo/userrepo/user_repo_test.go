package repo

import (
	"database/sql"
	"errors"
	"log"
	"pds-backend/orm/gorm/model"

	// "regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var name = "fahmi uhuy"
var name2 = "diva ihiy"
var email = "fahmi@gmail.com"
var email2 = "diva@gmail.com"
var password = "test123"
var password2 = "asd123f"
var gender = "M"
var gender2 = "F"
var bio = "ini bio loh"
var bio2 = "itu bio loh"
var div = uint(2)
var div2 = uint(3)
var role = uint(1)
var role2 = uint(2)
var dummyUsers = []model.User{
	{
		FullName:   &name,
		Email:      &email,
		Password:   &password,
		Picture:    nil,
		Gender:     &gender,
		Biography:  &bio,
		DivisionID: &div,
		RoleID:     &role,
	},
	{
		FullName:   &name2,
		Email:      &email2,
		Password:   &password2,
		Gender:     &gender2,
		Biography:  &bio2,
		DivisionID: &div2,
		RoleID:     &role2,
		Model:      gorm.Model{ID: 2},
	},
}

var upName = "ganti nama"
var upEmail = "ganti email"
var updatedDummy = model.User{
	FullName: &upName,
	Email:    &upEmail,
}

var passwordDummy = "passBaruw"

type UserRepoTestSuite struct {
	suite.Suite
	mockResource *gorm.DB
	mock         sqlmock.Sqlmock
}

func (suite *UserRepoTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)
	db, suite.mock, err = sqlmock.New()
	if err != nil {
		log.Panicln(err)
	}

	if db == nil {
		log.Panicln("mock db is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	suite.mockResource, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to open gorm v2 db, got error: %v\n", err)
	}
}

// func (suite *UserRepoTestSuite) TearDownTest() {
// 	suite.mockResource.Close()
// }

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestUserRepo_CreateOne() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(20)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO (.+)").WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	result, err := repo.CreateOne(dummyUsers[0])
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), uint(20), result.ID)
}

func (suite *UserRepoTestSuite) TestUserRepo_FindOne() {
	d := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"full_name", "email"})
	rows.AddRow(d.FullName, d.Email)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" (.+)`).WithArgs(d.Email).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	result, err := repo.FindOne(*d.Email)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), *d.FullName, *result.FullName)
}

func (suite *UserRepoTestSuite) TestUserRepo_FindOne_Failed() {
	d := dummyUsers[0]
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" (.+)`).WithArgs(d.Email).WillReturnError(errors.New("failed"))
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	_, err := repo.FindOne(*d.Email)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *UserRepoTestSuite) TestUserRepo_FindOneById() {
	d := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"full_name", "email"})
	rows.AddRow(d.FullName, d.Email)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" (.+)`).WithArgs(d.ID).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	result, err := repo.FindOneById(d.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), *d.FullName, *result.FullName)
}

func (suite *UserRepoTestSuite) TestUserRepo_FindOneById_Failed() {
	d := dummyUsers[0]
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" (.+)`).WithArgs(d.ID).WillReturnError(errors.New("failed"))
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	_, err := repo.FindOneById(d.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *UserRepoTestSuite) TestUserRepo_FindRoleByRoleId() {
	d := dummyUsers[0]
	dummyRole := "electrical"
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(d.RoleID, dummyRole)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" (.+)`).WithArgs(&d.RoleID).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	result, err := repo.FindRoleByRoleId(*d.RoleID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyRole, *result)
}

func (suite *UserRepoTestSuite) TestUserRepo_FindRoleByRoleId_Failed() {
	d := dummyUsers[0]
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" (.+)`).WithArgs(&d.RoleID).WillReturnError(errors.New("failed"))
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	_, err := repo.FindRoleByRoleId(*d.RoleID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *UserRepoTestSuite) TestUserRepo_FindDivisionByDivisionId() {
	d := dummyUsers[0]
	dummyDivision := "BE"
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(d.DivisionID, dummyDivision)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" (.+)`).WithArgs(&d.DivisionID).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	result, err := repo.FindDivisionByDivisionId(*d.DivisionID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDivision, *result)
}

func (suite *UserRepoTestSuite) TestUserRepo_FindDivisionByDivisionId_Failed() {
	d := dummyUsers[0]
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" (.+)`).WithArgs(&d.DivisionID).WillReturnError(errors.New("failed"))
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	_, err := repo.FindDivisionByDivisionId(*d.DivisionID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *UserRepoTestSuite) TestUserRepo_UpdatePassword() {
	d := dummyUsers[0]
	dummyNewPass := "ihiy"
	rows := sqlmock.NewRows([]string{"password"})
	rows.AddRow(d.Password)	
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`UPDATE 'users' SET (.+)`).WithArgs(dummyNewPass, *d.Email, *d.Password).WillReturnError(nil)
	suite.mock.ExpectCommit()

	// repo := UserRepo{orm: suite.mockResource}
	// err := repo.UpdatePassword(*d.Email, *d.Password, dummyNewPass)
	// log.Println(err.Error())
	// assert.Nil(suite.T(), err)
}


func (suite *UserRepoTestSuite) TestUserRepo_UpdateProfile() {
	d := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"full_name", "email"})
	rows.AddRow(*d.FullName, *d.Email)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`UPDATE (.+) SET (.+) WHERE (.+)`).WithArgs(*d.Email).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := UserRepo{orm: suite.mockResource}
	repo.UpdateProfile(*d.Email, d)
}