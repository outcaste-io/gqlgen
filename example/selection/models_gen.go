// Code generated by github.com/outcaste-io/gqlgen, DO NOT EDIT.

package selection

import (
	"time"
)

type Event interface {
	IsEvent()
}

type Like struct {
	Reaction  string    `json:"reaction"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection"`
	Collected []string  `json:"collected"`
}

func (Like) IsEvent() {}

type Post struct {
	Message   string    `json:"message"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection"`
	Collected []string  `json:"collected"`
}

func (Post) IsEvent() {}
