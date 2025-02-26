package handlers

import (
	"apitest/internal/web/users"
	"apitest/userService"
	"context"
	"fmt"
)

type Handlers struct {
	Service *userService.UserService
}

func (h Handlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUser()
	if err != nil {
		return users.GetUsers500JSONResponse{}, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.Id,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}
func (h Handlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return users.PostUsers500JSONResponse{}, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.Id,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	fmt.Println(userToCreate)
	return response, nil
}

func (h Handlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := uint(request.Id)
	err := h.Service.DeleteUser(id)
	if err != nil {
		return users.DeleteUsersId500JSONResponse{}, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h Handlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := uint(request.Id)
	body := request.Body
	if body == nil {
		return users.PatchUsersId500JSONResponse{}, nil
	}
	userValue := ""
	if body.Email != nil {
		userValue = *body.Email
	}
	PasswordValue := "" // Значение по умолчанию для isDone
	if body.Password != nil {
		PasswordValue = *body.Password
	}
	updatedUser := userService.User{
		Id:       id,
		Email:    userValue,
		Password: PasswordValue,
	}
	upUser, err := h.Service.UpdateUser(id, updatedUser)
	if err != nil {
		return users.PatchUsersId500JSONResponse{}, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &upUser.Id,
		Email:    &upUser.Email,
		Password: &upUser.Password,
	}
	fmt.Println(id)
	return response, nil
}

func NewHandlers(service *userService.UserService) *Handlers {
	return &Handlers{
		Service: service,
	}
}
