package enum

type ReportType int

const (
	POST ReportType = 1 + iota
	COMMENT
	REPLY
)

var reportTypeList = []string{
	"POST",
	"COMMENT",
	"REPLY",
}

func (m ReportType) String() string { return reportTypeList[(m - 1)] }

func StringToReportType(reportType string) ReportType {
	var MapEnumStringToReportType = func() map[string]ReportType {
		m := make(map[string]ReportType)
		for i := POST; i <= REPLY; i++ {
			m[i.String()] = i
		}
		return m
	}()

	return MapEnumStringToReportType[reportType]
}
