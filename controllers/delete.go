package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type UserId struct {
	ID uint `json:"id"`
}

func Delete(c *gin.Context) {
	var id UserId
	if err := c.BindJSON(&id); err == nil {
		fmt.Println(id)
		if err := db.DB.Where("id=?", id.ID).Delete(&static.TestMsg{}).Error; err == nil {
			c.JSON(200, gin.H{"message": "删除成功"})
			return
		}
		c.JSON(500, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(500, gin.H{"message": "删除失败"})
	return
}
