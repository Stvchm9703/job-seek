package model

import (
	// "job-seek/pkg/config"
	// "job-seek/pkg/database"

	"fmt"
	"job-seek/pkg/protos"
	"job-seek/pkg/request/seek_api"

	"github.com/k0kubun/pp/v3"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type JobCacheListModel struct {
	gorm.Model
	CacheRef       string
	UserQueryId    string
	Job            []JobModel `gorm:"foreignKey:PostId"`
	PageNumber     int
	TotalCount     int
	TotalPage      int
	CurrentKeyword string
	SearchParams   seek_api.SeekSearchApiParams `gorm:"foreignKey:ID"`
	ExpiredDate    string
}

func (JobCacheListModel) TableName() string {
	return "job_cache_list"
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
		Job:        lo.Map(m.Job, func(x JobModel, _ int) *protos.Job { return x.ToProto() }),
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
	m.Job = lo.Map(p.Data, func(x seek_api.SeekSearchApiResponseData, _ int) JobModel {
		return JobModel{PostId: fmt.Sprintf("%d", x.ID)}
	})
}

func (m *JobCacheListModel) CreateJobCacheList(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return fmt.Errorf("cache ref is empty")
	}
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *JobCacheListModel) GetJobCacheList(db *gorm.DB) (*protos.JobSearchResponse, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if m.CacheRef == "" {
		return nil, fmt.Errorf("cache ref is empty")
	}

	if err := db.Model(&JobCacheListModel{}).Preload("Job").First(m).Error; err != nil {
		return nil, err
	}
	pp.Println(m)
	return m.ToProto(), nil
}

func (m *JobCacheListModel) DefineModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	return sd.AutoMigrate(&JobCacheListModel{})

}
