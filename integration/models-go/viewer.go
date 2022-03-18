package models

import "github.com/outcaste-io/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
