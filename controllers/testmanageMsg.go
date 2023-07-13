package controllers

import (
	"github.com/gin-gonic/gin"
	"nodejs_project/db"
	"nodejs_project/static"
)

type DisplayMsg struct {
	ID           uint   `gorm:"primary_key"`
	UpdatedAt    string `json:"updated_at"`
	Topic        string `json:"topic"`
	Class        string `json:"class"`
	Score        string `json:"score"`
	QuestionType string `json:"questiontype"`
	Difficulty   string `json:"difficulty"`
	Author       string `json:"author"`
	Answer       string `json:"answer"`
}

func TestmanageMsg(c *gin.Context) {
	var msgs []static.TestMsg
	db.DB.Model(&static.TestMsg{}).Find(&msgs)
	var displayMsgs []DisplayMsg
	for _, msg := range msgs {
		displayMsgs = append(displayMsgs, DisplayMsg{
			ID:           msg.ID,
			UpdatedAt:    msg.UpdatedAt.Format("2006-01-02 15:04:05"),
			Topic:        msg.Topic,
			Class:        msg.Class,
			Score:        msg.Score,
			QuestionType: msg.QuestionType,
			Difficulty:   msg.Difficulty,
			Author:       msg.Author,
			Answer:       msg.Answer,
		})
	}
	//fmt.Println(displayMsgs)
	c.JSON(200, displayMsgs)
}
