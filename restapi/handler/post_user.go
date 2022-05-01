package handler

import (
    "github.com/go-openapi/runtime/middleware"
    "go-swagger/models"
    "go-swagger/restapi/operations"
	"fmt"
)

var (
    Users map[uint64]models.User
)

func init()  {
    Users = map[uint64]models.User{}
}

type PostUserHandler struct {}

func (h *PostUserHandler) Handle(params operations.PostUserParams) middleware.Responder {
	fmt.Println(params)
    u := models.User{
        ID:params.Body.ID,
        Name: params.Body.Name,
    }

    Users[params.Body.ID] = u

	return operations.NewPostUserCreated()
}