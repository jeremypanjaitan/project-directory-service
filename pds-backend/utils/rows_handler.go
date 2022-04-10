package utils

import (
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
)

func RowsHandler(list []model.ProjectWithLikeViewComment) []jsonmodels.ProjectListResponseBody {
	var rows []jsonmodels.ProjectListResponseBody
	for i := 0; i < len(list); i++ {
		row := jsonmodels.ProjectListResponseBody{
			ID:            list[i].ID,
			Title:         *list[i].Title,
			Picture:       *list[i].Picture,
			Description:   *list[i].Description,
			CategoryID:    *list[i].CategoryID,
			TotalLikes:    *list[i].TotalLikes,
			TotalViews:    *list[i].TotalViews,
			TotalComments: *list[i].TotalComments,
		}
		rows = append(rows, row)
	}
	return rows
}
