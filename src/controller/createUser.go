package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dbls-enterprise/crud-go-basic/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewRestErr("error creating user", "internal_server_error", 500, nil)
	c.JSON(err.Code, err)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
}