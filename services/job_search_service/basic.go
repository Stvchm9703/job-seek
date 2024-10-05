package job_search_service

import (
	"fmt"
	"job-seek/pkg/config"
	"job-seek/pkg/request"
	linkedin "job-seek/pkg/request/linkedin_search"
	"job-seek/pkg/request/seek_api"
	gql "job-seek/pkg/request/seek_gql"
	"job-seek/pkg/scoring"
	"job-seek/pkg/store"
	"job-seek/pkg/template"
	"log"

	ollamaApi "job-seek/pkg/request/generation_service/ollama"

	"time"

	pp "github.com/k0kubun/pp/v3"
)

var (
	SeekRequestChan  chan int
	SeekRequestCount int
	RateLimit        int = 36 // 36 request per minute
	// DebugComppanyProfile []interface{}
	// DebugComppanyTag     []interface{}
)

func main_process() {

	pp.Default.SetColorScheme(pp.ColorScheme{
		StructName: pp.NoColor,
	})

	readConfig := config.ReadSearchConfig()
	store.InitDB()
	defer store.CloseDB()

	pp.Println(readConfig)

	searchParamsPreset := seek_api.CreateSeekSearchApiParamsFromConfig(readConfig.SearchParamsPreset)

	// postData := []request.SeekSearchApiResponseData{}
	// postDetails := []request.SeekPostDetails{}
	combinedKeywords := seek_api.CreateSearchCombinations(*readConfig.SearchKeywords)

	postData := getPostPobsList(combinedKeywords, searchParamsPreset)

	log.Println("total post data count : ", len(postData))
	waitSecond := int64(15 * len(postData))
	time.Sleep(time.Millisecond * time.Duration(waitSecond))

	// log.Println("start scraping post details")

	fileStore, pageName, initIndex := store.CreateFetchedJobStore()

	for idx, post := range postData {
		// 0.0277777778
		var postDetail = new(request.SeekPostDetails)
		// var coverLetterResp = new(ollamaApi.GenerateResponse)
		// var wg sync.WaitGroup
		// wg.Add(2)

		// // fix time rate limit
		// go func() {
		// 	defer wg.Done()
		// 	t1 := time.Now()
		// 	time.Sleep(150 * time.Millisecond)
		// 	t2 := time.Now()
		// 	log.Println("api cold down time taken : ", t2.Sub(t1))
		// }()

		// // remark: it take 2 post request to get Job detail
		// go func() {
		// 	defer wg.Done()
		t1 := time.Now()
		time.Sleep(50 * time.Millisecond)

		postDetail = batchPostDetailProcess(&post, searchParamsPreset, &readConfig)
		store.WriteRecordToPage(fileStore, pageName, postDetail, initIndex+1+idx)

		t2 := time.Now()
		log.Println("fetch job time taken : ", t2.Sub(t1))
		// pp.Println("post detail : ", postDetail)
		log.Println("progess : ", idx, " of ", len(postData))
		if idx%32 == 0 {
			if err := store.Save(fileStore); err != nil {
				log.Println(err)
			}
			if err := store.StorePoint(); err != nil {
				log.Println(err)
			}

		}

		// go batchGenerateCoverLetter(postDetail, &readConfig)
		// }()
		// wg.Wait()
	}

	// log.Println("DebugComppanyTags : ")
	// debugStr, _ := json.Marshal(DebugComppanyTag)
	// log.Println(string(debugStr))

	if err := store.StorePoint(); err != nil {
		log.Println(err)
	}
	if err := store.Save(fileStore); err != nil {
		log.Println(err)
	}
}

func getPostPobsList(combinedKeywords []string, searchParamsPreset *seek_api.SeekSearchApiParams) []seek_api.SeekSearchApiResponseData {

	postData := []seek_api.SeekSearchApiResponseData{}
	for _, keywordCombination := range combinedKeywords {
		tempPostData, _ := fetchJobs(searchParamsPreset, keywordCombination)
		// if len(postData) == 0 {
		// 	postData = append(postData, tempPostData...)
		// } else {
		for _, post := range tempPostData {
			isExist, _ := store.CheckKeyPostDetailCache(fmt.Sprintf("%d", post.ID))
			if !isExist {
				postData = append(postData, post)
			}
			// }
		}
	}
	pp.Printf("post data : %d\n", postData)

	return postData
}

func batchPostDetailProcess(post *seek_api.SeekSearchApiResponseData, searchParamsPreset *seek_api.SeekSearchApiParams, readConfig *config.SearchConfig) *request.SeekPostDetails {
	collector := linkedin.CreateSearchEngineCollector()
	gqlData := gql.GetPostDetail(fmt.Sprintf("%d", post.ID))
	store.SetPostGQLCache(&gqlData.Data.JobDetails.Job)

	postDetail := gql.ConvertPostGQLToPostDetail(gqlData)
	store.SetPostDetailCache(postDetail)

	if postDetail.PayRange == "" && post.Salary != "" {
		postDetail.PayRange = post.Salary
	}

	// check company profile is in the cache
	var companyDetail *request.SeekCompanyDetails = nil

	if gqlData.Data.JobDetails.CompanyProfile != nil {
		companyDetail, _ = store.GetCompanyDetailCache(gqlData.Data.JobDetails.CompanyProfile.ID)
	} else if gqlData.Data.JobDetails.Job.Advertiser.ID != "" {
		companyDetail, _ = store.GetCompanyDetailCache(gqlData.Data.JobDetails.Job.Advertiser.ID)
	}

	if companyDetail == nil {
		companyDetail = linkedin.ExtractCompanyProfileGQL(&gqlData.Data.JobDetails)
		jobPosted, _ := linkedin.GetCompanyPostList(searchParamsPreset, post)
		companyDetail.JobPosted = jobPosted
		// companyDetail.ContactPerson = gqlData.Data.JobDetails
		linkedin.SearchCompany(collector, companyDetail)
		store.SetCompanyDetailCache(companyDetail)
	}

	postDetail.CompanyDetails = companyDetail
	contextScore, positiveMatch := scoring.CalculateMatchingScores(readConfig, postDetail)

	postDetail.Score = contextScore
	scoring.OtherScoring(readConfig, post, postDetail)

	postDetail.HittedKeywords = positiveMatch

	// pp.Printf("post detail : %s\n", postDetail)
	return postDetail
}

func batchGenerateCoverLetter(postDetail *request.SeekPostDetails, readConfig *config.SearchConfig) {
	coverLetterResp := ollamaApi.GenerateCoverLetter(readConfig, postDetail)
	// pp.Println("cover letter : ", coverLetterResp.Response)
	coverLetterContent := template.GenerateCoverLetterMail(readConfig, postDetail, coverLetterResp)
	log.Println("cover letter content : ", coverLetterContent)
}

func fetchJobs(preset *seek_api.SeekSearchApiParams, keyword string) ([]seek_api.SeekSearchApiResponseData, error) {
	postData := []seek_api.SeekSearchApiResponseData{}

	data, err := seek_api.SeekSearchApiWithPreset(preset, keyword, 1, "", "")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// log.Printf("totel count : %d\n", data.TotalCount)

	pageTotal := data.TotalPages

	for pageNumber := 1; pageNumber <= pageTotal; pageNumber++ {
		time.Sleep(80 * time.Millisecond)
		data, err := seek_api.SeekSearchApiWithPreset(preset, keyword, pageNumber, "", data.UserQueryID)
		if err != nil {
			log.Println(err)
		}
		postData = append(postData, data.Data...)
	}
	return postData, nil
}
