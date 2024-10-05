// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"job-seek/pkg/protos"
	"log"
	"strings"
	"text/template"

	"github.com/k0kubun/pp/v3"
	"github.com/samber/lo"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type JobBookmarkModel struct {
	JobId     string `json:"job_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	// Job       string `json:"job,omitempty"`
}

func (m *JobBookmarkModel) ToProto() protos.JobBookmark {
	return protos.JobBookmark{
		JobId:     m.JobId,
		UserId:    m.UserId,
		CreatedAt: &m.CreatedAt,
		UpdatedAt: &m.UpdatedAt,
		DeletedAt: &m.DeletedAt,
		Job:       nil,
	}
}

func (m *JobBookmarkModel) FromProto(p *protos.JobBookmark) {
	m.JobId = p.JobId
	m.UserId = p.UserId
	m.CreatedAt = p.GetCreatedAt()
	m.UpdatedAt = p.GetUpdatedAt()
	m.DeletedAt = p.GetDeletedAt()
	// m.Job = p.GetJob().PostId
}

type JobBookmarkUnmarshalModel struct {
	Id        string             `json:"id"`
	JobId     string             `json:"job_id"`
	UserId    string             `json:"user_id"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at,omitempty"`
	DeletedAt string             `json:"deleted_at,omitempty"`
	Job       *JobUnmarshalModel `json:"job,omitempty"`
}

func (m *JobBookmarkUnmarshalModel) ToProto() *protos.JobBookmark {
	var jobProto *protos.Job = nil
	if m.Job != nil {
		jobProto = m.Job.ToProto()
	}
	return &protos.JobBookmark{
		JobId:     m.JobId,
		UserId:    m.UserId,
		CreatedAt: &m.CreatedAt,
		UpdatedAt: &m.UpdatedAt,
		DeletedAt: &m.DeletedAt,
		Job:       jobProto,
	}
}

func (m *JobBookmarkModel) GetModel(db *surrealdb.DB) (*protos.JobBookmark, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf(`
	SELECT *, 
		(SELECT *, 
			(SELECT * FROM CompanyDetail  WHERE id = $parent.CompanyDetail Limit 1)[0] AS CompanyDetail 
			FROM Job 
			WHERE PostId = $parent.JobId
		)[0] AS Job
	FROM JobBookmark:[%s, %s];
	`, m.UserId, m.JobId)
	result, err := db.Query(query, nil)

	if err != nil {
		return nil, err
	}

	// resultMap := result.([]map[string]map[string][]interface{})

	var queryResult []QueryResult[JobBookmarkUnmarshalModel]
	// err = surrealdb.Unmarshal(result, jobqueryResult)
	jsonResult, _ := json.Marshal(result)
	err = json.Unmarshal(jsonResult, &queryResult)
	if err != nil {
		errorWrap := errors.Join(err, fmt.Errorf("query: %s", query), fmt.Errorf("raw: %s", jsonResult))
		log.Fatalf("error: %v", errorWrap)
		return nil, errorWrap
		// return nil, err
	}
	// pp.Println("jobs:", jobqueryResult)
	if len(queryResult) == 0 || len(queryResult[0].Result) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	// pp.Println("jobs:", jobqueryResult[0].Result[0])
	return queryResult[0].Result[0].ToProto(), nil
}

