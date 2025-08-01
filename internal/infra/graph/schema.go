package graph

import _ "embed"

//go:embed schema.graphqls
var sdl string

// MustLoadSchema retorna o SDL embutido de schema.graphqls
func MustLoadSchema() string {
	return sdl
}
