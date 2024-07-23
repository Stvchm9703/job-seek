package request

import (
	protos "job-seek/pkg/protos"
)

type SeekPostDetails struct {
	PostId         string
	PostTitle      string
	PostUrl        string
	PayRange       string
	DebugText      string
	HittedKeywords []string
	Score          int
	Role           string
	WorkType       string
	CompanyDetails *SeekCompanyDetails
	Locations      string
	ExpiringDate   string
}

func (p *SeekPostDetails) ToProto() *protos.Job {
	score := int32(p.Score)

	return &protos.Job{
		PostId:         p.PostId,
		PostTitle:      p.PostTitle,
		PostUrl:        p.PostUrl,
		PayRange:       p.PayRange,
		DebugText:      p.DebugText,
		HittedKeywords: p.HittedKeywords,
		Score:          &score,
		Role:           p.Role,
		WorkType:       p.WorkType,
		CompanyDetails: p.CompanyDetails.ToProto(),
		Locations:      p.Locations,
		ExpiringDate:   p.ExpiringDate,
	}

}

type SeekCompanyDetails struct {
	ReferenceId string
	Name        string
	Url         string
	Linkedin    string
	// SNS         string
	Description  string
	Industry     string
	JobPosted    int
	GroupSize    string
	HeadQuarters string
	Specialties  []string
	Locations    string
	// ContactPerson string
	// ContactEmail  string
}

func (c *SeekCompanyDetails) ToProto() *protos.CompanyDetail {

	groupSize := protos.CompanySize_SIZE_A
	switch c.GroupSize {
	case "1-10 employees":
		groupSize = protos.CompanySize_SIZE_A
	case "11-50 employees":
		groupSize = protos.CompanySize_SIZE_B
	case "51-200 employees":
		groupSize = protos.CompanySize_SIZE_C
	case "201-500 employees":
		groupSize = protos.CompanySize_SIZE_D
	case "501-1000 employees":
		groupSize = protos.CompanySize_SIZE_E
	case "1001-5000 employees":
		groupSize = protos.CompanySize_SIZE_F
	case "5001-10000 employees":
		groupSize = protos.CompanySize_SIZE_G
	case "10001+ employees":
		groupSize = protos.CompanySize_SIZE_H
	}

	return &protos.CompanyDetail{
		ReferenceId:  c.ReferenceId,
		Name:         c.Name,
		Url:          c.Url,
		Linkedin:     c.Linkedin,
		Description:  c.Description,
		Industry:     c.Industry,
		JobPosted:    int32(c.JobPosted),
		GroupSize:    groupSize,
		HeadQuarters: c.HeadQuarters,
		Specialties:  c.Specialties,
		Locations:    c.Locations,
	}
}
