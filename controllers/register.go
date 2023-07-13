package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/config"
	"nodejs_project/db"
	"nodejs_project/static"
)

type RegisterRequest struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Rpassword string `json:"rpassword"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	}
	var users []static.UserInfo
	if req.Account == "" || req.Password == "" {
		c.JSON(500, gin.H{"message": "账号或密码不能为空"})
		return
	}
	if len(req.Account) > 20 {
		c.JSON(400, gin.H{"message": "账号长度超过限制"})
		return
	}
	if len(req.Password) < 6 || len(req.Password) > 20 {
		c.JSON(400, gin.H{"message": "密码格式不正确"})
		return
	}
	db.DB.Model(&static.UserInfo{}).Where("account=?", req.Account).Find(&users)
	fmt.Println(users)
	if len(users) > 0 {
		c.JSON(500, gin.H{"message": "账号已存在"})
		return
	}
	if req.Password != req.Rpassword {
		c.JSON(500, gin.H{"message": "两次输入的密码不一致"})
		return
	}
	err := config.SaveUserInfo(req.Account, req.Password, req.Rpassword)
	if err == nil {
		c.JSON(200, gin.H{"message": "注册成功"})
		return
	} else {
		c.JSON(500, gin.H{"message": "服务器出错啦！"})
	}

}
