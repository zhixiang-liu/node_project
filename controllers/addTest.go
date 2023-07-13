package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

func AddTest(c *gin.Context) {
	var newTest static.TestMsg
	var addIntegral static.UserInfo
	if err := c.BindJSON(&newTest); err == nil {
		fmt.Println(newTest)
		if err := db.DB.Model(&static.TestMsg{}).Create(&newTest).Error; err == nil {
			db.DB.Model(&static.UserInfo{}).Where("account = ?", newTest.Author).Find(&addIntegral)
			addIntegral.Integral += 10
			db.DB.Save(&addIntegral)
			c.JSON(200, gin.H{"message": "添加成功"})
			return
		} else {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "添加失败"})
			return
		}
	}
	c.JSON(500, gin.H{"message": "添加失败"})
	return
}
