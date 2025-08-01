package model

import "github.com/graph-gophers/graphql-go"

// Order é só um DTO — o binding será feito pelos
// métodos no orderResolver, não por campos diretamente.
type Order struct {
	ID         graphql.ID
	Price      float64
	Tax        float64
	FinalPrice float64
}
