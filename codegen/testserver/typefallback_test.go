package testserver

import (
	"context"
	"testing"

	"github.com/outcaste-io/gqlgen/client"
	"github.com/outcaste-io/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestTypeFallback(t *testing.T) {
	resolvers := &Stub{}

	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolvers})))

	resolvers.QueryResolver.Fallback = func(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error) {
		return arg, nil
	}

	t.Run("fallback to string passthrough", func(t *testing.T) {
		var resp struct {
			Fallback string
		}
		c.MustPost(`query { fallback(arg: A) }`, &resp)
		require.Equal(t, "A", resp.Fallback)
	})
}
