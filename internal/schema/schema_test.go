package schema

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestSchemaMatchesConfig(t *testing.T) {
	var buf bytes.Buffer
	if err := ExportSchema(&buf); err != nil {
		t.Fatalf("failed to export schema: %v", err)
	}
	newSchema := buf.Bytes()

	modRoot, err := FindModRoot()
	schemaPath := filepath.Join(modRoot, ConfigSchemaPath)

	currentSchema, err := os.ReadFile(schemaPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("Schema file is missing from %s. Run 'make schema-gen' to create it.", schemaPath)
		}
		t.Fatalf("failed to read committed schema: %v", err)
	}

	if !bytes.Equal(newSchema, currentSchema) {
		t.Errorf("The configuration schema is out of date with domain.Config structure.\n" +
			"Please run 'make schema-gen' locally and commit the updated schema file.")
	}
}
