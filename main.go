package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader(canonicalPath("/json-test/paymentIntent-schema.json"))
	documentLoader := gojsonschema.NewReferenceLoader(canonicalPath("/json-test/paymentIntent-example4.json"))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

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
	dirPath, _ := os.Getwd()
	canonicalPath := fileSchemaPrefix + dirPath + filePath
	fmt.Printf("::: Canonical path: [%s]\n", canonicalPath)
	return canonicalPath
}
