package model

import (
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SurveyJobPreferenceModel struct {
	gorm.Model
	//  string user_id = 1;
	// string survey_id = 2;
	// repeated job_seek.job_search.Job jobs = 3;
	// repeated SurveyJobQuestionSet survey_job_questions = 4;
	UserID             int                      `json:"-"`
	User               UserAccountModel         `json:"user" gorm:"foreignKey:UserID"`
	SurveyID           int                      `json:"survey_id"`
	SurveyJobQuestions []SurveyJobQuestionModel `json:"survey_job_questions" gorm:"foreignKey:SurveyJobPreferenceID"`
}

func (SurveyJobPreferenceModel) TableName() string {
	return "survey_job_preference"
}

func (SurveyJobPreferenceModel) DefineModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	return sd.AutoMigrate(&SurveyJobPreferenceModel{})
}

type SurveyJobQuestionModel struct {
	gorm.Model
	// string pair_id = 1;
	// SurveyJobSet job_a = 2;
	// SurveyJobSet job_b = 3;
	// repeated string similarities = 4;
	// repeated string differences = 5;
	// optional float overall_similarity = 6;
	SurveyJobPreferenceID uint                                  `json:"-"`
	PairID                string                                `json:"pair_id"`
	JobA                  datatypes.JSONType[SurveyJobSetModel] `json:"job_a"`
	JobB                  datatypes.JSONType[SurveyJobSetModel] `json:"job_b"`
	Similarities          datatypes.JSONSlice[string]           `json:"similarities"`
	Differences           datatypes.JSONSlice[string]           `json:"differences"`
}

func (SurveyJobQuestionModel) TableName() string {
	return "survey_job_question"
}

func (SurveyJobQuestionModel) DefineModel(sd *gorm.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	return sd.AutoMigrate(&SurveyJobQuestionModel{})
}

type SurveyJobSetModel struct {
	// string job_id = 1;
	// float user_preference_score = 2;
	// SurveyJobFeature features = 3;

	JobId               string  `json:"job_id"`
	UserPreferenceScore float32 `json:"user_preference_score"`
	Features            struct {
		// float adjusted_job_title_similarity = 1;
		// float adjusted_job_industry_similarity = 2;
		// float adjusted_company_size = 3;
		// float adjusted_job_sector = 4;
		// float adjusted_company_culture = 5;
		// float adjusted_work_model = 6;
		// float adjusted_salary_expectation = 7;
		// float adjusted_role_type = 8;
		// float adjusted_distance_score = 9;
		// float pay_average_norm = 10;
		// float descriptions_similarity_to_resume = 11;

		AdjustedJobTitleSimilarity     float32 `json:"adjusted_job_title_similarity"`
		AdjustedJobIndustrySimilarity  float32 `json:"adjusted_job_industry_similarity"`
		AdjustedCompanySize            float32 `json:"adjusted_company_size"`
		AdjustedJobSector              float32 `json:"adjusted_job_sector"`
		AdjustedCompanyCulture         float32 `json:"adjusted_company_culture"`
		AdjustedWorkModel              float32 `json:"adjusted_work_model"`
		AdjustedSalaryExpectation      float32 `json:"adjusted_salary_expectation"`
		AdjustedRoleType               float32 `json:"adjusted_role_type"`
		AdjustedDistanceScore          float32 `json:"adjusted_distance_score"`
		PayAverageNorm                 float32 `json:"pay_average_norm"`
		DescriptionsSimilarityToResume float32 `json:"descriptions_similarity_to_resume"`
	} `json:"features"`
}
