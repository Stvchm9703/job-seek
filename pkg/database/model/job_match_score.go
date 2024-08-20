// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"job-seek/pkg/protos"

	"github.com/samber/lo"
)

type JobMatchScoreModel struct {
	UserId         string   `json:"user_id"`
	JobId          string   `json:"job_id"`
	PredictScore   float32  `json:"predict_score"`
	HittedKeywords []string `json:"hitted_keywords"`
	Job            string   `json:"job,omitempty"`
	UserProfile    string   `json:"user_profile,omitempty"`
}

type JobMatchScoreUnmarshalModel struct {
	Id             string                     `json:"id"`
	UserId         string                     `json:"user_id"`
	JobId          string                     `json:"job_id"`
	PredictScore   float32                    `json:"predict_score"`
	HittedKeywords []*PreferenceKeywordModel  `json:"hitted_keywords"`
	Job            *JobUnmarshalModel         `json:"job,omitempty"`
	UserProfile    *UserProfileUnmarshalModel `json:"user_profile,omitempty"`
}

func (m *JobMatchScoreUnmarshalModel) ToProto() *protos.JobMatchScore {
	return &protos.JobMatchScore{
		UserId:         m.UserId,
		JobId:          m.JobId,
		PredictScore:   m.PredictScore,
		HittedKeywords: lo.Map(m.HittedKeywords, func(p *PreferenceKeywordModel, _ int) *protos.PreferenceKeyword { return p.ToProto() }),
		Job:            m.Job.ToProto(),
		UserProfile:    m.UserProfile.ToProto(),
	}
}

func (m *JobMatchScoreModel) FromProto(p *protos.JobMatchScore) {
	m.UserId = p.UserId
	m.JobId = p.JobId
	m.PredictScore = p.PredictScore
	m.HittedKeywords = lo.Map(p.HittedKeywords, func(k *protos.PreferenceKeyword, _ int) string { return k.Keyword })
	m.Job = p.Job.PostId
	m.UserProfile = p.UserProfile.UserId
}