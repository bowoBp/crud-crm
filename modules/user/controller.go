package user

import (
	"crud/dto"
	"crud/entity"
	"fmt"
	"time"
)

type ControllerUser interface {
	CreateUser(req UserParam) (any, error)
	GetUserById(id uint) (FindUser, error)
	UpdateUser(req UserParam, id uint) (any, error)
	DeleteUser(email string) (any, error)
	GetAllUser(page uint, usernameStr string) (FindAllCustomer, error)
}

type controllerUser struct {
	userUseCase UseCaseUser
}

func (uc controllerUser) CreateUser(req UserParam) (any, error) {

	user, err := uc.userUseCase.CreateUser(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: UserParam{
			Name:   user.Name,
			Email:  user.Email,
			Gender: user.Gender,
			Phone:  user.Phone,
		},
	}
	return res, nil
}

func (uc controllerUser) GetUserById(id uint) (FindUser, error) {
	var res FindUser
	user, err := uc.userUseCase.GetUserById(id)
	if err != nil {
		return FindUser{}, err
	}
	res.Data = user
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get user",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerUser) UpdateUser(req UserParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.userUseCase.UpdateUser(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerUser) DeleteUser(email string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.userUseCase.DeleteUser(email)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}

func (c controllerUser) GetAllUser(page uint, usernameStr string) (FindAllCustomer, error) {
	start := time.Now()
	page, perPage, total, totalPages, customerEntities, err := c.userUseCase.GetAllUser(page, usernameStr)

	if err != nil {
		return FindAllCustomer{}, err
	}

	data := make([]entity.User, len(customerEntities))
	for i, customerEntity := range customerEntities {
		data[i] = customerEntity
	}

	res := FindAllCustomer{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success find customer",
			Message:      "Success find all",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
		Data:       data,
	}

	return res, nil
}
