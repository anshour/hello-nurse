package v1routes

import (
	userController "hello-nurse/src/http/controllers/user"
	middleware "hello-nurse/src/utils/middlewares"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userController := userController.New(&userController.V1User{
		DB: i.DB,
	})
	g.POST("/it/register", userController.ITRegister)
	g.POST("/it/login", userController.ITLogin)
	g.POST("/nurse/login", userController.ITLogin)

	g.Use(middleware.Authentication(), middleware.OnlyRole("it"))

	g.GET("/", userController.UserList)
	g.POST("/nurse/register", userController.NurseRegister)
	g.PUT("/nurse/:userId", userController.NurseEdit)
	g.DELETE("/nurse/:userId", userController.NurseDelete)
	g.POST("/nurse/:userId/access", userController.NurseAccess)

}
