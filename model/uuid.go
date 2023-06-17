package model

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UUID uuid.UUID

// StringToUUID -> parse string to MYTYPE
func StringToUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	return UUID(id), err
}

// String -> String Representation of Binary16
func (id UUID) String() string {
	return uuid.UUID(id).String()
}

// GormDataType -> sets type to binary(16)
func (id UUID) GormDataType() string {
	return "binary(16)"
}

func (id UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(id)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (id *UUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*id = UUID(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (id *UUID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*id = UUID(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (id UUID) Value() (driver.Value, error) {
	return uuid.UUID(id).MarshalBinary()
}
