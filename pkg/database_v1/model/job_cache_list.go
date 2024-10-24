package model

import (
	// "job-seek/pkg/config"
	// "job-seek/pkg/database"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/pkg/request/seek_api"
	"log"
	"strings"
	"text/template"

	"github.com/k0kubun/pp/v3"
	"github.com/samber/lo"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type JobCacheListUnmarshalModel struct {
	CacheRef       string
	UserQueryId    string
	Job            []*JobUnmarshalModel
	PageNumber     int
	TotalCount     int
	TotalPage      int
	CurrentKeyword string
	SearchParams   seek_api.SeekSearchApiParams
	ExpiredDate    string
}

func (m *JobCacheListUnmarshalModel) ToProto() *protos.JobSearchResponse {
	pageNumber := int32(m.PageNumber)
	totalCount := int32(m.TotalCount)
	totalPage := int32(m.TotalPage)

	return &protos.JobSearchResponse{
		CacheRef: &m.CacheRef,
		// UserQueryId: &m.UserQueryId,
		PageNumber: &pageNumber,
		TotalCount: &totalCount,
		TotalPage:  &totalPage,
		Message:    &m.CurrentKeyword,
		Job:        lo.Map(m.Job, func(x *JobUnmarshalModel, _ int) *protos.Job { return x.ToProto() }),
	}
}

type JobCacheListModel struct {
	CacheRef       string
	UserQueryId    string
	Job            []string
	PageNumber     int
	TotalCount     int
	TotalPage      int
	CurrentKeyword string
	SearchParams   seek_api.SeekSearchApiParams
	ExpiredDate    string
}

func (m *JobCacheListModel) ToProto() *protos.JobSearchResponse {
	pageNumber := int32(m.PageNumber)
	totalCount := int32(m.TotalCount)
	totalPage := int32(m.TotalPage)

	return &protos.JobSearchResponse{
		CacheRef: &m.CacheRef,
		// UserQueryId: &m.UserQueryId,
		PageNumber: &pageNumber,
		TotalCount: &totalCount,
		TotalPage:  &totalPage,
		Message:    &m.CurrentKeyword,
		Job:        []*protos.Job{},
	}
}

func (m *JobCacheListModel) FromSearchResult(r *seek_api.SeekSearchApiParams, p *seek_api.SeekSearchApiResponse) {
	m.SearchParams = *r
	if p.UserQueryID != "" {
		m.UserQueryId = p.UserQueryID
	} else if p.SearchParams.UserQueryId != "" {
		m.UserQueryId = p.SearchParams.UserQueryId
	}
	m.PageNumber = r.Page
	m.TotalCount = p.TotalCount
	m.TotalPage = p.TotalPages
	m.CurrentKeyword = r.Keywords
	m.Job = lo.Map(p.Data, func(x seek_api.SeekSearchApiResponseData, _ int) string { return fmt.Sprintf("Job:%d", x.ID) })
}

func (m *JobCacheListModel) CreateJobCacheList(db *surrealdb.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return fmt.Errorf("cache ref is empty")
	}
	queryTemplate, _ := template.New("createJob").Parse(`
CREATE JobCacheList:[s"{{.CacheRef}}","{{.UserQueryId}}",{{.PageNumber}}] CONTENT {
	CacheRef       : s"{{.CacheRef}}",
	UserQueryId    : s"{{.UserQueryId}}",
	Job         :   [{{- range $i, $v := .Job}}{{- if $i}},{{- end}}r"{{- $v}}"{{- end}}],
	PageNumber : {{- .PageNumber}},
	TotalCount : {{- .TotalCount}},
	TotalPage : {{- .TotalPage}},
	CurrentKeyword : s"{{- .CurrentKeyword}}",
	SearchParams : {
		SiteKey : s"{{- .SearchParams.SiteKey}}",
		Where : s"{{- .SearchParams.Where}}",
		Page : {{- .SearchParams.Page}},
		PageSize : {{- .SearchParams.PageSize}},
		Keywords : s"{{- .SearchParams.Keywords}}",
		SalaryType : s"{{- .SearchParams.SalaryType}}",
		SalaryRange : s"{{- .SearchParams.SalaryRange}}",
		Locale : s"{{- .SearchParams.Locale}}",
		SeekerId : s"{{- .SearchParams.SeekerId}}",
		Classification : s"{{- .SearchParams.Classification}}",
		AdvertiserId : s"{{- .SearchParams.AdvertiserId}}",
		UserQueryId : s"{{- .SearchParams.UserQueryId}}"
	},
	ExpiredDate : s"{{.ExpiredDate}}",
}
`)
	var doc bytes.Buffer
	var err error
	err = queryTemplate.Execute(&doc, m)
	if err != nil {
		return err
	}
	query := strings.ReplaceAll(doc.String(), "\n", " ")
	query = strings.ReplaceAll(query, "\t", " ")
	query = strings.ReplaceAll(query, "\r", " ")
	query = strings.ReplaceAll(query, `\"`, `"`)
	query = strings.TrimSpace(query)
	result, err := db.Query(query, m)
	var message map[string]interface{}
	surrealdb.Unmarshal(result, message)
	return errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("message:", message))
}

