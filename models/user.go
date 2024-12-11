package models

import (
	"BTS/structs"
	"BTS/system"
	"time"
)

func Register(req map[string]interface{}) (map[string]interface{}, error) {
	var responseUser map[string]interface{}
	dataReq := system.CreatedTimeNow(req)

	tx := ModelsDb.Table("Users").Create(dataReq).Limit(1).Order("UserId Desc").Find(&responseUser)

	responseUser["CreatedAt"] = responseUser["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	responseUser["UpdatedAt"] = responseUser["UpdatedAt"].(time.Time).Format("2006-01-02 15:04:05")

	return responseUser, tx.Error
}

func CheckUser(email string) (alreadyRegis bool) {
	var User map[string]interface{}
	ModelsDb.Table("Users").Where("Email = ?", email).Find(&User)

	alreadyRegis = false
	if User != nil {
		alreadyRegis = true
	}

	return alreadyRegis
}

func GetUser(req structs.User) (map[string]interface{}, error) {
	var DataUser map[string]interface{}
	tx := ModelsDb.Table("Users")

	if req.UserId > 0 {
		tx.Where("UserId = ?", req.UserId)
	}
	if req.Email != "" {
		tx.Where("Email = ?", req.Email)
	}

	tx.Find(&DataUser)

	return DataUser, nil
}

func UpdateStatus(UserId, Status int64) error {
	tx := ModelsDb.Table("Users").Where("UserId = ?", UserId).Update("Status", Status)

	return tx.Error
}
