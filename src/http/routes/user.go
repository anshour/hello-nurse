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

	g.Use(middleware.Authentication())
	g.GET("/", userController.UserList, middleware.OnlyRole("it"))
	g.POST("/nurse/register", userController.NurseRegister, middleware.OnlyRole("it"))
	g.PUT("/nurse/:userId", userController.NurseEdit, middleware.OnlyRole("it"))
	g.DELETE("/nurse/:userId", userController.NurseDelete, middleware.OnlyRole("it"))
	g.POST("/nurse/:userId/access", userController.NurseAccess, middleware.OnlyRole("it"))

}
