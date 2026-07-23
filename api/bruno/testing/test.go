package controller

import "github.com/gin-gonic/gin"

func Testing(c *gin.Context) {
    nickname	:= c.Param("nickname")
    senha 		:= 	c.Param("senha")

    if nickname == "davidlaryssa" && senha == "abcd1234" {
        c.JSON(200, gin.H{
            "message": "Usuário autenticado com sucesso!",
        })
		
    } else {
        c.JSON(401, gin.H{
            "message": "Falha na autenticação.",
        })
    }
}
