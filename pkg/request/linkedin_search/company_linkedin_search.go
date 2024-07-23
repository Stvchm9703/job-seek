package linkedin_search

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"job-seek/pkg/request"
	seekAPI "job-seek/pkg/request/seek_api"
	seekGQL "job-seek/pkg/request/seek_gql"

	lo "github.com/samber/lo"

	"github.com/dghubble/sling"
	"github.com/gocolly/colly"
	pp "github.com/k0kubun/pp/v3"
)

func ExtractCompanyProfile(postData *seekAPI.SeekSearchApiResponseData) *request.SeekCompanyDetails {
	name := postData.CompanyName
	if name == "" {
		name = postData.Advertiser.Description
	}
	return &request.SeekCompanyDetails{
		ReferenceId: postData.Advertiser.ID,
		Name:        postData.CompanyName,
		Description: postData.Advertiser.Description,
	}

}

func ExtractCompanyProfileGQL(post *seekGQL.JobDetails) *request.SeekCompanyDetails {
	if post.CompanyProfile == nil {
		return &request.SeekCompanyDetails{
			ReferenceId: post.Job.Advertiser.ID,
			Name:        post.Job.Advertiser.Name,
			// Description: post.Advertiser.Description,
		}
	}

	return &request.SeekCompanyDetails{
		ReferenceId: post.CompanyProfile.ID,
		Name:        post.CompanyProfile.Name,
		Description: strings.Join(post.CompanyProfile.Overview.Description.Paragraphs, "\n"),
		Industry:    post.CompanyProfile.Overview.Industry,
		GroupSize:   post.CompanyProfile.Overview.Size.Description,
		Url:         post.CompanyProfile.Overview.Website.URL,
		// HeadQuarters: *post.Job.Location.Label,
	}
}

func CreateSearchEngineCollector() *colly.Collector {

	collector := colly.NewCollector(
		// colly.AllowedDomains("www.google.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Cafari/537.36"),
		colly.AllowURLRevisit(),
		// colly.AllowedDomains("au.search.yahoo.com", "au.linkedin.com"),
	)

	collector.Visit("https://au.yahoo.com/")

	return collector
}

