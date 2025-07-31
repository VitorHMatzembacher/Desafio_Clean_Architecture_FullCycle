package graph

import _ "embed"

//go:embed schema.graphqls
var schemaSDL string

// MustLoadSchema retorna o conteúdo do seu schema.graphqls
func MustLoadSchema() string {
	return schemaSDL
}
