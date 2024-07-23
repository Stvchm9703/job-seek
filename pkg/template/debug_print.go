package template

import (
	"fmt"
	"io"
	seekAPI "job-seek/pkg/request/seek_api"
	"strings"

	fastTemplate "github.com/valyala/fasttemplate"
)

func DebugPrintPostListData(postList *[]seekAPI.SeekSearchApiResponseData) string {
	template := "{{idx}}: {{ID}}, {{Title}}, {{WorkType}}, {{Salary}}, {{CompanyName}};"
	engine := fastTemplate.New(template, "{{", "}}")
	templatedCollection := []string{}

	for idx, post := range *postList {
		stringPost := engine.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
			switch tag {
			case "idx":
				return w.Write([]byte(fmt.Sprintf("%d", idx)))
			case "ID":
				return w.Write([]byte(fmt.Sprintf("%d", post.ID)))
			case "Title":
				return w.Write([]byte(post.Title))
			case "WorkType":
				return w.Write([]byte(post.WorkType))
			case "Salary":
				return w.Write([]byte(post.Salary))
			case "CompanyName":
				return w.Write([]byte(post.CompanyName))
			default:
				return w.Write([]byte(fmt.Sprintf("[unknown tag %q]", tag)))
			}
		})
		templatedCollection = append(templatedCollection, stringPost)
	}

	return strings.Join(templatedCollection, "\n")
}
