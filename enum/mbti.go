package enum

import "database/sql/driver"

type MBTI string

const (
	INTJ MBTI = "INTJ"
	INTP MBTI = "INTP"
	ENTJ MBTI = "ENTJ"
	ENTP MBTI = "ENTP"

	INFJ MBTI = "INFJ"
	INFP MBTI = "INFP"
	ENFJ MBTI = "ENFJ"
	ENFP MBTI = "ENFP"

	ISFJ MBTI = "ISFJ"
	ISTJ MBTI = "ISTJ"
	ESFJ MBTI = "ESFJ"
	ESTJ MBTI = "ESTJ"

	ISFP MBTI = "ISFP"
	ISTP MBTI = "ISTP"
	ESFP MBTI = "ESFP"
	ESTP MBTI = "ESTP"
)

func (mbti *MBTI) Scan(value interface{}) error {
	*mbti = MBTI(value.([]byte))
	return nil
}

func (mbti MBTI) Value() (driver.Value, error) {
	return string(mbti), nil
}
