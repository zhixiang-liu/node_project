package controllers

import (
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

func PointsAcquistiom(c *gin.Context) {
	var req static.UserInfo
	var res static.UserInfo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "数据格式不正确"})
		return
	}
	if err := db.DB.Model(&static.UserInfo{}).Where("account=?", req.Account).Find(&res).Error; err != nil {
		c.JSON(500, gin.H{"message": "数据查询失败"})
		return
	}
	c.JSON(200, res)
	return
}
