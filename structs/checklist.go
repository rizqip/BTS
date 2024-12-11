package structs

type Checklist struct {
	ChecklistId int64  `json:"ChecklistId,omitempty" form:"ChecklistId" query:"ChecklistId" gorm:"column:ChecklistId"`
	UserId      int64  `json:"UserId,omitempty" form:"UserId" query:"UserId" gorm:"column:UserId"`
	Title       string `json:"Title,omitempty" form:"Title" query:"Title" gorm:"column:Title"`
	Description string `json:"Description,omitempty" form:"Description" query:"Description" gorm:"column:Description"`
	IsDeleted   string `json:"IsDeleted,omitempty" form:"IsDeleted" query:"IsDeleted" gorm:"column:IsDeleted"`
	CreatedAt   string `json:"CreatedAt,omitempty" form:"CreatedAt" query:"CreatedAt" gorm:"column:CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt,omitempty" form:"UpdatedAt" query:"UpdatedAt" gorm:"column:UpdatedAt"`
}
