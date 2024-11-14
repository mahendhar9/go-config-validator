package schema

import (
	"fmt"
)

// Represents the allowed data types in schema
type SchemaType string

const (
	TypeString  SchemaType = "string"
	TypeNumber  SchemaType = "number"
	TypeBoolean SchemaType = "boolean"
	TypeObject  SchemaType = "object"
	TypeArray   SchemaType = "array"
)

type SchemaValidator struct {
	schema map[string]interface{}
}

// Creates a new SchemaValidator instance
func NewSchemaValidator(schema map[string]interface{}) *SchemaValidator {
	return &SchemaValidator{
		schema: schema,
	}
}

// Validates schema against provided config
func (v *SchemaValidator) Validate(config map[string]interface{}) []error {
	var errors []error
	return v.validateObject(config, v.schema, "", &errors)
}

func (v *SchemaValidator) validateObject(config, schema map[string]interface{}, path string, errors *[]error) []error {
	// Check required fields
	if required, ok := schema["required"].([]interface{}); ok {
		for _, req := range required {
			if reqField, ok := req.(string); ok {
				if _, exists := config[reqField]; !exists {
					*errors = append(*errors, fmt.Errorf("missing required field: %s", joinPath(path, reqField)))
				}
			}
		}
	}

	// Check properties
	if properties, ok := schema["properties"].(map[string]interface{}); ok {
		for fieldName, fieldSchema := range properties {
			fieldPath := joinPath(path, fieldName)

			// Get field schema
			fieldSchemaDef, ok := fieldSchema.(map[string]interface{})
			if !ok {
				*errors = append(*errors, fmt.Errorf("invalid schema definition for field: %s", fieldPath))
				continue
			}

			// Check if field exists in config
			configValue, exists := config[fieldName]
			if !exists {
				continue // Field is not required if not in required list
			}

			// Validate field type
			if err := v.validateType(configValue, fieldSchemaDef, fieldPath); err != nil {
				*errors = append(*errors, err)
			}

			// Recursively validate nested objects
			if fieldSchemaDef["type"] == "object" && configValue != nil {
				if configObj, ok := configValue.(map[string]interface{}); ok {
					v.validateObject(configObj, fieldSchemaDef, fieldPath, errors)
				}
			}
		}
	}

	return *errors
}

// Checks if a value matches its schema type
func (v *SchemaValidator) validateType(value interface{}, schema map[string]interface{}, path string) error {
	schemaType, ok := schema["type"].(string)
	if !ok {
		return fmt.Errorf("missing type definition for field: %s", path)
	}

	switch SchemaType(schemaType) {
	case TypeString:
		if _, ok := value.(string); !ok {
			return fmt.Errorf("field %s must be string", path)
		}
	case TypeNumber:
		switch value.(type) {
		case float64, int, int64, float32:
			// Valid number types
		default:
			return fmt.Errorf("field %s must be a number", path)
		}
	case TypeBoolean:
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("field %s must be a boolean", path)
		}
	case TypeArray:
		if _, ok := value.([]interface{}); !ok {
			return fmt.Errorf("field %s must be an array", path)
		}
	case TypeObject:
		if _, ok := value.(map[string]interface{}); !ok {
			return fmt.Errorf("field %s must be an object", path)
		}
	default:
		return fmt.Errorf("unsupported type %s for field: %s", schemaType, path)
	}

	return nil
}

func joinPath(base, field string) string {
	if base == "" {
		return field
	}
	return base + "." + field
}
