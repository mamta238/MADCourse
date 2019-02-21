package usercrudhandler

import (
	"bytes"
	"fmt"
	logger "log"

	customerrors "packages/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateRestCreateUpdateRequest(rStr string) (bool, error) {
	//logger := loggerutils.GetLogger()
	schemaStr := `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		
		"properties": {
		  "name": {
            "type": "string",
			"minLength": 1,
			"maxLength": 30
		  },
		  "address": {
            "type": "string",
			"minLength": 1,
			"maxLength": 30
		  },
		  "address_line_2": {
            "type": "string",
			"minLength": 1,
			"maxLength": 30
		  },
		  "url": {
            "type": "string",
			"minLength": 1,
			"maxLength": 30
		  },
		  "outcode": {
            "type": "string",
			"minLength": 1,
			"maxLength": 10
		  },
		  "postcode": {
            "type": "string",
			"minLength": 1,
			"maxLength": 10
		  },
		  "rating": {
            "type": float
		  },
		  "type_of_food": {
            "type": "string",
			"minLength": 1,
			"maxLength": 15
		  },
		  
		},
		"required": [
		  "name",
		  "address",
		  "address_line_2",
		  "url",
		  "outcode",
		  "postcode",
		  "rating",
		  "type_of_food"
		]
	  }`

	schema := gojsonschema.NewStringLoader(schemaStr)
	content := gojsonschema.NewStringLoader(rStr)
	result, err := gojsonschema.Validate(schema, content)
		logger.Println("Hello")
	if err != nil {
		logger.Fatalf("Invalid Json Schema Error: %v", err)
		return false, customerrors.InternalError(fmt.Sprintf("Invalid Json Schema Error: %v", err))
		//panic(err)
	}
	if result.Valid() {
		return true, nil
	} else {
		var buffer bytes.Buffer
		for _, resulterr := range result.Errors() {
			logger.Printf("- %s\n", resulterr)
			errString := fmt.Sprintf("Field: %s - %s, ", resulterr.Field(), resulterr.Description())
			buffer.Write([]byte(errString))

		}
		errorDesc := buffer.String()
		return false, customerrors.BadRequest(errorDesc)
	}

}
