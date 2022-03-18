package config

import (
	"go/types"
	"testing"

	"github.com/outcaste-io/gqlgen/internal/code"

	"github.com/outcaste-io/gqlparser/v2"
	"github.com/outcaste-io/gqlparser/v2/ast"
	"github.com/stretchr/testify/require"
)

func TestBindingToInvalid(t *testing.T) {
	binder, schema := createBinder(Config{})
	_, err := binder.TypeReference(schema.Query.Fields.ForName("messages").Type, &types.Basic{})
	require.EqualError(t, err, "Message has an invalid type")
}

func TestSlicePointerBinding(t *testing.T) {
	t.Run("without OmitSliceElementPointers", func(t *testing.T) {
		binder, schema := createBinder(Config{
			OmitSliceElementPointers: false,
		})

		ta, err := binder.TypeReference(schema.Query.Fields.ForName("messages").Type, nil)
		if err != nil {
			panic(err)
		}

		require.Equal(t, ta.GO.String(), "[]*github.com/outcaste-io/gqlgen/example/chat.Message")
	})

	t.Run("with OmitSliceElementPointers", func(t *testing.T) {
		binder, schema := createBinder(Config{
			OmitSliceElementPointers: true,
		})

		ta, err := binder.TypeReference(schema.Query.Fields.ForName("messages").Type, nil)
		if err != nil {
			panic(err)
		}

		require.Equal(t, ta.GO.String(), "[]github.com/outcaste-io/gqlgen/example/chat.Message")
	})
}

func createBinder(cfg Config) (*Binder, *ast.Schema) {
	cfg.Models = TypeMap{
		"Message": TypeMapEntry{
			Model: []string{"github.com/outcaste-io/gqlgen/example/chat.Message"},
		},
	}
	cfg.Packages = &code.Packages{}

	cfg.Schema = gqlparser.MustLoadSchema(&ast.Source{Name: "TestAutobinding.schema", Input: `
		type Message { id: ID }

		type Query {
			messages: [Message!]!
		}
	`})

	b := cfg.NewBinder()

	return b, cfg.Schema
}
