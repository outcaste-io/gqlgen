// Code generated by github.com/outcaste-io/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/outcaste-io/gqlgen/example/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location"`
}
