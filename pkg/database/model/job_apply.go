// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"fmt"
	"job-seek/pkg/protos"
	"strconv"
	"time"

	// surrealdb "github.com/surrealdb/surrealdb.go"
	"gorm.io/gorm"
)

type JobApplyModel struct {
	gorm.Model
	Id          string           `json:"id,omitempty" gorm:"primaryKey"`
	JobID       int              `json:"-"`
	Job         JobModel         `gorm:"foreignKey:JobID"`
	UserID      int              `json:"-"`
	User        UserAccountModel `gorm:"foreignKey:UserID"`
	Status      string           `json:"status,omitempty"`
	CoverLetter string           `json:"cover_letter,omitempty"`
	CvContent   string           `json:"cv_content,omitempty"`
	CvFile      []byte           `json:"cv_file,omitempty"`
	Message     string           `json:"message,omitempty"`
}

func (JobApplyModel) TableName() string {
	return "job_apply"
}

func (m *JobApplyModel) ToProto() *protos.JobApply {
	var status = protos.JobStatus(protos.JobStatus_value[m.Status])

	createdAt := m.CreatedAt.Format(time.RFC3339)
	updatedAt := m.UpdatedAt.Format(time.RFC3339)
	deleteAt := ""
	// deletedAtValue, _ := m.DeletedAt.Value()
	// if deletedAtValue != nil {
	// 	deleteAt = deletedAtValue.Format(time.RFC3339)
	// }

	return &protos.JobApply{
		JobId:       m.Job.PostId,
		UserId:      fmt.Sprintf("%d", m.User.ID),
		Status:      &status,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
		DeletedAt:   &deleteAt,
		CoverLetter: &m.CoverLetter,
		CvContent:   &m.CvContent,
		CvFile:      m.CvFile,
		Job:         m.Job.ToProto(),
		Message:     &m.Message,
	}
}

func (m *JobApplyModel) FromProto(p *protos.JobApply) {
	// m.JobId = p.Job.PostId
	m.Job = JobModel{PostId: p.GetJobId()}
	// m.UserId = p.User.ID
	uidv, _ := strconv.Atoi(p.GetUserId())
	// m.User = UserAccountModel{ID: p.GetUserId()}
	m.User.ID = uint(uidv)
	m.Status = p.GetStatus().String()
	m.CoverLetter = p.GetCoverLetter()
	m.CvContent = p.GetCvContent()
	m.CvFile = p.GetCvFile()
	m.Message = p.GetMessage()
}

// type JobApplyUnmarshalModel struct {
// 	Id          string             `json:"id,omitempty"`
// 	JobId       string             `json:"job_id"`
// 	UserId      string             `json:"user_id"`
// 	Status      string             `json:"status,omitempty"`
// 	CreatedAt   string             `json:"created_at"`
// 	UpdatedAt   string             `json:"updated_at,omitempty"`
// 	DeletedAt   string             `json:"deleted_at,omitempty"`
// 	CoverLetter string             `json:"cover_letter,omitempty"`
// 	CvContent   string             `json:"cv_content,omitempty"`
// 	CvFile      []byte             `json:"cv_file,omitempty"`
// 	Message     string             `json:"message,omitempty"`
// 	Job         *JobUnmarshalModel `json:"job,omitempty"`
// }

// func (m *JobApplyUnmarshalModel) ToProto() *protos.JobApply {
// 	var jobProto *protos.Job = nil
// 	if m.Job != nil {
// 		jobProto = m.Job.ToProto()
// 	}
// 	status := protos.JobStatus(protos.JobStatus_value[m.Status])

// 	return &protos.JobApply{
// 		JobId:       m.JobId,
// 		UserId:      m.UserId,
// 		CreatedAt:   &m.CreatedAt,
// 		UpdatedAt:   &m.UpdatedAt,
// 		DeletedAt:   &m.DeletedAt,
// 		Job:         jobProto,
// 		Status:      &status,
// 		CoverLetter: &m.CoverLetter,
// 		CvContent:   &m.CvContent,
// 		CvFile:      m.CvFile,
// 		Message:     &m.Message,
// 	}
// }

func (m *JobApplyModel) GetModel(db *gorm.DB) (*protos.JobApply, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	// var res JobApplyModel
	result := db.First(m)
	if result.Error != nil {
		return nil, result.Error
	}

	return m.ToProto(), nil
}

// func (m *JobApplyModel) GetModelList(db *gorm.DB) ([]*protos.JobApply, error) {
// 	query := fmt.Sprintf(`
// 	SELECT *, (SELECT * FROM Job WHERE PostId = $parent.JobId) AS Job
// 	FROM JobApply WHERE UserId = $1
// 	ORDER BY CreatedAt DESC;
// 	`, m.UserId)

