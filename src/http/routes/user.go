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

	group := i.Echo.Group("/user")
	group.Use(middleware.Authentication(), middleware.OnlyRole("it"))

	group.GET("", userController.UserList)
	group.POST("/nurse/register", userController.NurseRegister)
	group.PUT("/nurse/:userId", userController.NurseEdit)
	group.DELETE("/nurse/:userId", userController.NurseDelete)
	group.POST("/nurse/:userId/access", userController.NurseAccess)

}
