// source: JobSearchService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/database/model"
	"job-seek/pkg/protos"
	"job-seek/pkg/request"
	linkedin "job-seek/pkg/request/linkedin_search"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCompanyDetail implements GetCompanyDetail from JobSearchServiceServer
// generated by protoc-gen-go-grpc.JobSearchServiceServer
func (s JobSearchServiceServerImpl) GetCompanyDetail(ctx context.Context, req *protos.CompanyDetailRequest) (*protos.CompanyDetail, error) {
	// todo()
	var companyDetail *protos.CompanyDetail
	var err error
	companyId := req.GetReferenceId()

	if companyId != "" {
		companyDetail, err = s.getCompanyDetailFromDB(companyId)
		if err != nil {
		}
		return companyDetail, nil
	}

	companyName := req.GetName()
	if companyName != "" {
		companyDetail, err = s.getCompanyDetailFromAPI(companyName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to get company detail from API")
		}

		var jobPostCount int
		jobPostCount, err = linkedin.GetCompanyPostListForApi(&s.config.SeekService.ApiService, companyId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to get company post list from API")
		}

		companyDetail.JobPosted = int32(jobPostCount)

		return companyDetail, nil
	}

	return nil, status.Errorf(codes.Unimplemented, "method  GetCompanyDetail  not implemented")
}

func (s JobSearchServiceServerImpl) getCompanyDetailFromDB(company_id string) (*protos.CompanyDetail, error) {
	companyModel := model.CompanyDetailModel{
		ReferenceId: company_id,
	}

	data, err := companyModel.GetModel(s.dbClient)
	if err != nil {
		s.log.WithFields(map[string]interface{}{
			"company_id": company_id,
			"error":      err,
			"method":     "getCompanyDetailFromDB",
		}).Error("Failed to get company detail from DB")
		return nil, err
	}
	return data, nil
}

func (s JobSearchServiceServerImpl) getCompanyDetailFromAPI(company_name string) (*protos.CompanyDetail, error) {

	companyDetail := request.SeekCompanyDetails{
		Name: company_name,
	}

	collector := linkedin.CreateSearchEngineCollector()
	_, err := linkedin.SearchCompanyForApi(collector, &s.config.YahooSearch, &companyDetail)
	if err != nil {
		s.log.WithFields(map[string]interface{}{
			"company_name": company_name,
			"config":       s.config.YahooSearch,
			"error":        err,
			"method":       "getCompanyDetailFromAPI",
		}).Error("Failed to search company from API")
		return nil, err
	}
	return companyDetail.ToProto(), nil
}

func (s JobSearchServiceServerImpl) storeCompanyDetailToDB(companyDetail *protos.CompanyDetail) error {
	companyModel := model.CompanyDetailModel{}
	companyModel.FromProto(companyDetail)

	err := companyModel.CreateModel(s.dbClient)
	if err != nil {
		s.log.WithFields(map[string]interface{}{
			"company_detail": companyDetail,
			"error":          err,
			"method":         "storeCompanyDetailToDB",
		}).Error("Failed to save company detail to DB")
		return err
	}
	return nil
}
