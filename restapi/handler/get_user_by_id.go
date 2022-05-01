package handler

import (
    "github.com/go-openapi/runtime/middleware"
    "strconv"
    "go-swagger/restapi/operations"
)

type GetUserByIDHandler struct {}

func (h *GetUserByIDHandler) Handle(params operations.GetUserByIDParams) middleware.Responder {
    id, _ := strconv.ParseUint(params.ID, 10, 64)

    u := Users[id]

    return operations.NewGetUserByIDOK().WithPayload(&u)
}