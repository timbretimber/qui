package schema

import (
	"encoding/json"
	"io"

	"github.com/autobrr/qui/internal/domain"
	"github.com/invopop/jsonschema"
)

const SiteConfigSchemaPath = "static/config.schema.json"
const ConfigSchemaPath = "documentation/" + SiteConfigSchemaPath

func ExportSchema(w io.Writer) error {
	reflector := &jsonschema.Reflector{
		FieldNameTag:               "toml",
		RequiredFromJSONSchemaTags: true,
	}

	schema := reflector.Reflect(&domain.Config{})
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	return encoder.Encode(schema)
}
