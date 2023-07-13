package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

func Modification(c *gin.Context) {
	var modifyMsg static.TestMsg
	if err := c.BindJSON(&modifyMsg); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		fmt.Println(err.Error())
		return
	}
	fmt.Println(modifyMsg)
	if result := db.DB.Model(&static.TestMsg{}).Where("id = ?", modifyMsg.ID).Updates(&modifyMsg).Error; result != nil {
		c.JSON(500, gin.H{"message": result})
		fmt.Println(result)
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
	return
}