func (m *JobCacheListModel) GetJobCacheListWithQueryId(db *surrealdb.DB) (*protos.JobSearchResponse, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return nil, fmt.Errorf("cache ref is empty")
	}
	query := fmt.Sprintf(`
		SELECT *, 
			(SELECT *,
				(SELECT * FROM CompanyDetail
					WHERE ReferanceId = $parent.Job[*].CompanyDetail
					Limit 1
				)[0] AS CompanyDetail
			FROM Job 
			WHERE id INSIDE $parent.Job) AS Job
		FROM JobCacheList:[%s,%s,%d];`, m.CacheRef, m.UserQueryId, m.PageNumber)
	result, err := db.Query(query, nil)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("raw", result))
	}
	// pp.Printf("result: %v", result)

	data := make([]JobCacheListUnmarshalModel, 1)
	err = surrealdb.Unmarshal(result, data)

	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("parsing data error"), err,
			fmt.Errorf("query: %s", query), pp.Errorf("raw", result))
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return data[0].ToProto(), nil
}

func (m *JobCacheListModel) GetJobCacheList(db *surrealdb.DB) (*protos.JobSearchResponse, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return nil, fmt.Errorf("cache ref is empty")
	}
	query := fmt.Sprintf(`
		SELECT *, 
			(SELECT *,
				(SELECT * FROM CompanyDetail
					WHERE ReferanceId = $parent.Job[*].CompanyDetail
					Limit 1
				)[0] AS CompanyDetail
			FROM Job 
			WHERE id INSIDE $parent.Job) AS Job
		FROM JobCacheList
		WHERE CacheRef = s'%s'
		;`, m.CacheRef)
	result, err := db.Query(query, nil)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("raw", result))
	}
	// pp.Printf("result: %v", result)

	var queryResult []QueryResult[JobCacheListUnmarshalModel]
	// err = surrealdb.Unmarshal(result, jobqueryResult)
	jsonResult, _ := json.Marshal(result)
	err = json.Unmarshal(jsonResult, &queryResult)
	if err != nil {
		errorWrap := errors.Join(err, fmt.Errorf("query: %s", query), fmt.Errorf("raw: %s", jsonResult))
		log.Fatalf("error: %v", errorWrap)
		return nil, errorWrap
		// return nil, err
	}

	if len(queryResult) == 0 || len(queryResult[0].Result) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return queryResult[0].Result[0].ToProto(), nil
}

func (m *JobCacheListModel) DefineModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := `
-- Table definition
DEFINE TABLE IF NOT EXISTS JobCacheList SCHEMAFULL;
-- Field definition
  DEFINE FIELD IF NOT EXISTS  CacheRef        ON TABLE JobCacheList TYPE    string;
  DEFINE FIELD IF NOT EXISTS  UserQueryId      ON TABLE JobCacheList TYPE    string;
  DEFINE FIELD IF NOT EXISTS  Job             ON TABLE JobCacheList TYPE    array<record<Job>>;
  DEFINE FIELD IF NOT EXISTS  PageNumber       ON TABLE JobCacheList TYPE    number;
  DEFINE FIELD IF NOT EXISTS  TotalCount       ON TABLE JobCacheList TYPE    number;
  DEFINE FIELD IF NOT EXISTS  TotalPage       ON TABLE JobCacheList TYPE    number;
  DEFINE FIELD IF NOT EXISTS  CurrentKeyword   ON TABLE JobCacheList TYPE    string;
  DEFINE FIELD IF NOT EXISTS  SearchParams     ON TABLE JobCacheList FLEXIBLE TYPE    object;
  DEFINE FIELD IF NOT EXISTS  ExpiredDate     ON TABLE JobCacheList TYPE    string;
-- Event definition
  DEFINE EVENT IF NOT EXISTS CreateHook ON TABLE JobCacheList 
    WHEN $event = "CREATE" OR $event = "INSERT"
    THEN (
      UPDATE JobCacheList SET ExpiredDate = time::format(time::now()+duration::from::days(5),"%+")
        WHERE CacheRef = $after.CacheRef AND UserQueryId = $after.UserQueryId AND PageNumber = $after.PageNumber
    );  
-- END OF table definition
    `
	_, err := sd.Query(query, nil)
	return err

}
