package graph

import (
	"graphql/graph/dbutil"
	"graphql/graph/generated"
)

// Resolver is the base resolver struct for gqlgen
type Resolver struct {
	DB dbutil.Database
}

// Mutation connects the mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query connects the query resolver
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
