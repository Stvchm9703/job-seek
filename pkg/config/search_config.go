package config

import (
	"github.com/BurntSushi/toml"
)

type SearchConfig struct {
	IgnoreKeywords     *[]string                 `toml:"ignore_keywords"`
	SearchKeywords     *[]string                 `toml:"search_keywords"`
	SearchParamsPreset *SearchParamsPresetConfig `toml:"search_params_preset"`
	Template           *TemplateConfig           `toml:"template"`
	SeekToken          *SeekToken                `toml:"seek_token"`
	ApiEndpoint        struct {
		Ollama *ApiEndpoint `toml:"ollama"`
	} `toml:"api_endpoint"`
}

type SearchParamsPresetConfig struct {
	UserId         string `toml:"user_id"`
	SiteKey        string `toml:"site_key"`
	SalaryType     string `toml:"salary_type"`
	MinSalary      string `toml:"min_salary"`
	MaxSalary      string `toml:"max_salary"`
	LangLocale     string `toml:"lang_locale"`
	WorkLocale     string `toml:"work_locale"`
	WorkType       string `toml:"work_type"`
	Classification string `toml:"classification"`
	CompanySize    string `toml:"company_size"`
}

type TemplateConfig struct {
	CoverLetterPath string `toml:"cover_letter_path"`
	CvPath          string `toml:"cv_path"`
	MailPath        string `toml:"mail_path"`
}

type SeekToken struct {
	SessionId    string `toml:"session_id" json:"session_id"`
	ClientId     string `toml:"client_id" json:"client_id"`
	AccessToken  string `toml:"access_token" json:"access_token"`
	ExpiresIn    int    `toml:"expires_in" json:"expires_in"`
	TokenType    string `toml:"token_type" json:"token_type"`
	IdToken      string `toml:"id_token" json:"id_token"`
	RefreshToken string `toml:"refresh_token" json:"refresh_token"`
	Scope        string `toml:"scope" json:"scope"`
}

type ApiEndpoint struct {
	Url    string `toml:"url"`
	Model  string `toml:"model"`
	Prompt string `toml:"prompt"`
}

func ReadSearchConfig() SearchConfig {
	var config SearchConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
	return config
}

func CreateSearchConfig() ([]byte, error) {
	return toml.Marshal(SearchConfig{})
}
