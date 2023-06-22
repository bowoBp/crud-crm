package user

import (
	"crud/dto"
	"crud/entity"
)

type UserParam struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data UserParam `json:"data"`
}

type FindUser struct {
	dto.ResponseMeta
	Data entity.User `json:"data"`
}

type FindAllCustomer struct {
	dto.ResponseMeta
	Page       uint          `json:"page,omitempty"`
	PerPage    uint          `json:"per_page,omitempty"`
	Total      int           `json:"total,omitempty"`
	TotalPages uint          `json:"total_pages,omitempty"`
	Data       []entity.User `json:"data"`
}
