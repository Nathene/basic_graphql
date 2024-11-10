package graph

import (
	"graphql/graph/dbutil"
)

// Resolver is the base resolver struct for gqlgen
type Resolver struct {
	DB dbutil.Database
}
