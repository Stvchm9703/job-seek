package schema

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	pp "github.com/k0kubun/pp/v3"
	"github.com/samber/lo"
)

var (
	tagName = "surrealdb"
)

type DataModelImpl interface {
	ToProto() interface{}
	FromProto(interface{})
	GetModelQuery(interface{}) (interface{}, error)
	CreateModelQuery(interface{}) error
	DefineModelQuery() error
	GetId() string
	SetId(string)
	getFieldList() []DataModelFieldSet
	// getStructTag(string) []string
	FromSurrealInterface(interface{}) (bool, error)
}

type DataModel struct {
	DataModelImpl
	fieldSet *[]DataModelFieldSet
	id       string `json:"id" surrealdb:"primaryId"`
}

type DataModelFieldSet struct {
	fieldName string
	// string, int, float, bool,  []string, []int, []float, []bool, record, []record
	fieldType        string
	isPrimaryId      bool
	mapToField       string
	mapAsType        string
	isArray          bool
	isPointer        bool
	isAutoCreateTime bool
	isAutoUpdateTime bool
	isAutoDeleteTime bool
	isSkipUpdate     bool
}

func GetFieldMap(m any) []DataModelFieldSet {
	fieldMap := make([]DataModelFieldSet, 0)
	if reflect.TypeOf(m).Kind() != reflect.Struct {
		return fieldMap
	}
	numOfField := reflect.TypeOf(m).NumField()
	fmt.Println(numOfField)
	for i := 0; i < numOfField; i++ {
		field := reflect.TypeOf(m).Field(i)
		tagString := field.Tag.Get(tagName)
		tags := strings.Split(tagString, ",")

		fieldSet := DataModelFieldSet{
			fieldName:        field.Name,
			fieldType:        field.Type.String(),
			isPrimaryId:      slices.Contains(tags, "primaryId"),
			mapToField:       "",
			mapAsType:        "",
			isArray:          field.Type.Kind() == reflect.Slice,
			isPointer:        field.Type.Kind() == reflect.Ptr,
			isAutoCreateTime: slices.Contains(tags, "autoCreateTime"),
			isAutoUpdateTime: slices.Contains(tags, "autoUpdateTime"),
			isAutoDeleteTime: slices.Contains(tags, "autoDeleteTime"),
			isSkipUpdate:     slices.Contains(tags, "skipUpdate"),
		}
		mapToTag, _, hasMapTo := lo.FindIndexOf(tags, func(tag string) bool {
			return strings.HasPrefix(tag, "field")
		})
		if hasMapTo {
			tmp := strings.Split(mapToTag, ":")
			fieldSet.mapToField = tmp[1]
		}

		mapAsTypeTag, _, hasMapAsType := lo.FindIndexOf(tags, func(tag string) bool {
			return strings.HasPrefix(tag, "as")
		})
		if hasMapAsType {
			tmp := strings.Split(mapAsTypeTag, ":")
			fieldSet.mapAsType = tmp[1]
		}
		fieldMap = append(fieldMap, fieldSet)
	}
	// pp.Println("fieldMap:", fieldMap)
	return fieldMap
}

