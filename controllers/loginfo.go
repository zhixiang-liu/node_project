package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type LogMsg struct {
	ID        uint   `json:"ID"`
	OpName    string `json:"opname"`
	OpUser    string `json:"opuser"`
	OpTime    string `json:"optime"`
	OpHash    string `json:"hash"`
	HashValue string `json:"hashvalue"`
}

func LogInfo(c *gin.Context) {
	var opmsg []static.LogInfo
	var displayOpmsg []LogMsg
	if err := db.DB.Model(&static.LogInfo{}).Find(&opmsg).Error; err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "查询失败"})
		return
	}
	for _, msg := range opmsg {
		displayOpmsg = append(displayOpmsg, LogMsg{
			ID:     msg.ID,
			OpName: msg.OpName,
			OpUser: msg.OpUser,
			OpTime: msg.UpdatedAt.Format("2006-01-02 15:04:05"),
			OpHash: msg.OpHash,
		})
	}
	c.JSON(200, displayOpmsg)
	return
}
