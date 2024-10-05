// Path: job-seek/pkg/database/model/
// code generated by tools/generate_db_model_query/main.go

package model

import (
	"bytes"
	"errors"
	"fmt"
	"job-seek/pkg/protos"
	"strings"
	"text/template"

	"github.com/k0kubun/pp/v3"
	"github.com/samber/lo"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type PreferenceKeywordModel struct {
	Id         string `json:"kw_id"`
	UserId     string `json:"user_id"`
	Keyword    string `json:"keyword"`
	Value      string `json:"value"`
	Type       string `json:"type"`
	IsPositive bool   `json:"is_positive"`
}

func (m *PreferenceKeywordModel) ToProto() *protos.PreferenceKeyword {
	return &protos.PreferenceKeyword{
		KwId:       m.Id,
		UserId:     m.UserId,
		Keyword:    m.Keyword,
		Value:      m.Value,
		Type:       m.Type,
		IsPositive: m.IsPositive,
	}
}

func (m *PreferenceKeywordModel) FromProto(p *protos.PreferenceKeyword) {
	m.Id = p.KwId
	m.UserId = p.UserId
	m.Keyword = p.Keyword
	m.Value = p.Value
	m.Type = p.Type
	m.IsPositive = p.IsPositive
}

func (m *PreferenceKeywordModel) GetModel(db *surrealdb.DB) (*protos.PreferenceKeyword, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	result, err := db.Select(fmt.Sprintf("PreferenceKeyword:%s", m.Id))
	if err != nil {
		return nil, err
	}

	data := new(PreferenceKeywordModel)
	err = surrealdb.Unmarshal(result, data)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to unmarshal PreferenceKeywordModel"), err, pp.Errorf("result", result))
		// return nil, err
	}

	return data.ToProto(), nil

}

func (m *PreferenceKeywordModel) ListModel(db *surrealdb.DB) ([]*protos.PreferenceKeyword, error) {
	query := fmt.Sprintf(`
	SELECT * , search::score(1) as score FROM PreferenceKeyword WHERE Value @1@ "%s" ORDER BY score DESC;
	`, m.Value)
	result, err := db.Query(query, nil)
	if err != nil {
		return nil, err
	}
	details := []*PreferenceKeywordModel{}
	err = surrealdb.Unmarshal(result, details)

	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to unmarshal CompanyDetailModel"), err, pp.Errorf("result", result))
	}
	fun := lo.Map(details, func(item *PreferenceKeywordModel, index int) *protos.PreferenceKeyword {
		return item.ToProto()
	})
	return fun, nil
}

func (m *PreferenceKeywordModel) CreateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	if m.UserId == "" {
		return fmt.Errorf("UserId is empty")
	}

	// result, err := sd.Create("PreferenceKeyword", m)

	queryTemplate, _ := template.New("createCompanyDetail").Parse(`
INSERT INTO CompanyDetail  {
	UserId: r"{{.UserId}}",
	Keyword: s"{{.Keyword}}",
	Value: s"{{.Value}}",
	Type: s"{{.Type}}",
	IsPositive {{.IsPositive}},
};
	`)
	var doc bytes.Buffer
	var err error
	err = queryTemplate.Execute(&doc, m)
	if err != nil {
		return err
	}
	// _, err := sd.Create(fmt.Sprintf("CompanyDetail:%s", m.ReferenceId), m)
	query := strings.ReplaceAll(doc.String(), "\n", " ")
	query = strings.ReplaceAll(query, "\t", " ")
	query = strings.ReplaceAll(query, "\r", " ")
	// query = strings.ReplaceAll(query, "\"", "'")
	query = strings.Join(strings.Fields(strings.TrimSpace(query)), " ")
	result, err := sd.Query(query, m)
	var message map[string]interface{}
	surrealdb.Unmarshal(result, message)
	if err != nil {
		fmt.Println("query:", query)
		pp.Println("message:", message)
		return errors.Join(err, fmt.Errorf("query: %s", query), pp.Errorf("message: %v", message))
	}
	return nil
}

func (m *PreferenceKeywordModel) UpdateModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	_, err := sd.Update(fmt.Sprintf("PreferenceKeyword:%s", m.Id), m)
	return err
}

func (m *PreferenceKeywordModel) DefineModel(sd *surrealdb.DB) error {
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	query := `
DEFINE TABLE IF NOT EXISTS PreferenceKeyword SCHEMAFULL;
-- Field definition
	DEFINE FIELD IF NOT EXISTS	UserId 					ON TABLE PreferenceKeyword TYPE		record<UserAccount>;
	DEFINE FIELD IF NOT EXISTS	Keyword 				ON TABLE PreferenceKeyword TYPE		string;
	DEFINE FIELD IF NOT EXISTS	Value 					ON TABLE PreferenceKeyword TYPE		string;
	DEFINE FIELD IF NOT EXISTS	Type 						ON TABLE PreferenceKeyword TYPE		string;
	DEFINE FIELD IF NOT EXISTS	IsPositive			ON TABLE PreferenceKeyword TYPE		bool;

`
	_, err := sd.Query(query, nil)
	return err
}