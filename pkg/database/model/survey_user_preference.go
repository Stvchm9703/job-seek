package model

import (
	"fmt"

	"gorm.io/gorm"
)

type SurveyUserPreferenceModel struct {
	gorm.Model
	UserID   int                      `json:"-"`
	User     UserAccountModel         `json:"user" gorm:"foreignKey:UserID"`
	Keywords []PreferenceKeywordModel `json:"keywords" gorm:"many2many:survey_user_preference_keyword;"`
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