func DefineModelString(m any) (string, error) {
	if reflect.TypeOf(m).Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}
	structName := reflect.TypeOf(m).Name()
	fmt.Println("structName:", structName)
	fieldList := GetFieldMap(m)
	writer := strings.Builder{}
	writer.WriteString("-- Table definition \n")
	writer.WriteString(fmt.Sprintf("DEFINE TABLE IF NOT EXISTS %s SCHEMAFULL;\n", structName))
	writer.WriteString("-- Field definition \n")
	for _, field := range fieldList {
		fieldType := field.fieldType

		if field.mapAsType != "" {
			fieldType = field.mapAsType
		} else if field.fieldType == "[]byte" || field.fieldType == "[]uint8" {
			fieldType = "bytes"
		}
		if field.isArray && !(field.fieldType == "[]byte" || field.fieldType == "[]uint8") {
			fieldType = fmt.Sprintf("array<%s>", fieldType)
		} else if field.isPointer {
			fieldType = fmt.Sprintf("optional<%s>", fieldType)
		}

		if fieldType == "object" {
			// fieldType = "FEXIABLE object"
			structName += " FLEXIBLE"
		}

		fieldName := field.fieldName
		if field.mapToField != "" {
			fieldName = field.mapToField
		}
		if !field.isPrimaryId {
			writer.WriteString(fmt.Sprintf("DEFINE FIELD IF NOT EXISTS %s ON TABLE %s TYPE %s;\n", fieldName, structName, fieldType))
		} else {
			writer.WriteString(fmt.Sprintf("--- Alias  %s as id;\n", field.fieldName))

		}
	}
	writer.WriteString("-- EVENT definition \n")

	autoCreateTimeField := lo.Filter(fieldList, func(field DataModelFieldSet, _ int) bool {
		return field.isAutoCreateTime
	})
	if len(autoCreateTimeField) > 0 {
		writer.WriteString(fmt.Sprintf("DEFINE EVENT IF NOT EXISTS CreateHook ON TABLE %s  \n", structName))
		writer.WriteString("WHEN $event = \"CREATE\" OR $event = \"INSERT\" \n ")
		writer.WriteString("THEN {\n ")
		writer.WriteString(fmt.Sprintf("UPDATE  %s SET \n", structName))
		for _, field := range autoCreateTimeField {
			writer.WriteString(fmt.Sprintf("   %s = time::format(time::now(),\"%s\")", field.fieldName, "%+"))
			if field != autoCreateTimeField[len(autoCreateTimeField)-1] {
				writer.WriteString(",\n")
			}
		}
		writer.WriteString("WHERE id = $after.id ; \n ")
		writer.WriteString("}; \n ")
	}

	autoUpdateTimeField := lo.Filter(fieldList, func(field DataModelFieldSet, _ int) bool {
		return field.isAutoUpdateTime
	})

	if len(autoUpdateTimeField) > 0 {
		writer.WriteString(fmt.Sprintf("DEFINE EVENT IF NOT EXISTS CreateHook ON TABLE %s  \n", structName))
		writer.WriteString("WHEN $event = \"CREATE\" OR $event = \"INSERT\"  \n ")
		writer.WriteString("THEN {\n ")
		writer.WriteString(fmt.Sprintf("UPDATE  %s SET \n", structName))
		for _, field := range autoUpdateTimeField {
			writer.WriteString(fmt.Sprintf("   %s = time::format(time::now(),\"%s\")", field.fieldName, "%+"))
			if field != autoUpdateTimeField[len(autoUpdateTimeField)-1] {
				writer.WriteString(",\n")
			}
		}
		writer.WriteString("WHERE id = $after.id ; \n ")
		writer.WriteString("}; \n ")
	}
	// fmt.Println("writer:", writer.String())
	return writer.String(), nil
}

func Unmarshal(m interface{}, data map[string]interface{}) error {

	if reflect.TypeOf(m).Kind() != reflect.Ptr {
		return fmt.Errorf("input is not a struct")
	}
	// structName := reflect.TypeOf(m).Kind().String()
	// fieldList := GetFieldMap(m)
	// numOfField := reflect.TypeOf(m)
	v := reflect.ValueOf(m).Elem()
	t := reflect.TypeOf(m).Elem()

	pp.Println("s:", t, v)
	fieldList := GetFieldMap(&m)
	pp.Println("s:", t, v)

	// dataColumns := make([]interface{}, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		instField := t.Field(i)
		tag := instField.Tag.Get(tagName)
		tagList := strings.Split(tag, ",")
		if slices.Contains(tagList, "skipUpdate") {
			continue
		}
		if slices.Contains(tagList, "primaryId") {
			if data["id"] != nil {
				v.Field(i).Set(reflect.ValueOf(data["id"]))
			} else {
				v.Field(i).Set(reflect.ValueOf(data[instField.Name]))
			}
			continue
		}
		if slices.Contains(tagList, "field") {
			tagValue, _ := lo.Find(tagList, func(tagSet string) bool {
				return strings.HasPrefix(tagSet, "field")
			})
			tagValue = strings.Split(tagValue, ":")[1]
			v.Field(i).Set(reflect.ValueOf(data[tagValue]))
			continue
		}
		pp.Println("instField:", instField.Name)
		v.Field(i).Set(reflect.ValueOf(data[instField.Name]))
		// instField.Elem().Set(reflect.ValueOf(p).Elem())
	}

	pp.Println("fieldList:", fieldList)

	return nil
}
