package schema_test

import (
	"job-seek/pkg/database/schema"
	"testing"
)

var (
	tagName = "surrealdb"
)

type TestDataModelA struct {
	ReferenceId   string            `surrealdb:"field:id,primaryId"`
	Name          string            ``
	Count         int               `surrealdb:"as:int"`
	CreatedAt     string            `surrealdb:"as:datetime,autoCreateTime"`
	UpdatedAt     string            `surrealdb:"as:datetime,autoUpdateTime"`
	DeletedAt     string            `surrealdb:"as:string,autoDeleteTime"`
	RawData       []byte            `surrealdb:""`
	ModelBs       []TestDataModelB  `surrealdb:"as:record"`
	ModelBsPtr    []*TestDataModelB `surrealdb:"as:record"`
	ModelBonly    TestDataModelB    `surrealdb:"as:record"`
	ModelBonlyPtr *TestDataModelB   `surrealdb:"as:record"`
	ModelCObject  TestDataModelC    `surrealdb:"as:object"`
}

type TestDataModelB struct {
	ReferenceId string `surrealdb:"mapTo:id,primaryId"`
	Name        string ``
	Url         string ``
	Count       int    `surrealdb:"as:int"`
}

type TestDataModelC struct {
	Rid string
}

func TestGetFieldMap(t *testing.T) {
	sample := TestDataModelA{}
	result := schema.GetFieldMap(sample)
	t.Logf("%v", result)

}

func TestDefineModelString(t *testing.T) {
	sample := TestDataModelA{}
	result, _ := schema.DefineModelString(sample)
	t.Logf("%v", result)
}

func TestUnmarshal(t *testing.T) {
	sample := TestDataModelA{}
	sampleData := map[string]interface{}{
		"ReferenceId": "123",
		"Name":        "test",
		"Count":       1,
		"CreatedAt":   "2024-07-31T13:09:13.732387211+00:00",
		"UpdatedAt":   "2024-07-31T13:09:13.732387211+00:00",
		"DeletedAt":   "2024-07-31T13:09:13.732387211+00:00",
		"RawData":     []byte("test"),
		"ModelBs": []map[string]interface{}{
			{
				"ReferenceId": "123",
				"Name":        "test",
				"Url":         "test",
				"Count":       1,
			},
		},
		"ModelBsPtr": []map[string]interface{}{
			{
				"ReferenceId": "123",
				"Name":        "test",
				"Url":         "test",
				"Count":       1,
			},
		},
		"ModelBonly": map[string]interface{}{
			"ReferenceId": "123",
			"Name":        "test",
			"Url":         "test",
			"Count":       1,
		},
		"ModelBonlyPtr": map[string]interface{}{
			"ReferenceId": "123",
			"Name":        "test",
			"Url":         "test",
			"Count":       1,
		},
		"ModelCObject": map[string]interface{}{
			"Rid": "123",
		},
	}

	schema.Unmarshal(&sample, sampleData)
	t.Logf("%v", sample)
}
