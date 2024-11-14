# Go Config Validator

A configuration validator that ensures your JSON and YAML configuration files conform to predefined schemas. Perfect for validating application configurations, infrastructure definitions, and more.

## Features

- ✅ Supports both JSON and YAML formats
- 🔍 Deep nested object validation
- 📝 Detailed error reporting
- 🎯 Type checking and validation
- 🔒 Required field validation
- 📦 Easy to use as a CLI tool

## Installation

```bash
# Clone the repository
git clone https://github.com/mahendhar9/go-config-validator
cd go-config-validator

# Build the binary
go build -o validator cmd/validator/main.go
```

## Usage

```bash
./validator -config <path-to-config-file> -schema <path-to-schema-file>
```

### Command Line Arguments

- `-config`: Path to your configuration file (JSON or YAML)
- `-schema`: Path to your schema definition file (JSON or YAML)

## Example Configurations

The repository includes example configurations in the `configs/` and `schemas/` directories.

### JSON Configuration Examples

1. Using the JSON schema (`schemas/app-schema.json`):

```bash
# Validate valid JSON config
./validator -config configs/valid-config.json -schema schemas/app-schema.json

# Validate invalid JSON config
./validator -config configs/invalid-config.json -schema schemas/app-schema.json
```

The JSON schema validates a complex application configuration including:

- Application metadata (name, version, environment)
- Server configuration (ports, host, timeouts)
- Database settings (primary and replica configurations)
- Feature flags and rate limiting

Example successful validation output:

```bash
Configuration is valid!
```

Example validation error output:

```bash
Validation errors:
  - field app.name must be a string
  - field server.ports.http must be a number
  - missing required field: app.environment
  - field database.replicas must be an array
```

### YAML Configuration Examples

1. Using the YAML schema (`schemas/app-schema.yaml`):

```bash
# Validate valid YAML config
./validator -config configs/valid-config.yaml -schema schemas/app-schema.yaml

# Validate invalid YAML config
./validator -config configs/invalid-config.yaml -schema schemas/app-schema.yaml
```

The YAML schema validates a complex system configuration including:

- Logging configuration (levels, outputs, retention)
- Cache settings (Redis configuration)
- Monitoring setup (intervals, alerting)
- Security settings (JWT configuration)

## Example Files Structure

```
go-config-validator/
├── schemas/
│   ├── app-schema.json     # JSON schema example
│   └── app-schema.yaml     # YAML schema example
├── configs/
│   ├── valid-config.json   # Valid JSON configuration
│   ├── invalid-config.json # Invalid JSON configuration
│   ├── valid-config.yaml   # Valid YAML configuration
│   └── invalid-config.yaml # Invalid YAML configuration
```

## Supported Types

- `string`: String values
- `number`: Numeric values (integers and floats)
- `boolean`: Boolean values (true/false)
- `object`: Nested objects
- `array`: Lists/arrays

## Project Structure

```
go-config-validator/
├── cmd/
│   └── validator/
│       └── main.go           # CLI entry point
├── internal/
│   └── schema/
│       └── schema.go         # Schema validation logic
├── pkg/
│   └── parser/
│       └── parser.go         # File parser implementation
├── schemas/                  # Example schema files
├── configs/                  # Example config files
└── README.md
```

## Development

### Prerequisites

- Go 1.16 or higher
- Make sure `go mod tidy` is run to install dependencies

### Building from Source

```bash
go build -o validator cmd/validator/main.go
```
