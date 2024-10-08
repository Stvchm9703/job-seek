// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"fmt"
	"job-seek/pkg/protos"
	"strconv"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type UserProfileModel struct {
	gorm.Model
	// Id            string                   `json:"id"`
	User          UserAccountModel         `json:"User" gorm:"foreignKey:ID"`
	Title         string                   `json:"Title"`
	Position      string                   `json:"Position"`
	Description   string                   `json:"Description"`
	Company       string                   `json:"Company"`
	CompanyDetail *CompanyDetailModel      `json:"CompanyDetail" gorm:"foreignKey:ID"`
	Salary        string                   `json:"Salary"`
	Type          string                   `json:"Type"`
	Keywords      []PreferenceKeywordModel `json:"Keywords" gorm:"foreignKey:ID"`
	StartDate     string                   `json:"StartDate"`
	EndDate       string                   `json:"EndDate"`
}

func (UserProfileModel) TableName() string {
	return "user_profile"
}

// type UserProfileUnmarshalModel struct {
// 	Id            string                   `json:"id"`
// 	UserId        string                   `json:"UserId"`
// 	Title         string                   `json:"Title"`
// 	Position      string                   `json:"Position"`
// 	Description   string                   `json:"Description"`
// 	Company       string                   `json:"Company"`
// 	Salary        string                   `json:"Salary"`
// 	Type          string                   `json:"Type"`
// 	Keywords      []PreferenceKeywordModel `json:"Keywords"`
// 	StartDate     string                   `json:"StartDate"`
// 	EndDate       string                   `json:"EndDate"`
// 	CompanyDetail CompanyDetailModel       `json:"CompanyDetail"`
// }

func (m *UserProfileModel) ToProto() *protos.UserProfile {
	mType := protos.UserProfileType_value[m.Type]

	return &protos.UserProfile{
		UserId:      fmt.Sprintf("%d", m.User.ID),
		Title:       m.Title,
		Position:    m.Position,
		Description: m.Description,
		Salary:      m.Salary,
		Type:        protos.UserProfileType(mType),
		Keywords:    lo.Map(m.Keywords, func(k PreferenceKeywordModel, _ int) *protos.PreferenceKeyword { return k.ToProto() }),
		StartDate:   m.StartDate,
		EndDate:     m.EndDate,
	}
}

func (m *UserProfileModel) FromProto(p *protos.UserProfile) {
	idv, _ := strconv.Atoi(p.UserId)
	m.User = UserAccountModel{}
	m.User.ID = uint(idv)
	m.Title = p.Title
	m.Position = p.Position
	m.Description = p.Description
	m.Salary = p.Salary
	m.Type = p.Type.String()
	m.Keywords = lo.Map(p.Keywords, func(k *protos.PreferenceKeyword, _ int) PreferenceKeywordModel {
		pk := PreferenceKeywordModel{}
		pk.FromProto(k)
		return pk
	})
	m.StartDate = p.StartDate
	m.EndDate = p.EndDate
}

func (m *UserProfileModel) GetModel(db *gorm.DB) (*protos.UserProfile, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if err := db.First(m).Error; err != nil {
		return nil, db.Error
	}
	return m.ToProto(), nil

}

func (m *UserProfileModel) GetModelByUserId(db *gorm.DB) ([]*protos.UserProfile, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	var result []*UserProfileModel
	if err := db.Model(&UserProfileModel{}).Where("user_id = ?", m.User.ID).Find(&result).Error; err != nil {
		return nil, err
	}
	return lo.Map(result, func(item *UserProfileModel, index int) *protos.UserProfile {
		return item.ToProto()
	}), nil
}

func (m *UserProfileModel) CreateModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.User.ID == 0 {
		return fmt.Errorf("user id is empty")
	}

	if err := sd.Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserProfileModel) UpdateModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.ID == 0 {
		return fmt.Errorf("ID is empty")
	}
	if err := sd.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func (m *UserProfileModel) DeleteModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.ID == 0 {
		return fmt.Errorf("ID is empty")
	}
	if err := sd.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

func (UserProfileModel) DefineModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}

	return sd.AutoMigrate(&UserProfileModel{})

}
