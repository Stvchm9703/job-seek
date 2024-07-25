// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"fmt"
	"job-seek/pkg/protos"

	"github.com/samber/lo"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type UserCVProfileModel struct {
	Id         string   `json:"id"`
	UserId     string   `json:"user_id"`
	CvData     []byte   `json:"cv_data"`
	CvKeywords []string `json:"cv_keywords"`
}

func (m *UserCVProfileModel) ToProto() protos.UserCVProfile {
	return protos.UserCVProfile{
		UserId:     m.UserId,
		CvId:       m.Id,
		CvData:     m.CvData,
		CvKeywords: []*protos.PreferenceKeyword{},
	}
}

func (m *UserCVProfileModel) FromProto(p *protos.UserCVProfile) {
	m.UserId = p.UserId
	m.Id = p.CvId
	m.CvData = p.CvData
	m.CvKeywords = lo.Map(p.CvKeywords, func(x *protos.PreferenceKeyword, _ int) string { return x.KwId })
}

func (m *UserCVProfileModel) GetModel(db *surrealdb.DB) (*protos.UserCVProfile, error) {
	result, err := db.Query(fmt.Sprintf(`
	SELECT *, (SELECT * FROM PreferenceKeyword WHERE id INSIDE $parent.CvKeywords) AS CvKeywords
	FROM UserCVProfile:%s;`, m.Id), nil)

	if err != nil {
		return nil, err
	}

	var data *protos.UserCVProfile
	err = surrealdb.Unmarshal(result, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (m *UserCVProfileModel) CreateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	result, err := sd.Create("UserCVProfile", map[string]interface{}{
		"UserId":     m.UserId,
		"CvData":     m.CvData,
		"CvKeywords": m.CvKeywords,
	})
	if err != nil {
		return err
	}
	var data *UserCVProfileModel
	err = surrealdb.Unmarshal(result, data)
	if err != nil {
		return err
	}
	m.Id = data.Id
	return nil
}

func (m *UserCVProfileModel) UpdateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	_, err := sd.Update(fmt.Sprintf("UserCVProfile:%s", m.Id), m)
	return err
}

func (m *UserCVProfileModel) DefineModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	query := `
DEFINE TABLE IF NOT EXISTS UserCVProfile SCHEMAFULL;
-- Field definition
	DEFINE FIELD IF NOT EXISTS	UserId 					ON TABLE UserAccount TYPE		string;
	DEFINE FIELD IF NOT EXISTS	CvData 					ON TABLE UserAccount TYPE		bytes;
	DEFINE FIELD IF NOT EXISTS	CvKeywords			ON TABLE UserAccount TYPE		array<record<PreferenceKeyword>>;
`
	_, err := sd.Query(query, nil)
	return err
}
