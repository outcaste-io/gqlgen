package cmd

import (
	// Import and ignore the ambient imports listed below so dependency managers
	// don't prune unused code for us. Both lists should be kept in sync.
	_ "github.com/outcaste-io/gqlgen/graphql"
	_ "github.com/outcaste-io/gqlgen/graphql/introspection"
	_ "github.com/outcaste-io/gqlparser/v2"
	_ "github.com/outcaste-io/gqlparser/v2/ast"
)
