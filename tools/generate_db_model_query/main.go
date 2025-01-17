package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
	pp "github.com/k0kubun/pp/v3"
	protoparser "github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

func main() {
	protoPath := flag.String("proto", "", "proto path")
	flag.Parse()

	if *protoPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	//read the proto file and parse it
	reader, err := os.Open(*protoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s, err %v\n", *protoPath, err)
		return
	}
	defer reader.Close()

	got, err := protoparser.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse, err %v\n", err)
		return
	}

	os.Mkdir(modelPath, os.ModePerm)

	for _, v := range got.ProtoBody {
		if s, ok := v.(*parser.Message); ok {

			if strings.HasSuffix(s.MessageName, "Request") || strings.HasSuffix(s.MessageName, "Response") || strings.HasPrefix(s.MessageName, "Get") {
				// fmt.Println("skip : ", s.MessageName)
				continue
			}
			pp.Println("Message: ", s)
			generateDataModel(*s)
		}
	}

}

var (
	protobufDefaultType = []string{
		"int32", "uint32", "sint32", "fixed32", "sfixed32",
		"int64", "uint64", "sint64", "fixed64", "sfixed64",
		"float",
		"double",
		"bool",
		"string",
		"bytes",
	}

	modelPath = "pkg/database/model"

	modelHeader = `
	// Path: job-seek/pkg/database/model/
	// code generated by tools/generate_db_model_query/main.go

	
	package model
	import (
		"fmt"
		"job-seek/pkg/config"
		"job-seek/pkg/database"
		"job-seek/pkg/database/schema"
		"job-seek/pkg/protos"	
		surrealdb "github.com/surrealdb/surrealdb.go"
	)
	`
)

func generateDataModel(message parser.Message) {
	// generate model folder
	modelName := message.MessageName

	outputFile, err := os.Create(fmt.Sprintf("%s/%s.go", modelPath, strcase.ToSnake(modelName)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create file %s, err %v\n", fmt.Sprintf("%s/%s.go", modelPath, strcase.ToSnake(modelName)), err)
		return
	}
	defer outputFile.Close()

	if _, err = outputFile.WriteString(modelHeader); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file %s.go, err %v\n", strcase.ToSnake(modelName), err)
		return
	}

	modelStructDefined := generateSurrealQueryModel(message)
	if _, err = outputFile.WriteString(modelStructDefined); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file %s.go, err %v\n", strcase.ToSnake(modelName), err)
		return
	}

	modelToProtocDefined := generateSurrealQueryModelToProtos(message)
	if _, err = outputFile.WriteString(modelToProtocDefined); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write function ToProto , %s.go, err %v\n", strcase.ToSnake(modelName), err)
		return
	}

	modelFromProtocDefined := generateSurrealQueryModelFromProto(message)
	if _, err = outputFile.WriteString(modelFromProtocDefined); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write function FromProto %s.go, err %v\n", strcase.ToSnake(modelName), err)
		return
	}

}

func generateSurrealQueryModel(message parser.Message) string {
	// generate model folder

	var modelStructDefined string = fmt.Sprintf("type %sModel struct {\n", message.MessageName)

	var fieldDef []string = []string{
		"schema.Model",
	}

	for _, v := range message.MessageBody {
		if field, ok := v.(*parser.Field); ok {

			fieldLine := []string{
				strcase.ToCamel(field.FieldName),
			}

			dataType := field.Type
			isMessgeStruct := false
			jsonTag := []string{
				field.FieldName,
			}
			if strings.HasPrefix(field.Type, "int") {
				dataType = "int"
			} else if field.Type == "double" {
				dataType = "float"
			} else if field.Type == "bytes" {
				dataType = "[]byte"
			} else if !slices.Contains(protobufDefaultType, field.Type) {
				dataType = "string"
				isMessgeStruct = true
			}
			fieldLine = append(fieldLine, dataType)

			if field.IsOptional {
				// fieldLine = append(fieldLine, fmt.Sprintf("*%s", dataType))
				jsonTag = append(jsonTag, "omitempty")
			}
			if field.IsRepeated {
				println("repeated : ", field.Type)
				fieldLine = append(fieldLine, fmt.Sprintf("[]%s", dataType))
			}
			if isMessgeStruct {
				fieldLine = append(fieldLine, fmt.Sprintf("// reference to %s as record link", field.Type))
			} else {
				fieldLine = append(fieldLine, "// field value")
			}

			fieldDef = append(
				fieldDef,
				fmt.Sprintf("%s\t%s\t`json:\"%s\"`", fieldLine[0], fieldLine[1], strings.Join(jsonTag, ",")),
			)

		}
	}
	modelStructDefined += strings.Join(fieldDef, "\n")
	modelStructDefined += "\n}\n"

	return modelStructDefined
}