// 	result, err := db.Query(query, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var queryResult []QueryResult[JobApplyUnmarshalModel]
// 	// err = surrealdb.Unmarshal(result, jobqueryResult)
// 	jsonResult, _ := json.Marshal(result)
// 	err = json.Unmarshal(jsonResult, &queryResult)
// 	if err != nil {
// 		errorWrap := errors.Join(err, fmt.Errorf("query: %s", query), fmt.Errorf("raw: %s", jsonResult))
// 		return nil, errorWrap
// 		// return nil, err
// 	}
// 	// pp.Println("jobs:", jobqueryResult)
// 	if len(queryResult) == 0 || len(queryResult[0].Result) == 0 {
// 		return nil, fmt.Errorf("no data found")
// 	}
// 	// pp.Println("jobs:", jobqueryResult[0].Result[0])
// 	protoResult := lo.Map(queryResult[0].Result, func(item JobApplyUnmarshalModel, _ int) *protos.JobApply {
// 		return item.ToProto()
// 	})
// 	return protoResult, nil
// }

// func (m *JobApplyModel) CreateModel(sd *gorm.DB) error {
// 	if sd == nil {
// 		return fmt.Errorf("database connection is nil")
// 	}

// 	queryTemplate, _ := template.New("createJobApply").Parse(`
// INSERT INTO JobApply {
// 	JobId       : r"{{.JobId}}",
// 	UserId      : r"{{.UserId}}",
// 	Status      : s"{{.Status}}",
// 	CreatedAt   : s"",
// 	UpdatedAt   : s"",
// 	DeletedAt   : s"",
// 	CoverLetter : s"{{.CoverLetter}}",
// 	CvContent   : s"{{.CvContent}}",
// 	CvFile      : b"{{.CvFile}}",
// 	Message 		: s"{{.Message}}",
// }	`)
// 	var doc bytes.Buffer
// 	var err error
// 	err = queryTemplate.Execute(&doc, m)
// 	if err != nil {
// 		return err
// 	}
// 	// _, err := sd.Create(fmt.Sprintf("CompanyDetail:%s", m.ReferenceId), m)
// 	query := strings.ReplaceAll(doc.String(), "\n", " ")
// 	query = strings.ReplaceAll(query, "\t", " ")
// 	query = strings.ReplaceAll(query, "\r", " ")
// 	query = strings.ReplaceAll(query, `\"`, `"`)
// 	query = strings.Join(strings.Fields(strings.TrimSpace(query)), " ")

// 	result, err := sd.Query(query, m)
// 	var message map[string]interface{}
// 	surrealdb.Unmarshal(result, message)
// 	if err != nil {
// 		fmt.Println("query:", query)
// 		pp.Println("message:", message)
// 		return errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("message: %v", message))
// 	}
// 	return nil
// }

// func (m *JobApplyModel) UpdateModel(sd *gorm.DB) error {
// 	if sd == nil {
// 		return fmt.Errorf("database connection is nil")
// 	}
// 	_, err := sd.Update(fmt.Sprintf("JobApply:%s", m.Id), m)
// 	return err
// }

// func (m JobApplyModel) DefineModel(sd *gorm.DB) error {
// 	if sd == nil {
// 		return fmt.Errorf("database connection is nil")
// 	}

// 	query := `
// -- Table definition
// DEFINE TABLE IF NOT EXISTS JobApply SCHEMAFULL;
// -- Field definition
//   DEFINE FIELD IF NOT EXISTS JobId      ON TABLE JobApply TYPE  record<Job>;
//   DEFINE FIELD IF NOT EXISTS UserId     ON TABLE JobApply TYPE  record<UserAccount>;
//   DEFINE FIELD IF NOT EXISTS Status     ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS CreatedAt    ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS UpdatedAt    ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS DeletedAt    ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS CoverLetter   ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS CVContent    ON TABLE JobApply TYPE  string;
//   DEFINE FIELD IF NOT EXISTS CvFile     ON TABLE JobApply TYPE  bytes;
//   DEFINE FIELD IF NOT EXISTS Message     ON TABLE JobApply TYPE  string;
// -- Index definition
//   DEFINE INDEX IF NOT EXISTS id       ON TABLE JobApply COLUMNS ReferenceId UNIQUE;
// -- Event definition
//   DEFINE EVENT IF NOT EXISTS CreateHook ON TABLE JobBookmark
//     WHEN $event = "CREATE" OR $event = "INSERT"
//     THEN (
//       UPDATE JobBookmark SET CreatedAt = time::format(time::now(),"%+")
//         WHERE JobId = $after.JobId AND UserId = $after.UserId
//     );
//   DEFINE EVENT IF NOT EXISTS UpdateHook ON TABLE JobBookmark
//     WHEN $event = "CREATE" OR $event = "INSERT"
//     THEN (
//       UPDATE JobBookmark SET UpdatedAt = time::format(time::now(),"%+")
//         WHERE JobId = $after.JobId AND UserId = $after.UserId
//     );
// -- END OF table definition
// `
// 	_, err := sd.Query(query, nil)
// 	return err
// }
