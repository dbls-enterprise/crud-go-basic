package routes

import (
	brunotesting "github.com/dbls-enterprise/crud-go-basic/api/bruno/testing"
	"github.com/dbls-enterprise/crud-go-basic/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

	// testes isolados para implementações futuras
	brunotesting.InitRoutes(r) // Adiciona as rotas de teste do pacote brunotesting
}
