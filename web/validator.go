package web

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

const (
	SchemaPath = "/json-test/paymentIntent-schema.json"
)

func Validate(dirPath string) {
	schemaLoader := gojsonschema.NewReferenceLoader(canonicalPath(dirPath + SchemaPath))
	schema, _ := gojsonschema.NewSchema(schemaLoader)
	validateCases(schema, dirPath)
}

func validateCases(schema *gojsonschema.Schema, dirPath string) {
	validateJsonFromFile(schema, dirPath+"/json-test/001-paymentIntent-required-ok.json")
	validateJsonFromFile(schema, dirPath+"/json-test/002-paymentIntent-required-error.json")
	validateJsonFromFile(schema, dirPath+"/json-test/003-paymentIntent-enum-ok.json")
	validateJsonFromFile(schema, dirPath+"/json-test/004-paymentIntent-enum-error.json")
	validateJsonFromFile(schema, dirPath+"/json-test/005-paymentIntent-dependent-ok.json")
	validateJsonFromFile(schema, dirPath+"/json-test/006-paymentIntent-dependent-error.json")
	validateJsonFromFile(schema, dirPath+"/json-test/007-paymentIntent-dependent-by-values-credit-ok.json")
	validateJsonFromFile(schema, dirPath+"/json-test/008-paymentIntent-dependent-by-values-debit-ok.json")
	validateJsonFromFile(schema, dirPath+"/json-test/009-paymentIntent-dependent-by-values-error.json")
	validateJsonFromFile(schema, dirPath+"/json-test/010-paymentIntent-dependent-by-values-min-max-error.json")
}

func validateJsonFromFile(schema *gojsonschema.Schema, jsonFilePath string) {
	fmt.Printf("\n::: Testing [%s]\n", jsonFilePath)
	documentLoader := gojsonschema.NewReferenceLoader(canonicalPath(jsonFilePath))
	validateJson(schema, documentLoader)
}

func validateJson(schema *gojsonschema.Schema, source gojsonschema.JSONLoader) {
	result, err := schema.Validate(source)

	if err != nil {
		fmt.Printf("::: There was an error validating the schema: \n[%v]\n", err)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")

		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func canonicalPath(filePath string) string {
	fileSchemaPrefix := "file://"
	canonicalPath := fileSchemaPrefix + filePath
	//fmt.Printf("::: Canonical path: [%s]\n", canonicalPath)
	return canonicalPath
}
