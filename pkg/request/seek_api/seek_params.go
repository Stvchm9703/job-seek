package seek_api

import "time"

type SeekSearchApiParams struct {
	SiteKey        string `url:"siteKey", json:"sitekey"`
	Where          string `url:"where", json:"where"`
	Page           int    `url:"page", json:"page"`
	Keywords       string `url:"keywords", json:"keywords"`
	SalaryType     string `url:"salarytype", json:"salarytype"`
	SalaryRange    string `url:"salaryrange", json:"salaryrange"`
	Locale         string `url:"locale", json:"locale"`
	SeekerId       string `url:"seekerId", json:"seekerId"`
	Classification string `url:"classification", json:"classification"`
	AdvertiserId   string `url:"advertiserid", json:"advertiserid"`
}

type SeekSearchApiResponse struct {
	Data                 []SeekSearchApiResponseData `json:"data"`
	Title                string                      `json:"title"`
	TotalCount           int                         `json:"totalCount"`
	TotalPages           int                         `json:"totalPages"`
	PaginationParameters struct {
		SeekSelectAllPages bool `json:"seekSelectAllPages"`
		HadPremiumListings bool `json:"hadPremiumListings"`
	} `json:"paginationParameters"`
	Info struct {
		TimeTaken  int    `json:"timeTaken"`
		Source     string `json:"source"`
		Experiment string `json:"experiment"`
	} `json:"info"`
	UserQueryID string `json:"userQueryId"`
	SortMode    []struct {
		IsActive bool   `json:"isActive"`
		Name     string `json:"name"`
		Value    string `json:"value"`
	} `json:"sortMode"`
	SolMetadata struct {
		RequestToken     string   `json:"requestToken"`
		Token            string   `json:"token"`
		Keywords         string   `json:"keywords"`
		SortMode         string   `json:"sortMode"`
		SalaryType       string   `json:"salaryType"`
		SalaryCurrency   string   `json:"salaryCurrency"`
		SalaryMin        int      `json:"salaryMin"`
		Locations        []string `json:"locations"`
		LocationDistance int      `json:"locationDistance"`
		PageSize         int      `json:"pageSize"`
		PageNumber       int      `json:"pageNumber"`
		TotalJobCount    int      `json:"totalJobCount"`
		Tags             struct {
			MordorSearchMarket     string `json:"mordor:searchMarket"`
			MordorRt               string `json:"mordor:rt"`
			MordorVs               string `json:"mordor_vs"`
			MordorFlights          string `json:"mordor__flights"`
			MordorRboPerfMngP95K20 string `json:"mordor__rbo_perfMng_p95_k20"`
			ChaliceSearchAPISolID  string `json:"chalice-search-api:solId"`
		} `json:"tags"`
	} `json:"solMetadata"`
	Location struct {
		AreaDescription         string `json:"areaDescription"`
		AreaID                  int    `json:"areaId"`
		Description             string `json:"description"`
		LocationDescription     string `json:"locationDescription"`
		LocationID              int    `json:"locationId"`
		Matched                 bool   `json:"matched"`
		StateDescription        string `json:"stateDescription"`
		SuburbParentDescription string `json:"suburbParentDescription"`
		Type                    string `json:"type"`
		WhereID                 int    `json:"whereId"`
		Descriptions            struct {
			En struct {
				ContextualName string `json:"contextualName"`
			} `json:"en"`
		} `json:"descriptions"`
		ShortLocationName string `json:"shortLocationName"`
	} `json:"location"`

	JoraCrossLink struct {
		CanCrossLink bool `json:"canCrossLink"`
	} `json:"joraCrossLink"`

	SearchParams struct {
		Sitekey     string `json:"sitekey"`
		Where       string `json:"where"`
		Page        string `json:"page"`
		Keywords    string `json:"keywords"`
		Salarytype  string `json:"salarytype"`
		Salaryrange string `json:"salaryrange"`
		Locale      string `json:"locale"`
		Solid       string `json:"solid"`
	} `json:"searchParams"`
}

type SeekSearchApiResponseData struct {
	Advertiser struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"advertiser"`
	Area               string `json:"area"`
	AreaID             int    `json:"areaId"`
	AreaWhereValue     string `json:"areaWhereValue"`
	AutomaticInclusion bool   `json:"automaticInclusion"`
	Branding           struct {
		ID     string `json:"id"`
		Assets struct {
			Logo struct {
				Strategies struct {
					JdpLogo  string `json:"jdpLogo"`
					SerpLogo string `json:"serpLogo"`
				} `json:"strategies"`
			} `json:"logo"`
		} `json:"assets"`
	} `json:"branding,omitempty"`
	BulletPoints   []string `json:"bulletPoints"`
	Classification struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"classification"`
	CompanyName                    string `json:"companyName,omitempty"`
	CompanyProfileStructuredDataID int    `json:"companyProfileStructuredDataId"`
	DisplayStyle                   struct {
		Search string `json:"search"`
	} `json:"displayStyle,omitempty"`
	DisplayType        string `json:"displayType"`
	ListingDateDisplay string `json:"listingDateDisplay"`
	Location           string `json:"location"`
	LocationID         int    `json:"locationId"`
	LocationWhereValue string `json:"locationWhereValue"`
	ID                 int    `json:"id"`
	IsPremium          bool   `json:"isPremium"`
	IsStandOut         bool   `json:"isStandOut"`
	JobLocation        struct {
		Label        string `json:"label"`
		CountryCode  string `json:"countryCode"`
		SeoHierarchy []struct {
			ContextualName string `json:"contextualName"`
		} `json:"seoHierarchy"`
	} `json:"jobLocation"`
	ListingDate time.Time `json:"listingDate"`
	Logo        struct {
		ID          string      `json:"id"`
		Description interface{} `json:"description"`
	} `json:"logo"`
	RoleID         string `json:"roleId"`
	Salary         string `json:"salary"`
	SearchInsights struct {
		UnmatchedKeywords []string `json:"unmatchedKeywords"`
	} `json:"searchInsights"`
	SolMetadata struct {
		SearchRequestToken string `json:"searchRequestToken"`
		Token              string `json:"token"`
		JobID              string `json:"jobId"`
		Section            string `json:"section"`
		SectionRank        int    `json:"sectionRank"`
		JobAdType          string `json:"jobAdType"`
		Tags               struct {
			MordorFlights string `json:"mordor__flights"`
			MordorS       string `json:"mordor__s"`
		} `json:"tags"`
	} `json:"solMetadata"`
	SubClassification struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"subClassification"`
	Suburb           string `json:"suburb"`
	SuburbID         int    `json:"suburbId"`
	SuburbWhereValue string `json:"suburbWhereValue"`
	Teaser           string `json:"teaser"`
	Title            string `json:"title"`
	Tracking         string `json:"tracking"`
	WorkType         string `json:"workType"`
	WorkArrangements struct {
		Data []struct {
			ID    string `json:"id"`
			Label struct {
				Text string `json:"text"`
			} `json:"label"`
		} `json:"data"`
	} `json:"workArrangements"`
	IsPrivateAdvertiser bool `json:"isPrivateAdvertiser"`
	Tags                []struct {
		Type  string `json:"type"`
		Label string `json:"label"`
	} `json:"tags,omitempty"`
}
