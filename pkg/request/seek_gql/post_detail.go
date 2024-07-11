package seek_gql

import (
	"fmt"
	"log"
	"strings"

	"job-seek/pkg/request"

	mdConv "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/dghubble/sling"
	"github.com/samber/lo"
)

type GQLRequest struct {
	Query     string            `json:"query"`
	Variables JobDetailVariable `json:"variables"`
}

type JobDetailVariable struct {
	CountryCode  string `json:"countryCode"`
	JobID        string `json:"jobId"`
	LanguageCode string `json:"languageCode"`
	Locale       string `json:"locale"`
	Timezone     string `json:"timezone"`
	Zone         string `json:"zone"`
}

var JobDetailQuery string = `
query jobDetails(
$jobId: ID!, $zone: Zone!, $locale: Locale!, $languageCode: LanguageCodeIso!, $countryCode: CountryCodeIso2!, $timezone: Timezone!) {
  jobDetails(id: $jobId) {
    job {
      id
      title
      phoneNumber
      isExpired
      expiresAt {
        dateTimeUtc
      }
      isLinkOut
      contactMatches {
        type
        value
      }
      isVerified
      abstract
      content(platform: WEB)
      status
      listedAt {
        label(context: JOB_POSTED, length: SHORT, timezone: $timezone, locale: $locale)
        dateTimeUtc
      }
      salary {
        currencyLabel(zone: $zone)
        label
      }
      workTypes {
        label(locale: $locale)
      }
      advertiser {
        id
        name(locale: $locale)
      }
      location {
        label(locale: $locale, type: LONG)
      }
      classifications {
        label(languageCode: $languageCode)
      }
      products {
        bullets
        questionnaire {
          questions
        }
      }
      
    }
    companyProfile(zone: $zone) {
      id
      name
      companyNameSlug
      shouldDisplayReviews
      overview {
        description {
          paragraphs
        }
        industry
        size {
          description
        }
        website {
          url
        }
      }
      reviewsSummary {
        overallRating {
          numberOfReviews {
            value
          }
          value
        }
      }
      perksAndBenefits {
        title
      }
    }
    companySearchUrl(zone: $zone, languageCode: $languageCode)
    companyTags {
      key(languageCode: $languageCode)
      value
    }
    restrictedApplication(countryCode: $countryCode) {
      label(locale: $locale)
    }
    sourcr {
      image
      imageMobile
      link
    }
    gfjInfo {
      location {
        countryCode
        country(locale: $locale)
        suburb(locale: $locale)
        region(locale: $locale)
        state(locale: $locale)
        postcode
      }
      workTypes {
        label
      }
    }
  }
}
`

func GetPostDetail(jobId string) *JobDetailResponse {
	url := "https://www.seek.com.au/graphql"

	requestData := &GQLRequest{
		Query: JobDetailQuery,
		Variables: JobDetailVariable{
			CountryCode:  "AU",
			JobID:        jobId,
			LanguageCode: "en",
			Locale:       "en-AU",
			Timezone:     "Australia/Sydney",
			Zone:         "anz-1",
		},
	}

	responseData := new(JobDetailResponse)

	_, err := sling.New().Post(url).
		Set("Accept", "*/*").
		Set("Content-Type", "application/json").
		BodyJSON(requestData).
		Receive(responseData, nil)

	if err != nil {
		log.Println("GQL error", err)
	}

	return responseData

}

func ConvertPostGQLToPostDetail(postGQL *JobDetailResponse) *request.SeekPostDetails {
	// postId, _ := strconv.Atoi(postGQL.Data.JobDetails.Job.ID)
	converter := mdConv.NewConverter("", true, nil)
	mdString, _ := converter.ConvertString(postGQL.Data.JobDetails.Job.Content)
	mdString = strings.ReplaceAll(mdString, "\u00a0", "\n")

	RoleDisplay := lo.Map(postGQL.Data.JobDetails.Job.Classifications, func(item RestrictedApplication, index int) string {
		return *item.Label
	})

	// companyProfile := postGQL.Data.JobDetails.CompanyProfile

	workType := ""
	if postGQL.Data.JobDetails.Job.WorkTypes.Label != nil {
		workType = *postGQL.Data.JobDetails.Job.WorkTypes.Label
	}

	location := ""
	if postGQL.Data.JobDetails.Job.Location.Label != nil {
		location = *postGQL.Data.JobDetails.Job.Location.Label
	}

	return &request.SeekPostDetails{
		PostTitle:    postGQL.Data.JobDetails.Job.Title,
		PayRange:     postGQL.Data.JobDetails.Job.Salary.Label,
		PostUrl:      fmt.Sprintf("https://www.seek.com.au/job/%s", postGQL.Data.JobDetails.Job.ID),
		PostId:       postGQL.Data.JobDetails.Job.ID,
		DebugText:    mdString,
		Role:         strings.Join(RoleDisplay, ", "),
		WorkType:     workType,
		Locations:    location,
		ExpiringDate: postGQL.Data.JobDetails.Job.ExpiresAt.DateTimeUTC,
	}
}
