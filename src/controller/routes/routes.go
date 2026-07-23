package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/dbls-enterprise/crud-go-basic/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	// r.GET("/Testing/:nickname/:senha")
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}