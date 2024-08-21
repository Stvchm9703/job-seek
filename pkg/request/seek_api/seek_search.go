package seek_api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	jsConfig "job-seek/pkg/config"

	"github.com/dghubble/sling"
	"github.com/k0kubun/pp/v3"
	combinations "github.com/mxschmitt/golang-combinations"
	lo "github.com/samber/lo"
)

var (
	SeekToken = &jsConfig.SeekToken{}
	client    = &http.Client{}
)

type SeekApiTokenRequest struct {
	JobseekerSessionId string `json:"JobseekerSessionId"`
	ClientId           string `json:"client_id"`
	GrantType          string `json:"grant_type"`
	IdentitySdkVersion string `json:"identity_sdk_version"`
	InitialScope       string `json:"initial_scope"`
	RedirectUri        string `json:"redirect_uri"`
	RefreshHref        string `json:"refresh_href"`
	RefreshToken       string `json:"refresh_token"`
}

// !deprecated
// func SeekRefreshToken(config *jsConfig.SeekToken) error {
// 	client := sling.New().
// 		Base("https://login.seek.com/").
// 		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

// 	resData := SeekApiTokenRequest{
// 		JobseekerSessionId: config.SessionId,
// 		ClientId:           config.ClientId,
// 		GrantType:          "refresh_token",
// 		IdentitySdkVersion: "6.54.0",
// 		InitialScope:       config.Scope,
// 		RedirectUri:        "https://www.seek.com.au/oauth/callback/",
// 		RefreshHref:        "https://www.seek.com.au/data-analyst-golang-jobs",
// 		RefreshToken:       config.RefreshToken,
// 	}

// 	responseData := jsConfig.SeekToken{}

// 	_, err := client.Post("oauth/token").
// 		BodyJSON(resData).
// 		Receive(responseData, nil)

// 	if err != nil {
// 		pp.Printf("refresh error:", err)
// 		return err
// 	}

// 	pp.Println("refreshed token", responseData)

// 	SeekToken.AccessToken = responseData.AccessToken
// 	SeekToken.ExpiresIn = responseData.ExpiresIn
// 	SeekToken.RefreshToken = responseData.RefreshToken
// 	SeekToken.IdToken = responseData.IdToken

// 	return nil
// }

// !deprecated
// func SeekSearchApiSimply(seekerId string, searchKeyWord string, pageNumber int) (SeekSearchApiResponse, *http.Response, error) {

// 	client := sling.New().Base("https://www.seek.com.au/api/chalice-search/v4/").
// 		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3").
// 		Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjVZSDBuSkpwRzYyTDFBZ2tEU3hfeiJ9.eyJodHRwOi8vc2Vlay9jbGFpbXMvY291bnRyeSI6IkFVIiwiaHR0cDovL3NlZWsvY2xhaW1zL2JyYW5kIjoic2VlayIsImh0dHA6Ly9zZWVrL2NsYWltcy9leHBlcmllbmNlIjoiY2FuZGlkYXRlIiwiaHR0cDovL3NlZWsvY2xhaW1zL3VzZXJfaWQiOiI1NzM3NTIzNTQiLCJpc3MiOiJodHRwczovL2xvZ2luLnNlZWsuY29tLyIsInN1YiI6ImF1dGgwfDY1ZjkzNzQxYTJmZWFlOTdiNjcyZTE3MyIsImF1ZCI6WyJodHRwczovL3NlZWsvYXBpL2NhbmRpZGF0ZSIsImh0dHBzOi8vc2Vla2Fuei5vbmxpbmVhdXRoLnByb2Qub3V0ZnJhLnh5ei91c2VyaW5mbyJdLCJpYXQiOjE3MjAzMjg1MDgsImV4cCI6MTcyMDMzMjEwOCwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBvZmZsaW5lX2FjY2VzcyIsImF6cCI6InlHQlZnZTY2SzVOSnBTTjV1NzFmVTkwVmNUbEVBU051In0.YA_u-TwHAr-j1E0FmBi376mp6W3rKeWKLADQZQdyqN_xLxM7lC6EJyttPYmTx0uM7JZ0SdMaqzlTYTcg3mySSVAMc4ah4wZnlwVipRa1jRYtZdDo_eDqXzK_JSaFm1MfvIuKiS7gA7P9K9mvgEZac6JfK9esGw_zsi6wgVnoIUkl5Xb7zuFCiJJWkHDl34MGinJK-xNzn4pAD0LmbdW9PArRMBl5bb_0GoLB4vqQ3nPnCBUgIc2xbB2fb6iP1zJK-DIcEsdGDu9LxtrXdX8NeW_skBMjbJ_4clpSruwj_301adr03Ox1P-whb42KRQCr83qwVeK2aWkHvJUHD5KztA")
// 	params := &SeekSearchApiParams{
// 		SiteKey:     "AU-Main",
// 		Where:       "All Sydney NSW",
// 		Page:        pageNumber,
// 		Keywords:    searchKeyWord,
// 		SalaryType:  "annual",
// 		SalaryRange: "100000-",
// 		Locale:      "en-AU",
// 		SeekerId:    seekerId,
// 	}