func (m *JobBookmarkModel) GetModelByUser(db *surrealdb.DB) ([]*protos.JobBookmark, error) {
	query := `
	SELECT 
		*, 
		(SELECT *, 
			(SELECT * FROM CompanyDetail  WHERE id = $parent.CompanyDetail Limit 1)[0] AS CompanyDetail 
			FROM Job 
			WHERE PostId = $parent.JobId
		)[0] AS Job
	FROM JobBookmark 
	WHERE UserId = $user_id
	;
	`

	result, err := db.Query(query, map[string]interface{}{
		"user_id": m.UserId,
	})

	if err != nil {
		return nil, err
	}
	var queryResult []QueryResult[JobBookmarkUnmarshalModel]
	// err = surrealdb.Unmarshal(result, jobqueryResult)
	jsonResult, _ := json.Marshal(result)
	err = json.Unmarshal(jsonResult, &queryResult)
	if err != nil {
		errorWrap := errors.Join(err, fmt.Errorf("query: %s", query), fmt.Errorf("raw: %s", jsonResult))
		log.Fatalf("error: %v", errorWrap)
		return nil, errorWrap
		// return nil, err
	}
	// pp.Println("jobs:", jobqueryResult)
	if len(queryResult) == 0 || len(queryResult[0].Result) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	// pp.Println("jobs:", jobqueryResult[0].Result[0])

	protoed := lo.Map(queryResult[0].Result, func(x JobBookmarkUnmarshalModel, _ int) *protos.JobBookmark {
		return x.ToProto()
	})
	return protoed, nil
}

func (m *JobBookmarkModel) CreateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}

	queryTemplate, _ := template.New("createJobBookmark").Parse(`
INSERT INTO JobBookmark {
JobId     : r"{{.JobId}}",
UserId    : r"{{.UserId}}",
CreatedAt : s"",
UpdatedAt : s"",
DeletedAt : s"",
}	`)
	var doc bytes.Buffer
	var err error
	err = queryTemplate.Execute(&doc, m)
	if err != nil {
		return err
	}
	// _, err := sd.Create(fmt.Sprintf("CompanyDetail:%s", m.ReferenceId), m)
	query := strings.ReplaceAll(doc.String(), "\n", " ")
	query = strings.ReplaceAll(query, "\t", " ")
	query = strings.ReplaceAll(query, "\r", " ")
	query = strings.ReplaceAll(query, `\"`, `"`)
	query = strings.Join(strings.Fields(strings.TrimSpace(query)), " ")

	result, err := sd.Query(query, m)
	var message map[string]interface{}
	surrealdb.Unmarshal(result, message)
	if err != nil {
		fmt.Println("query:", query)
		pp.Println("message:", message)
		return errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("message: %v", message))
	}
	return nil
}

func (m *JobBookmarkModel) UpdateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	_, err := sd.Update(fmt.Sprintf("JobBookmark:[%s,%s]", m.UserId, m.JobId), m)
	return err
}

func (m JobBookmarkModel) DefineModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := `
-- Table definition
DEFINE TABLE IF NOT EXISTS JobBookmark SCHEMAFULL;
-- Field definition
	DEFINE FIELD IF NOT EXISTS	JobId 					ON TABLE JobBookmark TYPE		record<Job>;
	DEFINE FIELD IF NOT EXISTS	UserId 					ON TABLE JobBookmark TYPE		record<UserAccount>;
	DEFINE FIELD IF NOT EXISTS	CreatedAt 			ON TABLE JobBookmark TYPE		string;
	DEFINE FIELD IF NOT EXISTS	UpdatedAt 			ON TABLE JobBookmark TYPE		string;
	DEFINE FIELD IF NOT EXISTS	DeletedAt 			ON TABLE JobBookmark TYPE		string;
-- Index definition
	DEFINE INDEX IF NOT EXISTS	id							ON TABLE JobBookmark COLUMNS UserId,JobId;
-- Event definition
	DEFINE EVENT IF NOT EXISTS CreateHook ON TABLE JobBookmark 
		WHEN $event = "CREATE" OR $event = "INSERT"
		THEN (
			UPDATE JobBookmark SET CreatedAt = time::format(time::now(),"%+") 
				WHERE JobId = $after.JobId AND UserId = $after.UserId
		);
	DEFINE EVENT IF NOT EXISTS UpdateHook ON TABLE JobBookmark 
		WHEN $event = "CREATE" OR $event = "INSERT"
		THEN (
			UPDATE JobBookmark SET UpdatedAt = time::format(time::now(),"%+") 
				WHERE JobId = $after.JobId AND UserId = $after.UserId
		);
-- END OF table definition
		`
	_, err := sd.Query(query, nil)
	return err

}