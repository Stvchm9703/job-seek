// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"fmt"
	"job-seek/pkg/protos"

	surrealdb "github.com/surrealdb/surrealdb.go"
)

type BookmarkJobModel struct {
	JobId     string `json:"job_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Job       string `json:"job,omitempty"`
}

func (m *BookmarkJobModel) ToProto() protos.BookmarkJob {
	return protos.BookmarkJob{
		JobId:     m.JobId,
		UserId:    m.UserId,
		CreatedAt: &m.CreatedAt,
		UpdatedAt: &m.UpdatedAt,
		DeletedAt: &m.DeletedAt,
		Job:       nil,
	}
}

func (m *BookmarkJobModel) FromProto(p *protos.BookmarkJob) {
	m.JobId = p.JobId
	m.UserId = p.UserId
	m.CreatedAt = p.GetCreatedAt()
	m.UpdatedAt = p.GetUpdatedAt()
	m.DeletedAt = p.GetDeletedAt()
	m.Job = p.GetJob().PostId
}

func (m *BookmarkJobModel) GetModel(db *surrealdb.DB) (*protos.BookmarkJob, error) {
	result, err := db.Query(
		`
	SELECT *, (SELECT * FROM Job WHERE PostId = $parent.JobId) AS Job
	FROM BookmarkJob:[$user_id,$job_id];
	`, map[string]interface{}{
			"job_id":  m.JobId,
			"user_id": m.UserId,
		})

	if err != nil {
		return nil, err
	}

	var data *protos.BookmarkJob
	err = surrealdb.Unmarshal(result, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (m *BookmarkJobModel) CreateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	_, err := sd.Create(
		fmt.Sprintf("BookmarkJob:[%s,%s]", m.UserId, m.JobId),
		m)
	return err
}

func (m *BookmarkJobModel) UpdateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	_, err := sd.Update(fmt.Sprintf("BookmarkJob:[%s,%s]", m.UserId, m.JobId), m)
	return err
}
