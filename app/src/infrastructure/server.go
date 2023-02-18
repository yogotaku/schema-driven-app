package infrastructure

import "github.com/yogotaku/schema-driven-app/app/src/interface/controllers"

type ApiServer struct {
	*controllers.UserController
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		UserController: controllers.NewUserController(),
	}
}
