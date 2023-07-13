package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nodejs_project/db"
	"nodejs_project/static"
)

type Data struct {
	IDs []int `json:"ids"`
}

func BatchDelete(c *gin.Context) {
	var data Data
	if err := c.BindJSON(&data); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(data)
	result := db.DB.Where("id IN (?)", data.IDs).Delete(&static.TestMsg{})
	if result.RowsAffected > 0 {
		c.JSON(200, gin.H{"message": fmt.Sprintf("%d 条数据已删除", result.RowsAffected)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "没有数据被删除"})
	}

}
