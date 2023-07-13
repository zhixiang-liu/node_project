package static

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Account   string `json:"account"  gorm:"not null"`
	Password  string `json:"password"  gorm:"not null"`
	Rpassword string `json:"rpassword"  gorm:"not null"`
	Integral  int    `json:"integral"  gorm:"not null"`
}

type TestMsg struct {
	gorm.Model
	Topic        string `json:"topic" default:""`
	Class        string `json:"class" default:""`
	Score        string `json:"score" default:""`
	QuestionType string `json:"questiontype" default:""`
	Difficulty   string `json:"difficulty" default:""`
	Author       string `json:"author" default:""`
	Answer       string `json:"answer" default:""`
}

type LogInfo struct {
	gorm.Model
	OpName string `json:"opName"`
	OpUser string `json:"opUser"`
	OpHash string `json:"opHash"`
	OpTime string `json:"opTime"`
}
