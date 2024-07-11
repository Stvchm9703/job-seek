package template

import (
	"job-seek/pkg/config"
	"job-seek/pkg/request"
	"job-seek/pkg/request/generation_service/ollama"
	"regexp"
)

type CoverLetterTemplate struct {
	Receiver string
	Title    string
	Content  string
}

func GenerateCoverLetterMail(readConfig *config.SearchConfig, postDetail *request.SeekPostDetails, coverLetter *ollama.GenerateResponse) string {
	pureContent := coverLetter.Response
	tagRegex := regexp.MustCompile(`\[[\s\w',()-]+\]`)
	pureContent = tagRegex.ReplaceAllString(pureContent, "")

	// emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// template.New("cover_letter").Parse(pureContent)
	// templateFile := readConfig.Template.CoverLetterPath
	// tmpl, _ := template.ParseFiles(readConfig.Template.CoverLetterPath)

	// var body bytes.Buffer
	// content := CoverLetterTemplate{
	// 	Receiver: "Hiring Manager",
	// 	Title:    postDetail.PostTitle,
	// 	Content:  pureContent,
	// }
	// if err := tmpl.Execute(&body, content); err != nil {
	// 	log.Fatalf("Failed to execute template: %v", err)
	// }

	// log.Printf("Mail content: %s\n", body.String())

	return pureContent

}
