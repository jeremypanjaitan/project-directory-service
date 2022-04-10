package repo

import (
	"database/sql"
	"log"
	"pds-backend/orm/gorm/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var title = "project ini"
var title2 = "project itu"
var pic = "photo ini"
var pic2 = "photo itu"
var desc = "desc ini"
var desc2 = "desc itu"
var story = "story ini"
var story2 = "story itu"
var category = uint(1)
var category2 = uint(2)
var user = uint(1)
var user2 = uint(2)
var dummyProjects = []model.Project{
	{
		Title:       &title,
		Picture:     &pic,
		Description: &desc,
		Story:       &story,
		CategoryID:  &category,
		UserID:      &user,
		Model:       gorm.Model{ID: 1},
	},
	{
		Title:       &title2,
		Picture:     &pic2,
		Description: &desc2,
		Story:       &story2,
		CategoryID:  &category2,
		UserID:      &user2,
		Model:       gorm.Model{ID: 1},
	},
}

type ProjectRepoTestSuite struct {
	suite.Suite
	mockResource *gorm.DB
	mock         sqlmock.Sqlmock
}

func (suite *ProjectRepoTestSuite) SetupTest() {
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

func TestProjectRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectRepoTestSuite))
}

func (suite *ProjectRepoTestSuite) TestProjectRepo_CreateOne() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(20)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO (.+)").WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := ProjectRepo{orm: suite.mockResource}
	result, err := repo.CreateOne(dummyProjects[0])
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), uint(20), result.ID)
}

func (suite *ProjectRepoTestSuite) TestUserRepo_FindProjectById() {
	d := dummyProjects[0]
	rows := sqlmock.NewRows([]string{"title", "story"})
	rows.AddRow(d.Title, d.Story)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "projects" (.+)`).WithArgs(d.ID).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := ProjectRepo{orm: suite.mockResource}
	result, err := repo.FindProjectById(d.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "project ini", *result.Title)
}

func (suite *ProjectRepoTestSuite) TestUserRepo_FindIdByEmail() {
	dummyEmail := "gongon@gmail.com"
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(uint(2))
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" (.+)`).WithArgs(dummyEmail).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := ProjectRepo{orm: suite.mockResource}
	result, err := repo.FindIdByEmail(dummyEmail)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), uint(2), *result)
}

func (suite *ProjectRepoTestSuite) TestUserRepo_ShowListProjectWithPagination() {
	// d := dummyProjects[0]
	rows := sqlmock.NewRows([]string{"count"}).AddRow(int64(3))
	log.Println(rows)
	dummyPageNumber := uint(1)
	dummyPageSize := uint(3)
	dummyOffset := uint(3)
	suite.mock.ExpectQuery(`SELECT COUNT(.+) FROM "projects" (.+)`).WillReturnRows(rows)
	suite.mock.ExpectBegin()	
	suite.mock.ExpectQuery(`SELECT (.+) FROM "projects" (.+)`).WithArgs(dummyPageSize, dummyOffset).WillReturnRows(rows)
	suite.mock.ExpectCommit()

	repo := ProjectRepo{orm: suite.mockResource}
	_, _, err := repo.ShowListProjectWithPagination(dummyPageNumber, dummyPageSize)
	log.Println(err.Error())
	// assert.Nil(suite.T(), err)
}

func (suite *ProjectRepoTestSuite) TestUserRepo_FindProjectByTitle() {
	// d := dummyProjects[0]
	rows := sqlmock.NewRows([]string{"count"}).AddRow(int64(3))
	dummyPageNumber := uint(1)
	dummyPageSize := uint(3)
	dummySearch := "project"
	suite.mock.ExpectQuery(`SELECT COUNT(.*) FROM "projects" (.+)`).WithArgs(dummySearch).WillReturnRows(rows)
	suite.mock.ExpectBegin()	
	suite.mock.ExpectQuery(`SELECT (.+) FROM "projects" (.+)`).WithArgs(dummySearch).WillReturnError(nil)
	suite.mock.ExpectCommit()

	repo := ProjectRepo{orm: suite.mockResource}
	_, _, err := repo.FindProjectByTitle(dummyPageNumber, dummyPageSize, dummySearch)
	log.Println(err.Error())
	// assert.Nil(suite.T(), err)
}
