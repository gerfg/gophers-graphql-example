package schema

import (
	"io/ioutil"
	"log"
)

// ReadSchemaFile read a schema file
func ReadSchemaFile(fileName string) string {
	schemaFile, err := ioutil.ReadFile("schema/" + fileName)
	if err != nil {
		log.Printf("Err:", err)
		return ""
	}
	return string(schemaFile)
}
