package ollama

import (
	"job-seek/pkg/config"
	"job-seek/pkg/request"
	"strings"

	"github.com/dghubble/sling"
	pp "github.com/k0kubun/pp/v3"
)

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Model              string  `json:"model"`
	CreatedAt          string  `json:"created_at"`
	Response           string  `json:"response"`
	Done               bool    `json:"done"`
	DoneReason         string  `json:"done_reason"`
	Context            []int64 `json:"context"`
	TotalDuration      int64   `json:"total_duration"`
	LoadDuration       int64   `json:"load_duration"`
	PromptEvalCount    int64   `json:"prompt_eval_count"`
	PromptEvalDuration int64   `json:"prompt_eval_duration"`
	EvalCount          int64   `json:"eval_count"`
	EvalDuration       int64   `json:"eval_duration"`
}

func GenerateCoverLetter(config *config.SearchConfig, postDetail *request.SeekPostDetails) *GenerateResponse {
	query := config.ApiEndpoint.Ollama.Prompt
	query = strings.ReplaceAll(query, "{{job_description}}", postDetail.DebugText)
	query = strings.ReplaceAll(query, "{{company_information}}", postDetail.CompanyDetails.Description)

	params := &GenerateRequest{
		Model:  config.ApiEndpoint.Ollama.Model,
		Prompt: query,
		Stream: false,
	}

	response := new(GenerateResponse)

	_, err := sling.New().Post(config.ApiEndpoint.Ollama.Url).BodyJSON(params).Receive(response, nil)

	if err != nil {
		pp.Println("err in GenerateCoverLetter", err)
	}

	pp.Println("response", response)

	return response
}
