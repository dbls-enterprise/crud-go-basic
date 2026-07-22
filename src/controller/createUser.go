package controller

import (
	"github.com/dbls-enterprise/crud-go-basic/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	restErr := rest_err.NewBadRequestValidationError(
		"Chamou a rota da forma errada.",
		"bad_request",
		400,
	)

	c.JSON(restErr.Code, restErr)
}
