package request

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
	Specialties  string
	Locations    string
	// ContactPerson string
	// ContactEmail  string
}

// // !deprecated
// func ScrapePostDetails(postId int, config jsConfig.SearchConfig) SeekPostDetails {

// 	c := colly.NewCollector()

// 	link := log.Sprintf("https://www.seek.com.au/job/%d", postId)
// 	scraped := SeekPostDetails{
// 		PostId:    postId,
// 		DebugText: "",
// 		PostUrl:   link,
// 		Score:     0,
// 	}
// 	converter := mdConv.NewConverter("", true, nil)
// 	c.OnRequest(func(r *colly.Request) {
// 		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Cafari/537.36")
// 		r.Headers.Set("Authorization", log.Sprintf("%s %s", SeekToken.TokenType, SeekToken.AccessToken))
// 	})
// 	c.OnHTML("div[data-automation='jobAdDetails']", func(e *colly.HTMLElement) {
// 		// pp.Println(e.Text)
// 		mdString := extractJobDetail(e, converter)
// 		scraped.DebugText = mdString
// 	})
// 	c.OnError(func(r *colly.Response, err error) {
// 		pp.Println("error:", r.StatusCode, err)
// 		// p.StatusCode = r.StatusCode
// 		// if r.StatusCode == 429 {
// 		// SeekRefreshToken(SeekToken)
// 		// return SeekSearchApiWithPreset(paramsPreset, searchKeyWord, pageNumber)
// 		// r.Request.Retry()
// 		// }
// 	})
// 	c.OnScraped(func(r *colly.Response) {
// 		// log.Printf("finished scraping post details for post id %d\n", postId)
// 		// log.Println("start scoping post details")

// 	})
// 	c.Visit(link)
// 	c.Wait()
// 	return scraped
// }

// func extractJobDetail(e *colly.HTMLElement, converter *mdConv.Converter) string {
// 	rawHtml, _ := e.DOM.Html()
// 	mdString, _ := converter.ConvertString(rawHtml)
// 	mdString = strings.ReplaceAll(mdString, "\u00a0", "\n")
// 	return mdString
// }
