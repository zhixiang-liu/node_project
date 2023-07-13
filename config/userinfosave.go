package config

import (
	"nodejs_project/db"
	"nodejs_project/static"
)

func SaveUserInfo(account, password, rpassword string) error {
	user := static.UserInfo{
		Account:   account,
		Password:  password,
		Rpassword: rpassword,
	}
	res := db.DB.Model(&static.UserInfo{}).Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
