package controllers

import (
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	if req.Account == "" || req.Password == "" {
		c.JSON(501, gin.H{"message": "账号密码不能为空"})
		return
	}
	if len(req.Account) > 20 {
		c.JSON(400, gin.H{"message": "账号长度超过限制"})
		return
	}
	var res = LoginRequest{}
	err := db.DB.Model(&static.UserInfo{}).Where("account=?", req.Account).Limit(1).Find(&res)
	if err == nil {
		c.JSON(200, gin.H{"error": err})
		return
	}
	if res.Account == "" {
		c.JSON(501, gin.H{"message": "账号不存在"})
		return
	}
	if req.Password != res.Password {
		c.JSON(502, gin.H{"message": "密码错误"})
		return
	}
	c.JSON(200, gin.H{"message": "登陆成功", "state": "200"})
}
