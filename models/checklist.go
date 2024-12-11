package models

import (
	"BTS/structs"
	"BTS/system"
	"time"
)

func CreateChecklist(req map[string]interface{}) (map[string]interface{}, error) {
	var responseChecklist map[string]interface{}
	dataReq := system.CreatedTimeNow(req)

	tx := ModelsDb.Table("Checklists").Create(dataReq).Limit(1).Order("ChecklistId Desc").Find(&responseChecklist)

	responseChecklist["CreatedAt"] = responseChecklist["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	responseChecklist["UpdatedAt"] = responseChecklist["UpdatedAt"].(time.Time).Format("2006-01-02 15:04:05")

	return responseChecklist, tx.Error
}

func ListChecklist(req structs.Filter) (map[string]interface{}, error) {
	var ListChecklist []map[string]interface{}

	tx := ModelsDb.Select("ChecklistId,Title")
	tx.Table("Checklists")

	if req.Search != "" {
		tx.Where("Title LIKE ?", "%"+req.Search+"%")
	}

	var count int64
	tx.Count(&count)

	if req.Limit > 0 {
		tx.Limit(req.Limit)
	}
	if req.Offset > 0 {
		tx.Offset(req.Offset)
	}

	tx.Find(&ListChecklist)

	resp := map[string]interface{}{
		"list":        ListChecklist,
		"total_items": count,
	}

	return resp, nil
}

func DetailChecklist(ChecklistId int64) (map[string]interface{}, error) {
	/* Get Detail Checklist */
	var DetailChecklist map[string]interface{}
	ModelsDb.Table("Checklists").Where("ChecklistId = ?", ChecklistId).Find(&DetailChecklist)

	/* Add New Response */
	DetailChecklist["CreatedAt"] = DetailChecklist["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	DetailChecklist["UpdatedAt"] = DetailChecklist["UpdatedAt"].(time.Time).Format("2006-01-02 15:04:05")

	return DetailChecklist, nil
}

func UpdateChecklist(req map[string]interface{}) (map[string]interface{}, error) {
	var responseChecklist map[string]interface{}
	dataReq := system.UpdatedTimeNow(req)

	tx := ModelsDb.Table("Checklists").Where("ChecklistId = ?", req["ChecklistId"]).Updates(dataReq).Find(&responseChecklist)

	responseChecklist["CreatedAt"] = responseChecklist["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	responseChecklist["UpdatedAt"] = responseChecklist["UpdatedAt"].(time.Time).Format("2006-01-02 15:04:05")

	return responseChecklist, tx.Error
}
