package repo

import (
	"errors"
	"log"
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var nameDivision1 = "Name1"
var nameDivision2 = "Name2"

var dummyDivisions = []model.Division{
	{
		Name: &nameDivision1,
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{
				Valid: false,
			},
		},
	},
	{
		Name: &nameDivision2,
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{
				Valid: false,
			},
		},
	},
}

type gormOrmMock struct {
	mock.Mock
}

func (o *gormOrmMock) GetOrm() *gorm.DB {
	args := o.Called()

	return args.Get(0).(*gorm.DB)
}

func (o *gormOrmMock) RunMigration() {
	return
}

type DivisionRepoTestSuite struct {
	suite.Suite
	orm       *gorm.DB
	ormEntity orm.GormOrmEntity
	mock      sqlmock.Sqlmock
}

func (suite *DivisionRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		log.Fatalf("mock db is null")
	}

	if mock == nil {
		log.Fatalf("sqlmock is null")
	}
	suite.mock = mock

	dialector := postgres.New(postgres.Config{
		DSN:                  "miniproject",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	suite.orm, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open gorm v2 db, got error: %v", err)
	}

	if suite.orm == nil {
		log.Fatalf("gorm db is null")
	}
	suite.ormEntity = new(gormOrmMock)
}

func (suite *DivisionRepoTestSuite) TestDivisionRepo_NewDivisionRepo() {
	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewDivisionRepo(suite.ormEntity)

	assert.NotNil(suite.T(), repo)
}

func (suite *DivisionRepoTestSuite) TestCategoryRepo_GetAll_Success() {
	rows := sqlmock.NewRows([]string{"name", "id", "created_at", "updated_at", "deleted_at"})
	for _, d := range dummyDivisions {
		rows.AddRow(d.Name, d.ID, d.CreatedAt, d.UpdatedAt, d.DeletedAt)
	}

	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" WHERE (.+) IS NULL`).WillReturnRows(rows)

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewDivisionRepo(suite.ormEntity)
	all, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), all[0], dummyDivisions[0])
}

func (suite *DivisionRepoTestSuite) TestDivisionRepo_GetAll_Failed() {
	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" WHERE (.+) IS NULL`).WillReturnError(errors.New("failed"))

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewDivisionRepo(suite.ormEntity)
	all, err := repo.GetAll()

	assert.Nil(suite.T(), all)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *DivisionRepoTestSuite) TestDivisionRepo_GetById_Success() {
	row := sqlmock.NewRows([]string{"name", "id", "created_at", "updated_at", "deleted_at"})
	d := dummyDivisions[0]
	row.AddRow(d.Name, d.ID, d.CreatedAt, d.UpdatedAt, d.DeletedAt)

	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" WHERE (.+) AND (.+) IS NULL ORDER BY (.+) LIMIT 1`).WithArgs(d.ID).WillReturnRows(row)

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewDivisionRepo(suite.ormEntity)
	all, err := repo.GetById(d.ID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), all.ID, dummyDivisions[0].ID)
}

func (suite *DivisionRepoTestSuite) TestDivisionRepo_GetById_Failed() {
	var id uint

	suite.mock.ExpectQuery(`SELECT (.+) FROM "divisions" WHERE (.+) AND (.+) IS NULL ORDER BY (.+) LIMIT 1`).WithArgs(id).WillReturnError(errors.New("failed"))

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewDivisionRepo(suite.ormEntity)
	selected, err := repo.GetById(id)

	assert.Nil(suite.T(), selected)
	assert.Equal(suite.T(), "failed", err.Error())
}

func TestDivisionRepoTestSuite(t *testing.T) {
	suite.Run(t, new(DivisionRepoTestSuite))
}
