package datatype

import "github.com/graphql-go/graphql"

var (
	Employee = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Employee",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.Field{
					Type: graphql.Int,
				},
				"position": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: Address,
				},
			},
		},
	)

	Address = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Address",
			Fields: graphql.Fields{
				"ward": &graphql.Field{
					Type: graphql.String,
				},
				"district": &graphql.Field{
					Type: graphql.String,
				},
				"province": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)
