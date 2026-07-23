package brunotesting

import "github.com/gin-gonic/gin"

func Testing(c *gin.Context) {
	nickname := c.Param("nickname")
	senha := c.Param("senha")

	if nickname == "davidlaryssa" && senha == "1234" {
		c.JSON(200, gin.H{
			"message": "Usuário autenticado com sucesso!",
		})
		return
	}

	c.JSON(401, gin.H{
		"message": "Falha na autenticação.",
	})
}
