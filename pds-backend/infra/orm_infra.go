package infra

import (
	"pds-backend/config"
	orm "pds-backend/orm/gorm"
)

type OrmInfraEntity interface {
	GetGormOrm() orm.GormOrmEntity
}

type OrmInfra struct {
	gormOrm orm.GormOrmEntity
}

func NewOrmInfra(appConfig config.AppConfigEntity) OrmInfraEntity {
	gormOrm := orm.NewGormOrm(appConfig)
	return &OrmInfra{
		gormOrm: gormOrm,
	}
}

func (o *OrmInfra) GetGormOrm() orm.GormOrmEntity {
	return o.gormOrm
}
