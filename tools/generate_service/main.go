package main

import (
	"flag"

	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	pp "github.com/k0kubun/pp/v3"
	protoparser "github.com/yoheimuta/go-protoparser/v4"
	parser "github.com/yoheimuta/go-protoparser/v4/parser"
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
	// pp.Println(got)

	for _, v := range got.ProtoBody {
		if s, ok := v.(*parser.Service); ok {
			gnerateService(*s)

			// pp.Sprintf("Service: %v\n", s)
			pp.Println("Service: ", s)
		}
	}

}
func gnerateService(service parser.Service) {
	// generate service folder
	serviceName := service.ServiceName
	servicePathName := serviceName
	if !strings.HasSuffix(strings.ToLower(servicePathName), "service") {
		servicePathName = servicePathName + "_service"
	}

	servicePathName = strcase.ToSnake(servicePathName)
	serviceCMDName := "JS" + strcase.ToCamel(serviceName)

	param := templateParams{
		ServiceName:          serviceName,
		ServicePathName:      servicePathName,
		ServiceCMDName:       serviceCMDName,
		ServiceNameWithSpace: strings.ToLower(serviceName) + "service",
		// ProtoServiceName : 	got.ProtoBody.ServiceName,
	}
	// craete service folder
	os.Mkdir(
		path.Join("services", servicePathName),
		os.ModePerm)

	// copy the main file
	generateFile(
		path.Join("services", servicePathName, "main.go"),
		"tools/generate_service/main_template/main.go.tmpl",
		param)

	// create cmd folder
	os.Mkdir(
		path.Join("services", servicePathName, "cmd"),
		os.ModePerm)
	generateFile(
		path.Join("services", servicePathName, "cmd", "root.go"),
		"tools/generate_service/cmd_template/root.go.tmpl",
		param)
	generateFile(
		path.Join("services", servicePathName, "cmd", "server.go"),
		"tools/generate_service/cmd_template/server.go.tmpl",
		param)

	// create config folder
	os.Mkdir(
		path.Join("services", servicePathName, "config"),
		os.ModePerm)
	generateFile(
		path.Join("services", servicePathName, "config", "server_conf.go"),
		"tools/generate_service/config_template/server_conf.go.tmpl",
		param)

	// create server folder
	os.Mkdir(
		path.Join("services", servicePathName, "server"),
		os.ModePerm)

	generateFile(
		path.Join("services", servicePathName, "server", "init.go"),
		"tools/generate_service/server_template/init.go.tmpl",
		param)

	// generate service function
	for _, v := range service.ServiceBody {
		if s, ok := v.(*parser.RPC); ok {
			generateServiceFunction(s, &param)
		}

	}

}

type templateParams struct {
	ServiceName          string
	ServicePathName      string
	ServiceCMDName       string
	ServiceNameWithSpace string
}

func generateFile(outputPath string, templatePath string, params templateParams) {
	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to craete file %s, err %v\n", outputPath, err)
	}
	defer output.Close()

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse template %s, err %v\n", templatePath, err)
		return
	}
	err = tmpl.Execute(output, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute template %s, err %v\n", templatePath, err)
	}

}

func generateServiceFunction(rpcfunction *parser.RPC, param *templateParams) {
	// generate service function
	if rpcfunction == nil {
		return
	}
	pp.Println("rpcfunction", rpcfunction)

	outputPath := path.Join("services", param.ServicePathName, "server", strcase.ToSnake(rpcfunction.RPCName)+".go")

	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to craete file %s, err %v\n", outputPath, err)
	}
	defer output.Close()

	templatePath := "tools/generate_service/server_template/_unary_func.go.tmpl"

	if rpcfunction.RPCRequest.IsStream || rpcfunction.RPCResponse.IsStream {
		// to stream function
		templatePath = "tools/generate_service/server_template/_stream_func.go.tmpl"
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse template %s, err %v\n", templatePath, err)
		return
	}

	parm := map[string]string{
		"ServiceName":     param.ServiceName,
		"ServiceCMDName":  param.ServiceCMDName,
		"ServicePathName": param.ServicePathName,
		"RequestType":     strcase.ToCamel(rpcfunction.RPCRequest.MessageType),
		"ResponseType":    strcase.ToCamel(rpcfunction.RPCResponse.MessageType),
		"RPCName":         rpcfunction.RPCName,
	}

	err = tmpl.Execute(output, parm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute template %s, err %v\n", templatePath, err)
	}
}
