package enum

import "database/sql/driver"

type ReportType string

const (
	POST    ReportType = "POST"
	COMMENT ReportType = "COMMENT"
)

func (reportType *ReportType) Scan(value interface{}) error {
	*reportType = ReportType(value.([]byte))
	return nil
}

func (reportType ReportType) Value() (driver.Value, error) {
	return string(reportType), nil
}
