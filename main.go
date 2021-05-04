package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func main() {
	schemaLoader := gojsonschema.NewReferenceLoader(canonicalPath("/json-test/paymentIntent-schema.json"))
	schema, _ := gojsonschema.NewSchema(schemaLoader)

	validateJsonFromFile(schema, "/json-test/001-paymentIntent-required-ok.json")
	validateJsonFromFile(schema, "/json-test/002-paymentIntent-required-error.json")
	validateJsonFromFile(schema, "/json-test/003-paymentIntent-enum-ok.json")
	validateJsonFromFile(schema, "/json-test/004-paymentIntent-enum-error.json")
	validateJsonFromFile(schema, "/json-test/005-paymentIntent-dependent-ok.json")
	validateJsonFromFile(schema, "/json-test/006-paymentIntent-dependent-error.json")
	validateJsonFromFile(schema, "/json-test/007-paymentIntent-dependent-by-values-credit-ok.json")
	validateJsonFromFile(schema, "/json-test/008-paymentIntent-dependent-by-values-debit-ok.json")
	validateJsonFromFile(schema, "/json-test/009-paymentIntent-dependent-by-values-error.json")
	validateJsonFromFile(schema, "/json-test/010-paymentIntent-dependent-by-values-min-max-error.json")
}

func canonicalPath(filePath string) string {
	fileSchemaPrefix := "file://"
	dirPath, _ := os.Getwd()
	canonicalPath := fileSchemaPrefix + dirPath + filePath
	//fmt.Printf("::: Canonical path: [%s]\n", canonicalPath)
	return canonicalPath
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
