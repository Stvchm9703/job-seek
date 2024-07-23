package scoring

import (
	"job-seek/pkg/config"
	"job-seek/pkg/request"
	seekAPI "job-seek/pkg/request/seek_api"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func OtherScoring(config *config.SearchConfig, post *seekAPI.SeekSearchApiResponseData, postDetail *request.SeekPostDetails) {
	calculateSalaryScore(config, post, postDetail)
	calculateCompanyFactorScore(config, post, postDetail)
}

func calculateSalaryScore(config *config.SearchConfig, post *seekAPI.SeekSearchApiResponseData, postDetail *request.SeekPostDetails) {
	score := 0

	tragetMinSalary, _ := strconv.Atoi(config.SearchParamsPreset.MinSalary)
	tragetMaxSalary, _ := strconv.Atoi(config.SearchParamsPreset.MaxSalary)

	salaryRange := strings.Split(post.Salary, "-")
	salaryRange = lo.Map(salaryRange, func(item string, idx int) string {
		txt := strings.TrimSpace(item)
		txt = strings.ReplaceAll(txt, "$", "")
		if strings.Contains(txt, "K") || strings.Contains(txt, "k") {
			txt = strings.ReplaceAll(txt, "K", "000")
			txt = strings.ReplaceAll(txt, "k", "000")
		}
		return txt
	})
	if len(salaryRange) == 2 {
		minSalary, _ := strconv.Atoi(salaryRange[0])
		maxSalary, _ := strconv.Atoi(salaryRange[1])

		if minSalary >= tragetMinSalary {
			score += 3
		}
		if maxSalary >= tragetMaxSalary {
			score += 3
		}
	}
	postDetail.Score += score
}

var (
	companySizeEnum = []string{"1-10", "11-50", "51-200", "201-500", "501-1000", "1001-5000", "5001-10000", "10001+"}
)

func calculateCompanyFactorScore(config *config.SearchConfig, post *seekAPI.SeekSearchApiResponseData, postDetail *request.SeekPostDetails) {
	// check the required Job post compare to the company size
	tmpScore := 0
	groupSize := strings.ReplaceAll(postDetail.CompanyDetails.GroupSize, " employees", "")
	// log.Println("groupSize", groupSize)

	indstr := strings.TrimSpace(groupSize)
	indstr = strings.ReplaceAll(indstr, ",", "")
	indstr = strings.ReplaceAll(indstr, "+", "")
	sizeRange := strings.Split(indstr, "-")

	sizeRangeInt := lo.Map(sizeRange, func(item string, _ int) int {
		val, _ := strconv.Atoi(item)
		return val
	})

	sizeIndex := lo.IndexOf(companySizeEnum, groupSize)
	rangeSizeCap := lo.IndexOf(companySizeEnum, config.SearchParamsPreset.CompanySize)

	if sizeIndex > 0 && len(sizeRangeInt) > 1 {
		if (sizeRangeInt[1] / 2) < postDetail.CompanyDetails.JobPosted {
			tmpScore -= 3
			postDetail.HittedKeywords = append(postDetail.HittedKeywords, "[neg:hiring rate too high]")

		} else if sizeRangeInt[0] < postDetail.CompanyDetails.JobPosted {
			tmpScore -= 2
			postDetail.HittedKeywords = append(postDetail.HittedKeywords, "[neg:hiring rate too high]")

		}
		// check the company scaling
	} else {
		tmpScore += 1
	}
	if sizeIndex > rangeSizeCap {
		tmpScore += 4
	} else {
		tmpScore += sizeIndex
	}

	spelities := strings.ToLower(strings.Join(postDetail.CompanyDetails.Specialties, ", "))
	// reject if the company is a recruitment agency
	if strings.Contains(strings.ToLower(postDetail.CompanyDetails.Name), "recruit") ||
		strings.Contains(strings.ToLower(postDetail.CompanyDetails.Description), "recruit") ||
		strings.Contains(spelities, "recruit") ||
		strings.Contains(strings.ToLower(postDetail.CompanyDetails.Name), "agency") ||
		strings.Contains(strings.ToLower(postDetail.CompanyDetails.Description), "agency") ||
		strings.Contains(spelities, "agency") {
		tmpScore -= 5
		postDetail.HittedKeywords = append(postDetail.HittedKeywords, "[neg:recruitment agency]")
	} else if postDetail.CompanyDetails.Name == "" || postDetail.CompanyDetails.Industry == "" {
		tmpScore -= 3
		postDetail.HittedKeywords = append(postDetail.HittedKeywords, "[neg:company name or industry missing]")

	}

	// log.Println("tmpScore", tmpScore)

	// check the company industry

	postDetail.Score += tmpScore

}
