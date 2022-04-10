package model

type ModelEntity interface {
	GetAllModel() []interface{}
}

type Model struct {
	models []interface{}
}

func NewModel() ModelEntity {
	model := Model{}

	roleModel := Role{}
	divisionModel := Division{}
	userModel := User{}
	projectModel := Project{}
	categoryModel := Category{}
	activityModel := Activity{}
	commentModel := Comment{}
	model.models = append(model.models,
		categoryModel,
		roleModel,
		divisionModel,
		userModel,
		projectModel,
		activityModel,
		commentModel,
	)
	return &model
}

func (m *Model) GetAllModel() []interface{} {
	return m.models
}
