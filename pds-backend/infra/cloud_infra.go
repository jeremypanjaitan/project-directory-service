package infra

import cloudengine "pds-backend/cloudengine/firebase"

type CloudInfraEntity interface {
	GetFirebaseCloudEngine() cloudengine.FirebaseCloudEngineEntity
}

type CloudInfra struct {
	firebaseCloudEngine cloudengine.FirebaseCloudEngineEntity
}

func NewCloudInfra() CloudInfraEntity {
	firebaseCloudEngine := cloudengine.NewFirebaseCloudEngine()
	return &CloudInfra{firebaseCloudEngine: firebaseCloudEngine}
}

func (c *CloudInfra) GetFirebaseCloudEngine() cloudengine.FirebaseCloudEngineEntity {
	return c.firebaseCloudEngine
}
