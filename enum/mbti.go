package enum

type MBTI int

const (
	INTJ MBTI = 1 + iota
	INTP
	ENTJ
	ENTP

	INFJ
	INFP
	ENFJ
	ENFP

	ISFJ
	ISTJ
	ESFJ
	ESTJ

	ISFP
	ISTP
	ESFP
	ESTP
)

var mbtiList = []string{
	"INTJ",
	"INTP",
	"ENTJ",
	"ENTP",

	"INFJ",
	"INFP",
	"ENFJ",
	"ENFP",

	"ISFJ",
	"ISTJ",
	"ESFJ",
	"ESTJ",

	"ISFP",
	"ISTP",
	"ESFP",
	"ESTP",
}

func (m MBTI) String() string { return mbtiList[(m - 1)] }

func StringToMBTI(mbti string) MBTI {
	var MapEnumStringToMBTI = func() map[string]MBTI {
		m := make(map[string]MBTI)
		for i := INTJ; i <= ESTP; i++ {
			m[i.String()] = i
		}
		return m
	}()

	return MapEnumStringToMBTI[mbti]
}
