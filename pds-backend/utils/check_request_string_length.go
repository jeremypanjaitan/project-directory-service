package utils

import "pds-backend/orm/gorm/model"

func CheckProjectResponseStringLength(project model.Project) bool {
	title := *project.Title
	desc := *project.Description
	story := *project.Story
	if len(title) >= 6 && len(title) <= 150 &&
		len(desc) >= 10 && len(desc) <= 300 &&
		len(story) >= 10 {
		return true
	}
	return false
}

func CheckProfileResponseStringLength(profile model.User) bool {
	fullName := *profile.FullName
	bio := *profile.Biography
	if len(fullName) <= 30 && len(bio) <= 150 {
		return true
	}
	return false
}