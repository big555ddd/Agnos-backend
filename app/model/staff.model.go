package model

import (
	"github.com/uptrace/bun"
)

type Staff struct {
	bun.BaseModel `bun:"table:staffs"`

	ID       string `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Username string `bun:"username,unique,notnull" json:"username"`
	Password string `bun:"password,notnull" json:"password"`
	Hospital string `bun:"hospital,notnull" json:"hospital"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
