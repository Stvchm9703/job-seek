package model

import (
	"fmt"

	"gorm.io/gorm"
)

type SurveyUserPreferenceModel struct {
	gorm.Model
	User     UserAccountModel         `json:"user" gorm:"foreignKey:ID"`
	Keywords []PreferenceKeywordModel `json:"keywords" gorm:"foreignKey:ID"`
}

func (SurveyUserPreferenceModel) TableName() string {
	return "survey_user_preference"
}

func (SurveyUserPreferenceModel) DefineModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	return sd.AutoMigrate(&SurveyUserPreferenceModel{})

}
