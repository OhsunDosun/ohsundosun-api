package enum

import "database/sql/driver"

type PostType string

const (
	DAILY  PostType = "DAILY"
	LOVE   PostType = "LOVE"
	FRIEND PostType = "FRIEND"
)

func (postType *PostType) Scan(value interface{}) error {
	*postType = PostType(value.([]byte))
	return nil
}

func (postType PostType) Value() (driver.Value, error) {
	return string(postType), nil
}
