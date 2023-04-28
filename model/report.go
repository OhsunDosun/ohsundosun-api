package model

import "ohsundosun-api/enum"

type Report struct {
	Key       string          `json:"key"`
	Type      enum.ReportType `json:"type"`
	TargetKey string          `json:"targetKey"`
	CreatedAt int64           `json:"createdAt"`
}
