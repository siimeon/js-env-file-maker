package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	var environmentVariablePrefix = flag.String("prefix", "APP_", "Environment variable prefix")
	var jsFileName = flag.String("f", "env.js", "JS file name")
	flag.Parse()

	var filteredEnvironmentVariables = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasPrefix(pair[0], *environmentVariablePrefix) {
			filteredEnvironmentVariables[strings.TrimPrefix(pair[0], *environmentVariablePrefix)] = pair[1]
		}
	}

	variablesTemplate, err := template.New("variablesTemplate").Parse("{{ range $key, $value := . }} {{ $key }}: \"{{ $value }}\",{{ end }}")
	if err != nil {
		panic(err)
	}
	var parsedVariablesTemplate bytes.Buffer
	err = variablesTemplate.Execute(&parsedVariablesTemplate, filteredEnvironmentVariables)
	if err != nil {
		panic(err)
	}
	var variablesTemplateString = parsedVariablesTemplate.String()
	variablesTemplateString = strings.TrimSuffix(variablesTemplateString, ",")
	variablesTemplateString = strings.TrimPrefix(variablesTemplateString, " ")
	fileTemplate, err := template.New("fileTemplate").Parse("window.env = { {{.}} }")
	errorCheck(err)
	var parsedFileTemplate bytes.Buffer
	err = fileTemplate.Execute(&parsedFileTemplate, variablesTemplateString)
	errorCheck(err)

	file, err := os.Create(*jsFileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, parsedFileTemplate.String())
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
