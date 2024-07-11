package seek_gql

type JobDetailResponse struct {
	Data JobDetailData `json:"data"`
}
type JobDetailData struct {
	JobDetails JobDetails `json:"jobDetails"`
}

type JobDetails struct {
	Job                   Job                   `json:"job"`
	CompanyProfile        *CompanyProfile       `json:"companyProfile"`
	CompanySearchURL      string                `json:"companySearchUrl"`
	CompanyTags           []interface{}         `json:"companyTags"`
	RestrictedApplication RestrictedApplication `json:"restrictedApplication"`
	Sourcr                interface{}           `json:"sourcr"`
	GfjInfo               GfjInfo               `json:"gfjInfo"`
}

type CompanyProfile struct {
	CompanyNameSlug string `json:"companyNameSlug"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Overview        struct {
		Description struct {
			Paragraphs []string `json:"paragraphs"`
		} `json:"description"`
		Industry string `json:"industry"`
		Size     struct {
			Description string `json:"description"`
		} `json:"size"`
		Website struct {
			URL string `json:"url"`
		} `json:"website"`
	} `json:"overview"`
	PerksAndBenefits []struct {
		Title string `json:"title"`
	} `json:"perksAndBenefits"`
	ReviewsSummary struct {
		OverallRating struct {
			NumberOfReviews struct {
				Value int `json:"value"`
			} `json:"numberOfReviews"`
			Value float64 `json:"value"`
		} `json:"overallRating"`
	} `json:"reviewsSummary"`
	ShouldDisplayReviews bool `json:"shouldDisplayReviews"`
}

type GfjInfo struct {
	Location  Location  `json:"location"`
	WorkTypes WorkTypes `json:"workTypes"`
}

type Location struct {
	CountryCode string      `json:"countryCode"`
	Country     string      `json:"country"`
	Suburb      string      `json:"suburb"`
	Region      interface{} `json:"region"`
	State       string      `json:"state"`
	Postcode    string      `json:"postcode"`
}

type WorkTypes struct {
	Label []string `json:"label"`
}

type Job struct {
	ID              string                  `json:"id"`
	Title           string                  `json:"title"`
	PhoneNumber     interface{}             `json:"phoneNumber"`
	IsExpired       bool                    `json:"isExpired"`
	ExpiresAt       ExpiresAt               `json:"expiresAt"`
	IsLinkOut       bool                    `json:"isLinkOut"`
	ContactMatches  []ContactMatch          `json:"contactMatches"`
	IsVerified      bool                    `json:"isVerified"`
	Abstract        string                  `json:"abstract"`
	Content         string                  `json:"content"`
	Status          string                  `json:"status"`
	ListedAt        ListedAt                `json:"listedAt"`
	Salary          Salary                  `json:"salary"`
	WorkTypes       RestrictedApplication   `json:"workTypes"`
	Advertiser      Advertiser              `json:"advertiser"`
	Location        RestrictedApplication   `json:"location"`
	Classifications []RestrictedApplication `json:"classifications"`
	Products        Products                `json:"products"`
}

type Advertiser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RestrictedApplication struct {
	Label *string `json:"label"`
}

type ContactMatch struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ExpiresAt struct {
	DateTimeUTC string `json:"dateTimeUtc"`
}

type ListedAt struct {
	Label       string `json:"label"`
	DateTimeUTC string `json:"dateTimeUtc"`
}

type Products struct {
	Bullets       []string      `json:"bullets"`
	Questionnaire Questionnaire `json:"questionnaire"`
}

type Questionnaire struct {
	Questions []string `json:"questions"`
}

type Salary struct {
	CurrencyLabel interface{} `json:"currencyLabel"`
	Label         string      `json:"label"`
}