func generateSurrealQueryModelToProtos(message parser.Message) string {
	var funcDefined string = fmt.Sprintf("func(m *%sModel) ToProto() *protos.%s {\n", message.MessageName, message.MessageName)

	funcDefined += fmt.Sprintf("return &protos.%s{\n", message.MessageName)
	var fieldArgs []string = []string{}

	for _, v := range message.MessageBody {
		if field, ok := v.(*parser.Field); ok {

			value := fmt.Sprintf("m.%s", strcase.ToCamel(field.FieldName))
			// funcDefined += fmt.Sprintf("%s: m.%s,\n", strcase.ToCamel(field.FieldName), strcase.ToCamel(field.FieldName))
			if (field.Type != "string" && field.Type != "bool" && field.Type != "bytes") && slices.Contains(protobufDefaultType, field.Type) {
				value = fmt.Sprintf("%s( m.%s )", field.Type, strcase.ToCamel(field.FieldName))
			}
			if field.IsOptional {
				value = fmt.Sprintf("&%s", value)
				// } else if field.IsRepeated && slices.Contains(protobufDefaultType, field.Type) {
				// 	value = fmt.Sprintf("m.%s", strcase.ToCamel(field.FieldName))
			} else if !field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
				value = fmt.Sprintf("m.%s.ToProto()", strcase.ToCamel(field.FieldName))
			} else if field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
				value = fmt.Sprintf("m.%s.ToProto()", strcase.ToCamel(field.FieldName))
			}
			fieldArgs = append(fieldArgs, fmt.Sprintf("%s: %s ,", strcase.ToCamel(field.FieldName), value))
		}
	}
	funcDefined += strings.Join(fieldArgs, "\n")
	funcDefined += "\n}\n"
	funcDefined += "}\n"
	return funcDefined
}

func generateSurrealQueryModelFromProto(message parser.Message) string {
	var funcDefined string = fmt.Sprintf("\nfunc(m *%sModel) FromProto(p *protos.%s) {\n", message.MessageName, message.MessageName)
	var fieldArgs []string = []string{}

	for _, v := range message.MessageBody {
		if field, ok := v.(*parser.Field); ok {

			value := fmt.Sprintf("p.%s", strcase.ToCamel(field.FieldName))
			// funcDefined += fmt.Sprintf("%s: m.%s,\n", strcase.ToCamel(field.FieldName), strcase.ToCamel(field.FieldName))
			if (field.Type != "string" && field.Type != "bool") && slices.Contains(protobufDefaultType, field.Type) {
				if strings.HasPrefix(field.Type, "int") {
					value = fmt.Sprintf("int(p.%s)", strcase.ToCamel(field.FieldName))
				} else if field.Type == "double" {
					value = fmt.Sprintf("float64(p.%s)", strcase.ToCamel(field.FieldName))
				} else if strings.HasPrefix(field.Type, "float") {
					value = fmt.Sprintf("float(p.%s)", strcase.ToCamel(field.FieldName))
				}
			}

			if field.IsOptional {
				value = fmt.Sprintf("p.Get%s()", value)
			} else if field.IsRepeated && slices.Contains(protobufDefaultType, field.Type) {
				value = fmt.Sprintf("p.%s", strcase.ToCamel(field.FieldName))
			} else if !field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
				// value = fmt.Sprintf("p.%s.ToProtoc()", strcase.ToCamel(field.FieldName))
			} else if field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
				// value = fmt.Sprintf("m.%s.ToProtoc()", strcase.ToCamel(field.FieldName))
			}
			fieldArgs = append(fieldArgs, fmt.Sprintf("m.%s = %s", strcase.ToCamel(field.FieldName), value))
		}
	}
	funcDefined += strings.Join(fieldArgs, "\n")
	funcDefined += "\n}\n"
	return funcDefined
}

func generateSurrealQueryModeCreateQuery(message parser.Message) string {
	var funcDefined string = fmt.Sprintf("\nfunc(m *%sModel) Create%s(sd *surrealdb.DB) error {\n", message.MessageName, message.MessageName)

	funcDefined += `
	if sd == nil {
		return fmt.Errorf("database connection is nil")
	}
	`
	funcDefined += fmt.Sprintf("_, err := sd.Create(\"%s:m.\", m)\n", strcase.ToSnake(message.MessageName))
	funcDefined += `return err`

	// var fieldArgs []string = []string{}

	// for _, v := range message.MessageBody {
	// 	if field, ok := v.(*parser.Field); ok {

	// 		value := fmt.Sprintf("p.%s", strcase.ToCamel(field.FieldName))
	// 		// funcDefined += fmt.Sprintf("%s: m.%s,\n", strcase.ToCamel(field.FieldName), strcase.ToCamel(field.FieldName))
	// 		if field.IsOptional {
	// 			value = fmt.Sprintf("p.Get%s()", value)
	// 		} else if field.IsRepeated && slices.Contains(protobufDefaultType, field.Type) {
	// 			value = fmt.Sprintf("p.%s", strcase.ToCamel(field.FieldName))
	// 		} else if !field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
	// 			// value = fmt.Sprintf("p.%s.ToProtoc()", strcase.ToCamel(field.FieldName))
	// 		} else if field.IsRepeated && !slices.Contains(protobufDefaultType, field.Type) {
	// 			// value = fmt.Sprintf("m.%s.ToProtoc()", strcase.ToCamel(field.FieldName))
	// 		}
	// 		fieldArgs = append(fieldArgs, fmt.Sprintf("m.%s = %s", strcase.ToCamel(field.FieldName), value))
	// 	}
	// }
	// funcDefined += strings.Join(fieldArgs, "\n")
	funcDefined += "\n}\n"
	return funcDefined
}
