package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type SearchReq struct {
	KeyWord string `json:"keyWord"`
}

func Search(c *gin.Context) {
	var keyWord SearchReq
	var result []static.TestMsg
	if err := c.BindJSON(&keyWord); err != nil {
		c.JSON(400, gin.H{"message": "请求参数错误"})
		return
	}
	condition := "%" + keyWord.KeyWord + "%"
	err := db.DB.Model(&static.TestMsg{}).Where("id LIKE ?", condition).Or("topic LIKE ?", condition).
		Or("class LIKE ?", condition).Or("score LIKE ?", condition).Or("question_type LIKE ?", condition).
		Or("difficulty LIKE ?", condition).Or("author LIKE ?", condition).Or("answer LIKE ?", condition).
		Find(&result).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "查询错误"})
		return
	}
	fmt.Println(len(result), result)
	c.JSON(200, result)

}
