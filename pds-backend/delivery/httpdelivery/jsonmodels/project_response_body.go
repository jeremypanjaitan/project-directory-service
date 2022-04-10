package jsonmodels

import "time"

type ProjectResponseBody struct {
	Title         string    `json:"title"`
	Picture       string    `json:"picture"`
	Description   string    `json:"description"`
	Story         string    `json:"story"`
	CategoryID    uint      `json:"categoryId"`
	UserID        uint      `json:"userId"`
	CanDelete     bool      `json:"canDelete"`
	CanEdit       bool      `json:"canEdit"`
	TotalLikes    uint      `json:"totalLikes"`
	TotalViews    uint      `json:"totalViews"`
	TotalComments uint      `json:"totalComments"`
	CreatedAt     time.Time `json:"createdAt"`
}

type UpdatedProjectResponseBody struct {
	Title       string    `json:"title"`
	Picture     string    `json:"picture"`
	Description string    `json:"description"`
	Story       string    `json:"story"`
	CategoryID  uint      `json:"categoryId"`
	UserID      uint      `json:"userId"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProjectListResponseBody struct {
	ID            uint   `json:"ID"`
	Title         string `json:"title"`
	Picture       string `json:"picture"`
	Description   string `json:"description"`
	CategoryID    uint   `json:"categoryId"`
	TotalLikes    uint   `json:"totalLikes"`
	TotalViews    uint   `json:"totalViews"`
	TotalComments uint   `json:"totalComments"`
}

type ListProject struct {
	Row         []ProjectListResponseBody `json:"row"`
	CurrentPage uint                      `json:"currentPage"`
	TotalPage   uint                      `json:"totalPage"`
}
