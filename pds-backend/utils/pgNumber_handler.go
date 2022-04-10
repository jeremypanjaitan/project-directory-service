package utils

import "pds-backend/delivery/httpdelivery/jsonmodels"

func PageNumberHandler(convPageNumber int, pagination uint, rows []jsonmodels.ProjectListResponseBody) []jsonmodels.ProjectListResponseBody {
	if uint(convPageNumber) > pagination {
		return []jsonmodels.ProjectListResponseBody{}
	} else {
		return rows
	}

}