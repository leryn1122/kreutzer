package entity

import (
	"github.com/leryn1122/kreutzer/v2/infra/id"
	"time"
)

type User struct {
	UID       id.Identifier
	Password  string
	CreatedAt time.Time
}

type Users []User
