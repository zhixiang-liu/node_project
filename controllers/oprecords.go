package controllers

import (
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type OpMsg struct {
	OpName string `json:"OpName"`
	OpUser string `json:"OpUser"`
	OpHash string `json:"OpHash"`
}

func OpRecord(c *gin.Context) {
	var oprecords OpMsg
	if err := c.BindJSON(&oprecords); err != nil {
		c.JSON(400, gin.H{"message": "数据格式不正确"})
		return
	}
	if err := db.DB.Model(&static.LogInfo{}).Create(&static.LogInfo{
		OpName: oprecords.OpName,
		OpUser: oprecords.OpUser,
		OpHash: oprecords.OpHash,
	}).Error; err != nil {
		c.JSON(500, gin.H{"message": "数据存储错误"})
		return
	}
	c.JSON(200, gin.H{"message": "数据记录成功"})
}
