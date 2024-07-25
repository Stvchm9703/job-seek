package model

import (
	// "job-seek/pkg/config"
	// "job-seek/pkg/database"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/pkg/request/seek_api"

	"github.com/samber/lo"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

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

func (m *JobCacheListModel) ToProto() protos.JobSearchResponse {
	pageNumber := int32(m.PageNumber)
	totalCount := int32(m.TotalCount)
	totalPage := int32(m.TotalPage)

	return protos.JobSearchResponse{
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
	m.Job = lo.Map(p.Data, func(x seek_api.SeekSearchApiResponseData, _ int) string { return fmt.Sprintf("%d", x.ID) })
}

func (m *JobCacheListModel) CreateJobCacheList(db *surrealdb.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return fmt.Errorf("cache ref is empty")
	}

	_, err := db.Create(
		fmt.Sprintf("JobCacheList[%s,%s,%d]", m.CacheRef, m.UserQueryId, m.PageNumber),
		m)
	if err != nil {
		return err
	}

	return nil
}

func (m *JobCacheListModel) GetJobCacheList(db *surrealdb.DB) (*protos.JobSearchResponse, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	result, err := db.Query(
		fmt.Sprintf(`
		SELECT *, 
			(SELECT * 
				(SELECT * FROM CompanyDetail
					WHERE ReferanceId = $parent.Job[*].CompanyDetail
				) AS CompanyDetail
				FROM Job 
				WHERE id INSIDE $parent.Job) AS Job
		FROM JobCacheList[%s,%s,%d];`, m.CacheRef, m.UserQueryId, m.PageNumber), nil)

	if err != nil {
		return nil, err
	}

	var data *protos.JobSearchResponse
	err = surrealdb.Unmarshal(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *JobCacheListModel) DefineModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := `
-- Table definition
DEFINE TABLE IF NOT EXISTS JobCacheList SCHEMAFULL;
-- Field definition
	DEFINE FIELD IF NOT EXISTS	CacheRef				ON TABLE JobCacheList TYPE		string;
	DEFINE FIELD IF NOT EXISTS	UserQueryId			ON TABLE JobCacheList TYPE		string;
	DEFINE FIELD IF NOT EXISTS	Job 						ON TABLE JobCacheList TYPE		array<record<Job>>;
	DEFINE FIELD IF NOT EXISTS	PageNumber 			ON TABLE JobCacheList TYPE		number;
	DEFINE FIELD IF NOT EXISTS	TotalCount 			ON TABLE JobCacheList TYPE		number;
	DEFINE FIELD IF NOT EXISTS	TotalPage 			ON TABLE JobCacheList TYPE		number;
	DEFINE FIELD IF NOT EXISTS	CurrentKeyword 	ON TABLE JobCacheList TYPE		string;
	DEFINE FIELD IF NOT EXISTS	SearchParams 		ON TABLE JobCacheList FLEXIBLE TYPE		object;
	DEFINE FIELD IF NOT EXISTS	ExpiredDate 		ON TABLE JobCacheList TYPE		string;
-- Index definition
	DEFINE INDEX IF NOT EXISTS	id							ON TABLE Job COLUMNS [CacheRef,UserQueryId,PageNumber];
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
