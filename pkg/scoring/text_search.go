package scoring

import (
	"fmt"
	"job-seek/pkg/config"
	"job-seek/pkg/request"

	"sort"
	"strings"

	"github.com/k0kubun/pp/v3"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/samber/lo"
)

func ExampleScore() {

	test := []string{"javascript", "typescript"}
	matches := fuzzy.RankFindFold("java", test)
	sort.Sort(matches)
	pp.Println("matches", matches)
}

func ScorePostMatching(config *config.SearchConfig, context string) (int, []string) {
	score := 0
	contextToken := strings.Split(strings.ReplaceAll(context, "\n", " "), " ")
	positiveMatch := []string{}
	for _, keyword := range *config.SearchKeywords {
		matched := fuzzy.RankFindFold(keyword, contextToken)
		for _, match := range matched {
			if match.Distance == 0 {
				score += 2
			} else if match.Distance <= 4 {
				score += 1
			}
			positiveMatch = append(positiveMatch, fmt.Sprintf("[%s:%s]", match.Source, match.Target))
		}
	}

	for _, keyword := range *config.IgnoreKeywords {
		matched := fuzzy.RankFindFold(keyword, contextToken)
		for _, match := range matched {
			if match.Distance == 0 {
				score -= 3
			} else if match.Distance <= 4 {
				score -= 2
			}
		}

	}
	// score += 1

	return score, positiveMatch
}

func ScorePostMatchingKeyword(config *config.SearchConfig, context string) (fuzzy.Ranks, fuzzy.Ranks) {
	// score := 0
	contextToken := strings.Split(strings.ReplaceAll(context, "\n", " "), " ")
	// positiveMatch := []string{}
	posMatch := make(fuzzy.Ranks, 0)
	negMatch := make(fuzzy.Ranks, 0)
	for _, keyword := range *config.SearchKeywords {
		matched := fuzzy.RankFindFold(keyword, contextToken)
		posMatch = append(posMatch, matched...)
	}
	posMatch = lo.Filter(posMatch, func(item fuzzy.Rank, _ int) bool {
		return strings.Contains(strings.ToLower(item.Target), strings.ToLower(item.Source))
	})

	for _, keyword := range *config.IgnoreKeywords {
		matched := fuzzy.RankFindFold(keyword, contextToken)
		negMatch = append(negMatch, matched...)
	}
	negMatch = lo.Filter(negMatch, func(item fuzzy.Rank, _ int) bool {
		return strings.Contains(strings.ToLower(item.Target), strings.ToLower(item.Source))
	})

	return posMatch, negMatch
}

func scoringMatch(posMatch, negMatch fuzzy.Ranks, posFullScore, posPartScore, negFullScore, negPartScore int, ignoreRepeated bool) int {
	score := 0

	pMatch := posMatch
	nMatch := negMatch
	if ignoreRepeated {
		pMatch = make(fuzzy.Ranks, 0)
		keys := lo.KeyBy(posMatch, func(item fuzzy.Rank) string {
			return item.Target
		})
		for keyRef := range keys {
			items := lo.Filter(posMatch, func(item fuzzy.Rank, _ int) bool {
				return item.Target == keyRef
			})
			item := lo.MinBy(items, func(item fuzzy.Rank, minItem fuzzy.Rank) bool {
				return item.Distance < minItem.Distance
			})
			pMatch = append(pMatch, item)
		}

		nMatch = make(fuzzy.Ranks, 0)
		keys = lo.KeyBy(negMatch, func(item fuzzy.Rank) string {
			return item.Target
		})
		for keyRef := range keys {
			items := lo.Filter(negMatch, func(item fuzzy.Rank, _ int) bool {
				return item.Target == keyRef
			})
			item := lo.MinBy(items, func(item fuzzy.Rank, minItem fuzzy.Rank) bool {
				return item.Distance < minItem.Distance
			})
			nMatch = append(nMatch, item)
		}

	}

	for _, match := range pMatch {
		if match.Distance == 0 {
			score += posFullScore
		} else if match.Distance <= 4 {
			score += posPartScore
		}
	}

	for _, match := range nMatch {
		if match.Distance == 0 {
			score -= negFullScore
		} else if match.Distance <= 4 {
			score -= negPartScore
		}
	}
	return score
}
func CalculateMatchingScores(readConfig *config.SearchConfig, postDetail *request.SeekPostDetails) (int, []string) {
	// positiveMatch := make(fuzzy.Ranks, 0)
	// negativeMatch := make(fuzzy.Ranks, 0)

	score := 0

	// title weighting
	positiveMatch, negativeMatch := ScorePostMatchingKeyword(readConfig, postDetail.PostTitle)
	score += scoringMatch(positiveMatch, negativeMatch, 5, 2, 5, 2, false)

	pMatch, nMatch := ScorePostMatchingKeyword(readConfig, postDetail.WorkType)
	score += scoringMatch(pMatch, nMatch, 5, 2, 5, 1, false)
	positiveMatch = append(positiveMatch, pMatch...)
	negativeMatch = append(negativeMatch, nMatch...)

	pMatch, nMatch = ScorePostMatchingKeyword(readConfig, postDetail.Role)
	score += scoringMatch(pMatch, nMatch, 5, 2, 5, 2, true)
	positiveMatch = append(positiveMatch, pMatch...)
	negativeMatch = append(negativeMatch, nMatch...)

	pMatch, nMatch = ScorePostMatchingKeyword(readConfig, postDetail.DebugText)
	score += scoringMatch(pMatch, nMatch, 3, 2, 3, 2, true)
	positiveMatch = append(positiveMatch, pMatch...)
	negativeMatch = append(negativeMatch, nMatch...)

	pMatch, nMatch = ScorePostMatchingKeyword(readConfig, postDetail.CompanyDetails.Description)
	score += scoringMatch(pMatch, nMatch, 2, 1, 2, 1, true)
	positiveMatch = append(positiveMatch, pMatch...)
	negativeMatch = append(negativeMatch, nMatch...)

	pMatch, nMatch = ScorePostMatchingKeyword(readConfig, postDetail.CompanyDetails.Specialties)
	score += scoringMatch(pMatch, nMatch, 2, 1, 2, 1, true)
	positiveMatch = append(positiveMatch, pMatch...)
	negativeMatch = append(negativeMatch, nMatch...)

	debugText := make([]string, 0)

	for _, match := range positiveMatch {
		debugText = append(debugText, fmt.Sprintf("[pos;%s:%s:%d]", match.Target, match.Source, match.Distance))
	}

	for _, match := range negativeMatch {
		debugText = append(debugText, fmt.Sprintf("[neg;%s:%s:%d]", match.Target, match.Source, match.Distance))
	}

	return score, debugText
}
