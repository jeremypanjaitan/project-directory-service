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

var nameRole1 = "Name1"
var nameRole2 = "Name2"

var dummyRoles = []model.Role{
	{
		Name: &nameRole1,
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
		Name: &nameRole2,
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

type RoleRepoTestSuite struct {
	suite.Suite
	orm       *gorm.DB
	ormEntity orm.GormOrmEntity
	mock      sqlmock.Sqlmock
}

func (suite *RoleRepoTestSuite) SetupTest() {
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

func (suite *RoleRepoTestSuite) TestRoleRepo_NewDivisionRepo() {
	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewRoleRepo(suite.ormEntity)

	assert.NotNil(suite.T(), repo)
}

func (suite *RoleRepoTestSuite) TestRoleRepo_GetAll_Success() {
	rows := sqlmock.NewRows([]string{"name", "id", "created_at", "updated_at", "deleted_at"})
	for _, d := range dummyRoles {
		rows.AddRow(d.Name, d.ID, d.CreatedAt, d.UpdatedAt, d.DeletedAt)
	}

	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" WHERE (.+) IS NULL`).WillReturnRows(rows)

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewRoleRepo(suite.ormEntity)
	all, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), all[0], dummyRoles[0])
}

func (suite *RoleRepoTestSuite) TestRoleRepo_GetAll_Failed() {
	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" WHERE (.+) IS NULL`).WillReturnError(errors.New("failed"))

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewRoleRepo(suite.ormEntity)
	all, err := repo.GetAll()

	assert.Nil(suite.T(), all)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *RoleRepoTestSuite) TestRoleRepo_GetById_Success() {
	row := sqlmock.NewRows([]string{"name", "id", "created_at", "updated_at", "deleted_at"})
	d := dummyRoles[0]
	row.AddRow(d.Name, d.ID, d.CreatedAt, d.UpdatedAt, d.DeletedAt)

	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" WHERE (.+) AND (.+) IS NULL ORDER BY (.+) LIMIT 1`).WithArgs(d.ID).WillReturnRows(row)

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewRoleRepo(suite.ormEntity)
	all, err := repo.GetById(d.ID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), all.ID, dummyRoles[0].ID)
}

func (suite *RoleRepoTestSuite) TestRoleRepo_GetById_Failed() {
	var id uint

	suite.mock.ExpectQuery(`SELECT (.+) FROM "roles" WHERE (.+) AND (.+) IS NULL ORDER BY (.+) LIMIT 1`).WithArgs(id).WillReturnError(errors.New("failed"))

	suite.ormEntity.(*gormOrmMock).On("GetOrm").Return(suite.orm)
	repo := NewRoleRepo(suite.ormEntity)
	selected, err := repo.GetById(id)

	assert.Nil(suite.T(), selected)
	assert.Equal(suite.T(), "failed", err.Error())
}

func TestRoleRepoTestSuite(t *testing.T) {
	suite.Run(t, new(RoleRepoTestSuite))
}
