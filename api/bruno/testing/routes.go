package brunotesting

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/Testing/:nickname/:senha", TestingHandler)
}

func TestingHandler(c *gin.Context) {
	nickname := c.Param("nickname")
	senha := c.Param("senha")

	if nickname == "davidlaryssa" && senha == "1234" {
		c.JSON(http.StatusOK, gin.H{
			"message": "autenticado",
			"usuario": nickname,
			"senha":   senha,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "acesso negado",
		"usuario": nickname,
		"senha":   senha,
	})
}
