package testserver

import (
	"context"
	"testing"

	"github.com/outcaste-io/gqlgen/client"
	"github.com/outcaste-io/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestNullBubbling(t *testing.T) {
	resolvers := &Stub{}
	resolvers.QueryResolver.Valid = func(ctx context.Context) (s string, e error) {
		return "Ok", nil
	}

	resolvers.QueryResolver.Errors = func(ctx context.Context) (errors *Errors, e error) {
		return &Errors{}, nil
	}
	resolvers.QueryResolver.ErrorBubble = func(ctx context.Context) (i *Error, e error) {
		return &Error{ID: "E1234"}, nil
	}

	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolvers})))

	t.Run("when function errors on non required field", func(t *testing.T) {
		var resp struct {
			Valid       string
			ErrorBubble *struct {
				Id                      string
				ErrorOnNonRequiredField *string
			}
		}
		err := c.Post(`query { valid, errorBubble { id, errorOnNonRequiredField } }`, &resp)

		require.EqualError(t, err, `[{"message":"boom","path":["errorBubble","errorOnNonRequiredField"]}]`)
		require.Equal(t, "E1234", resp.ErrorBubble.Id)
		require.Nil(t, resp.ErrorBubble.ErrorOnNonRequiredField)
		require.Equal(t, "Ok", resp.Valid)
	})

	t.Run("when function errors", func(t *testing.T) {
		var resp struct {
			Valid       string
			ErrorBubble *struct {
				NilOnRequiredField string
			}
		}
		err := c.Post(`query { valid, errorBubble { id, errorOnRequiredField } }`, &resp)

		require.EqualError(t, err, `[{"message":"boom","path":["errorBubble","errorOnRequiredField"]}]`)
		require.Nil(t, resp.ErrorBubble)
		require.Equal(t, "Ok", resp.Valid)
	})

	t.Run("when user returns null on required field", func(t *testing.T) {
		var resp struct {
			Valid       string
			ErrorBubble *struct {
				NilOnRequiredField string
			}
		}
		err := c.Post(`query { valid, errorBubble { id, nilOnRequiredField } }`, &resp)

		require.EqualError(t, err, `[{"message":"must not be null","path":["errorBubble","nilOnRequiredField"]}]`)
		require.Nil(t, resp.ErrorBubble)
		require.Equal(t, "Ok", resp.Valid)
	})

	t.Run("null args", func(t *testing.T) {
		var resp struct {
			NullableArg *string
		}
		resolvers.QueryResolver.NullableArg = func(ctx context.Context, arg *int) (i *string, e error) {
			v := "Ok"
			return &v, nil
		}

		err := c.Post(`query { nullableArg(arg: null) }`, &resp)
		require.Nil(t, err)
		require.Equal(t, "Ok", *resp.NullableArg)
	})

	t.Run("concurrent null detection", func(t *testing.T) {
		var resp interface{}
		resolvers.ErrorsResolver.A = func(ctx context.Context, obj *Errors) (i *Error, e error) { return nil, nil }
		resolvers.ErrorsResolver.B = func(ctx context.Context, obj *Errors) (i *Error, e error) { return nil, nil }
		resolvers.ErrorsResolver.C = func(ctx context.Context, obj *Errors) (i *Error, e error) { return nil, nil }
		resolvers.ErrorsResolver.D = func(ctx context.Context, obj *Errors) (i *Error, e error) { return nil, nil }
		resolvers.ErrorsResolver.E = func(ctx context.Context, obj *Errors) (i *Error, e error) { return nil, nil }

		err := c.Post(`{ errors { 
			a { id },
			b { id },
			c { id },
			d { id },
			e { id },
		} }`, &resp)

		require.Error(t, err)
		require.Contains(t, err.Error(), "must not be null")
	})
}
