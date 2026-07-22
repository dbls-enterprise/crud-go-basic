package controller

import "github.com/gin-gonic/gin"

func dizerOla(c *gin.Context) {
	c.JSON(200, gin.H{
		"message_para_laryssa": "Oi, amor!",
		"data_namoro": "2025-08-30",
	})

	c.Request.Context()
}
