package orm

import (
	"log"
	"pds-backend/config"
	"pds-backend/orm/gorm/model"

	"pds-backend/constant"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormOrmEntity interface {
	GetOrm() *gorm.DB
	RunMigration()
}

type GormOrm struct {
	orm    *gorm.DB
	models model.ModelEntity
}

func NewGormOrm(appConfig config.AppConfigEntity) GormOrmEntity {
	orm, err := gorm.Open(postgres.Open(appConfig.GetDataSourceName()), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: appConfig.GetDbSchema(), SingularTable: true}})
	if err != nil {
		log.Panicln(err)
	}

	models := model.NewModel()
	if appConfig.GetDebug() == constant.YES {
		orm = orm.Debug()
	}
	gormOrm := &GormOrm{orm: orm, models: models}

	if appConfig.GetRunMigration() == constant.YES {
		gormOrm.RunMigration()
	}

	if appConfig.GetDbSeed() == constant.YES {
		gormOrm.RunDbSeed()
	}
	return gormOrm
}

func (g *GormOrm) GetOrm() *gorm.DB {
	return g.orm
}

func (g *GormOrm) RunMigration() {
	err := g.orm.AutoMigrate(g.models.GetAllModel()...)
	if err != nil {
		log.Panicln(err)
	}
}

func (g *GormOrm) RunDbSeed() {
	roles := []string{
		"Front end Developer",
		"Back end Developer",
		"Mobile App Developer",
		"Data Scientist/Analyst",
		"Machine Learning Engineer",
		"Network Engineer",
		"UI/UX Designer",
		"Product Manager",
	}
	var roleModels []model.Role
	for i := 0; i < len(roles); i++ {
		roleModel := model.Role{Name: &roles[i]}
		roleModels = append(roleModels, roleModel)
	}

	categories := []string{
		"Backend Development",
		"Web Development",
		"Mobile App Development",
		"Network Development",
		"Data Analysis and Visualization",
		"Machine Learning",
		"Cybersecurities and Data Protection",
		"UI/UX Design",
	}
	var categoryModels []model.Category
	for i := 0; i < len(categories); i++ {
		categoryModel := model.Category{Name: &categories[i]}
		categoryModels = append(categoryModels, categoryModel)
	}

	divisions := []string{
		"Financial Service (Prodigy)",
		"Media (Genflix)",
		"Renewable Energy (SUNterra)",
		"Coal Mining (Berau Coal)",
		"Investment (Nanovest)",
	}
	var divisionModels []model.Division
	for i := 0; i < len(divisions); i++ {
		divisionModel := model.Division{Name: &divisions[i]}
		divisionModels = append(divisionModels, divisionModel)
	}

	g.GetOrm().Model(&model.Category{}).Save(categoryModels)
	g.GetOrm().Model(&model.Role{}).Save(roleModels)
	g.GetOrm().Model(&model.Division{}).Save(divisionModels)

}
