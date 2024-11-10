package graph

import (
	"graphql/graph/dbutil"
	"graphql/graph/generated"
)

// Resolver struct to hold the database instance.
type Resolver struct {
	DB dbutil.Database
}

// Mutation returns the mutation resolver.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns the query resolver.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
