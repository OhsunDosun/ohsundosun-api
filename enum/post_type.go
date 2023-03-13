package enum

type PostType int

const (
	DAILY PostType = 1 + iota
	LOVE
	FRIEND
)

var postTypeList = []string{
	"daily",
	"love",
	"friend",
}

func (m PostType) String() string { return postTypeList[(m - 1)] }

func StringToPostType(postType string) PostType {
	var MapEnumStringToPostType = func() map[string]PostType {
		m := make(map[string]PostType)
		for i := DAILY; i <= FRIEND; i++ {
			m[i.String()] = i
		}
		return m
	}()

	return MapEnumStringToPostType[postType]
}