// 	responseData := SeekSearchApiResponse{}

// 	res, err := client.Get("search").
// 		QueryStruct(params).
// 		Receive(&responseData, nil)

// 	if err != nil {
// 		return responseData, res, err
// 	}

// 	return responseData, res, nil
// }

func SeekSearchApiWithPreset(paramsPreset *SeekSearchApiParams, searchKeyWord string, pageNumber int, userId string, userQueryId string) (SeekSearchApiResponse, error) {

	client := sling.New().Base("https://www.seek.com.au/api/chalice-search/v4/").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
		// Set("Authorization", fmt.Sprintf("%s %s", SeekToken.TokenType, SeekToken.AccessToken))

	params := paramsPreset
	params.Page = pageNumber
	params.Keywords = searchKeyWord
	params.UserQueryId = userQueryId
	params.SeekerId = userId
	params.PageSize = 100
	responseData := SeekSearchApiResponse{}

	res, err := client.Get("search").
		QueryStruct(params).
		Receive(&responseData, nil)

	if err != nil {
		log.Println("error:", res.StatusCode, err)
		// if res.StatusCode == 429 {
		// 	SeekRefreshToken(SeekToken)
		// 	return SeekSearchApiWithPreset(paramsPreset, searchKeyWord, pageNumber)
		// }
		return responseData, err
	}

	return responseData, nil
}

func SeekSearchApiForApi(paramsPreset *SeekSearchApiParams, config *jsConfig.SeekServiceConfig) (SeekSearchApiResponse, error) {
	// pp.Println("SeekSearchApiForApi", config)
	client := sling.New().Base(config.Domain+"/api/chalice-search/v4/").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	params := paramsPreset
	params.PageSize = 100
	responseData := SeekSearchApiResponse{}
	if params.Locale == "" {
		params.Locale = "en"
	}
	res, err := client.Get("search").
		QueryStruct(params).
		Receive(&responseData, nil)
	if res != nil {
		fmt.Printf("req url: %s \n", res.Request.URL)
		pp.Println("query", res.Request.URL.Query())
	}

	if err != nil {
		return responseData, err
	}

	return responseData, nil
}

func CreateSeekSearchApiParamsFromConfig(config *jsConfig.SearchParamsPresetConfig) *SeekSearchApiParams {
	return &SeekSearchApiParams{
		SiteKey:        config.SiteKey,
		Where:          config.WorkLocale,
		Page:           1,
		Keywords:       "",
		SalaryType:     config.SalaryType,
		SalaryRange:    fmt.Sprintf("%s-%s", config.MinSalary, config.MaxSalary),
		Locale:         config.LangLocale,
		SeekerId:       config.UserId,
		Classification: config.Classification,
	}
}

func CreateSearchCombinations(keywords []string) []string {
	// exclude the important keywords

	importantKW := lo.Filter(keywords, func(kw string, _ int) bool {
		return strings.Contains(kw, "!") || strings.Contains(kw, "*")
	})

	importantKW = lo.Map(importantKW, func(kw string, _ int) string {
		s := strings.ReplaceAll(kw, "!", "")
		return strings.ReplaceAll(s, "*", "")
	})

	// log.Println("importantKW  ", importantKW)

	casualKW := lo.Filter(keywords, func(kw string, _ int) bool {
		return !strings.Contains(kw, "!") && !strings.Contains(kw, "*")
	})

	combinedKeywords := combinations.All(casualKW)
	// log.Println("combinedKeywords  ", combinedKeywords)

	casualKW = []string{}
	if len(importantKW) == 0 {
		for _, kw := range combinedKeywords {
			casualKW = append(casualKW, strings.Join(kw, " "))
		}

		return casualKW
	}
	for _, kw := range combinedKeywords {
		for _, imp := range importantKW {
			kw := append(kw, imp)
			casualKW = append(casualKW, strings.Join(kw, " "))
		}
	}
	// log.Println("kw  ", casualKW)

	return casualKW

}
