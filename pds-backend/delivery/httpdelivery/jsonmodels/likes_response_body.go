package jsonmodels

type LikesResponseBody struct {
	TotalLikes uint `json:"totalLikes"`
	IsUserLike bool `json:"isUserLike"`
}
