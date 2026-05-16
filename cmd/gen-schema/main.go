package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/autobrr/qui/internal/schema"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Schema generation failed: %v", err)
	}
	log.Println("Schema successfully updated.")
}

func run() error {
	modRoot, err := schema.FindModRoot()
	if err != nil {
		return fmt.Errorf("failed to locate module root: %w", err)
	}

	schemaPath := filepath.Join(modRoot, schema.ConfigSchemaPath)

	if err := os.MkdirAll(filepath.Dir(schemaPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to open schema file: %w", err)
	}
	defer f.Close()

	if err := schema.ExportSchema(f); err != nil {
		return fmt.Errorf("failed to export schema: %w", err)
	}

	return nil
}
