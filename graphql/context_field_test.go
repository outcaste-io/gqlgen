package graphql

import (
	"context"
	"testing"

	"github.com/outcaste-io/gqlparser/v2/ast"
	"github.com/stretchr/testify/require"
)

func TestGetResolverContext(t *testing.T) {
	require.Nil(t, GetFieldContext(context.Background()))

	rc := &FieldContext{}
	require.Equal(t, rc, GetFieldContext(WithFieldContext(context.Background(), rc)))
}

func testContext(sel ast.SelectionSet) context.Context {

	ctx := context.Background()

	rqCtx := &OperationContext{}
	ctx = WithOperationContext(ctx, rqCtx)

	root := &FieldContext{
		Field: CollectedField{
			Selections: sel,
		},
	}
	ctx = WithFieldContext(ctx, root)

	return ctx
}
