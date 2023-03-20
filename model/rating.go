package model

import (
	"database/sql"
)

type Rating struct {
	Key       string         `json:"key"`
	UserKey   string         `json:"userKey"`
	Rating    float32        `json:"rating"`
	Feedback  sql.NullString `json:"feedback"`
	CreatedAt int64          `json:"createdAt"`
}