func SearchCompany(c *colly.Collector, companyDetail *request.SeekCompanyDetails) *request.SeekCompanyDetails {
	// form submission
	companyName := companyDetail.Name
	if companyDetail.Name == "" {
		log.Println("company name is empty")
		return companyDetail
	}

	if checkFieldIsNotEmpty(companyDetail) {
		log.Println("it look like the company detail is already filled")
		return companyDetail
	}

	link := "https://au.search.yahoo.com/search?p=site%3Alinkedin.com%2Fcompany+%22" + strings.ReplaceAll(companyName, " ", "+") + "%22"
	log.Println("try search link", link)

	c2 := scrapeLinkedinDetail(c, companyDetail)

	c.OnHTML("li.first a", func(e *colly.HTMLElement) {
		// pp.Println(e.DOM.Html())
		resultUrl := e.Attr("href")
		if strings.Contains(resultUrl, "linkedin.com") && strings.Contains(resultUrl, "r.search.yahoo.com") {
			// pp.Println(companyName, resultUrl)
			// companyDetail.Url = resultUrl
			tmpUrl, _ := ExtractFinalURL(resultUrl)
			companyDetail.Linkedin = tmpUrl
			c2.Visit(tmpUrl)
			c2.Wait()

		} else if strings.Contains(resultUrl, "r.search.yahoo.com") {
			// log.Println("not linkedin url")
			tmpUrl, _ := ExtractFinalURL(resultUrl)
			companyDetail.Linkedin = tmpUrl
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		pp.Errorf("error: code : %s, \n %s; \n ", r.StatusCode, err)
	})
	// c.OnScraped(func(r *colly.Response) {
	// 	log.Println("")
	// 	pp.Println("companyDetail", companyDetail)
	// })
	c.Visit(link)

	c.Wait()

	return companyDetail
}

func checkFieldIsNotEmpty(companyDetail *request.SeekCompanyDetails) bool {
	return companyDetail.Description != "" &&
		companyDetail.Industry != "" &&
		companyDetail.GroupSize != "" &&
		companyDetail.HeadQuarters != "" &&
		len(companyDetail.Specialties) == 0 &&
		companyDetail.Url != ""

}

func scrapeLinkedinDetail(c *colly.Collector, companyDetail *request.SeekCompanyDetails) *colly.Collector {
	// converter := mdConv.NewConverter("", true, nil)

	c2 := c.Clone()

	if companyDetail.Description == "" {
		c2.OnHTML("p[data-test-id='about-us__description']", func(e *colly.HTMLElement) {
			companyDetail.Description = strings.TrimSpace(e.Text)
		})
	}

	if companyDetail.Url == "" {
		c2.OnHTML("div[data-test-id='about-us__website'] dd", func(e *colly.HTMLElement) {
			companyDetail.Url = e.Attr("href")
		})
	}

	if companyDetail.Industry == "" {
		c2.OnHTML("div[data-test-id='about-us__industry'] dd", func(e *colly.HTMLElement) {
			companyDetail.Industry = strings.TrimSpace(e.Text)
		})
	}

	if companyDetail.GroupSize == "" {
		c2.OnHTML("div[data-test-id='about-us__size'] dd", func(e *colly.HTMLElement) {
			companyDetail.GroupSize = strings.TrimSpace(e.Text)
		})
	}

	if companyDetail.HeadQuarters == "" {
		c2.OnHTML("div[data-test-id='about-us__headquarters'] dd", func(e *colly.HTMLElement) {
			companyDetail.HeadQuarters = strings.TrimSpace(e.Text)
		})
	}

	if len(companyDetail.Specialties) == 0 {
		c2.OnHTML("div[data-test-id='about-us__specialties'] dd", func(e *colly.HTMLElement) {
			companyDetail.Specialties = lo.Map(strings.Split(e.Text, ","), func(item string, _ int) string {
				return strings.TrimSpace(item)
			})
		})
	}

	c2.OnScraped(func(r *colly.Response) {
		log.Println("finished scraping post details for post id")
		companyDetail.Linkedin = r.Request.URL.String()
		// tmpC2Content := strings.Join(c2Content, "\n")
		// mdString, _ := converter.ConvertString(tmpC2Content)
		// companyDetail.Description = ReplaceMultipleNewlines(mdString)
		// pp.Println("companyDetail", companyDetail)
	})
	return c2
}

func GetCompanyPostList(paramsPreset *seekAPI.SeekSearchApiParams, postData *seekAPI.SeekSearchApiResponseData) (int, error) {
	client := sling.New().Base("https://www.seek.com.au/api/chalice-search/v4/")
	// params := map[string]string{
	// 	"advertiserid":   postData.Advertiser.ID,
	// 	"classification": paramsPreset.Classification,
	// }
	params := &seekAPI.SeekSearchApiParams{
		AdvertiserId:   postData.Advertiser.ID,
		Classification: paramsPreset.Classification,
	}

	responseData := seekAPI.SeekSearchApiResponse{}

	_, err := client.Get("search").
		QueryStruct(params).
		Receive(&responseData, nil)

	if err != nil {
		pp.Println(err)
		return 0, err
	}

	return responseData.TotalCount, nil
}

func GetCompanyPostListByCompanySearchURL(url string) (int, error) {
	// client := sling.New().Base("https://www.seek.com.au/api/chalice-search/v4/")
	// // params := map[string]string{
	// // 	"advertiserid":   postData.Advertiser.ID,
	// // 	"classification": paramsPreset.Classification,
	// // }
	// params := &SeekSearchApiParams{
	// 	AdvertiserId:   postData.Advertiser.ID,
	// 	Classification: paramsPreset.Classification,
	// }

	responseData := seekAPI.SeekSearchApiResponse{}

	_, err := sling.New().Get(url).
		Receive(&responseData, nil)

	if err != nil {
		pp.Println("GetCompanyPostListByCompanySearchURL", err)
		return 0, err
	}

	return responseData.TotalCount, nil
}

func ExtractFinalURL(yahooURL string) (string, error) {
	// Define the regex pattern to match the 'RU' parameter
	regexPattern := `RU=([^/]+)`
	re := regexp.MustCompile(regexPattern)

	// Find the first match
	match := re.FindStringSubmatch(yahooURL)
	if len(match) < 2 {
		return "", fmt.Errorf("no match found")
	}

	// The matched URL is URL-encoded, so we need to decode it
	decodedURL, err := url.QueryUnescape(match[1])
	if err != nil {
		return "", err
	}

	return decodedURL, nil
}

// ReplaceMultipleNewlines replaces multiple consecutive newlines with a single newline
func ReplaceMultipleNewlines(text string) string {
	// Define the regex pattern to match one or more newlines
	regexPattern := `\n+`
	re := regexp.MustCompile(regexPattern)

	// Replace multiple newlines with a single newline
	cleanedText := re.ReplaceAllString(text, "\n")

	return cleanedText
}
