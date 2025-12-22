package domain

type SystemSetting struct {
	Key   string `json:"key" validate:"required" gorm:"primaryKey;size:255"`
	Value string `json:"value" gorm:"type:text"`
	Detail string `json:"detail" gorm:"type:text"`
}

func (ss *SystemSetting) TableName() string {
	return "system_settings"
}

// SystemSetting doesn't use BaseModel, so no BeforeCreate hook needed

