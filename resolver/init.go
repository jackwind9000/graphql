package resolver

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jackwind9000/graphql/resolver/mutation"
	"github.com/jackwind9000/graphql/resolver/query"
)

func Init() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.Query(),
			Mutation: mutation.Mutation(),
		},
	)

	if err != nil {
		log.Fatalf("schema: Init: %v", err)
		return nil
	}

	return &schema
}
