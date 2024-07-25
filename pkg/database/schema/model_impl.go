package schema

import (
	"reflect"
	"slices"
	"strings"

	"github.com/samber/lo"
)

var (
	tagName = "surrealdb"
)

type DataModelImpl interface {
	ToProto() interface{}
	FromProto(interface{})
	GetModel(interface{}) (interface{}, error)
	CreateModel(interface{}) error
	DefineModel() error
	getFieldList() []DataModelFieldSet
	getStructTag(string) []string
}

type DataModel struct {
	DataModelImpl
	fieldSet *[]DataModelFieldSet
}

type DataModelFieldSet struct {
	fieldName string
	// string, int, float, bool,  []string, []int, []float, []bool, record, []record
	fieldType        string
	isPrimaryId      bool
	mapTo            string
	isAutoCreateTime bool
	isAutoUpdateTime bool
	isAutoDeleteTime bool
	isSkipUpdate     bool
}

func (m *DataModel) getStructTag(fieldName string) []string {
	field, ok := reflect.TypeOf(m).Elem().FieldByName(fieldName)
	if !ok {
		panic("Field not found")
	}
	return strings.Split(field.Tag.Get(tagName), ",")
}

func (m *DataModel) getFieldMap() []DataModelFieldSet {
	fieldMap := make([]DataModelFieldSet, 0)
	numOfField := reflect.TypeOf(m).NumField()
	for i := 0; 0 < numOfField; i++ {
		field := reflect.TypeOf(m).Field(i)
		tagString := field.Tag.Get(tagName)
		tags := strings.Split(tagString, ",")

		fieldSet := DataModelFieldSet{
			fieldName:        field.Name,
			fieldType:        field.Type.String(),
			isPrimaryId:      slices.Contains(tags, "primaryId"),
			mapTo:            "",
			isAutoCreateTime: slices.Contains(tags, "autoCreateTime"),
			isAutoUpdateTime: slices.Contains(tags, "autoUpdateTime"),
			isAutoDeleteTime: slices.Contains(tags, "autoDeleteTime"),
			isSkipUpdate:     slices.Contains(tags, "skipUpdate"),
		}
		mapToTag, _, hasMapTo := lo.FindIndexOf(tags, func(tag string) bool {
			return strings.HasPrefix(tag, "mapTo")
		})
		if hasMapTo {
			tmp := strings.Split(mapToTag, ":")
			fieldSet.mapTo = tmp[1]
		}

	}
	return fieldMap
}
