package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mahendhar9/go-config-validator/internal/schema"
	"github.com/mahendhar9/go-config-validator/pkg/parser"
)

func main() {
	configFile := flag.String("config", "", "Path to the configuration file (JSON or YAML)")
	schemaFile := flag.String("schema", "", "Path to the schema file (JSON or YAML)")

	flag.Parse()

	if *configFile == "" || *schemaFile == "" {
		fmt.Println("Error: Both config and schema files are required")
		flag.Usage()
		os.Exit(1)
	}

	log.SetPrefix("go-config-validator: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("Validating config file: %s against schema %s\n", *configFile, *schemaFile)

	// Parse config file
	configParser, err := parser.NewParser(*configFile)
	if err != nil {
		log.Fatalf("Failed to create config parser: %v", err)
	}

	config, err := configParser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// Parse schema file
	schemaParser, err := parser.NewParser(*schemaFile)
	if err != nil {
		log.Fatalf("Failed to create schema parser: %v", err)
	}

	schemaData, err := schemaParser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse schema file: %v", err)
	}

	// Validate config against schema
	validator := schema.NewSchemaValidator(schemaData)
	if errors := validator.Validate(config); len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Println("Configuration is valid!")
}
