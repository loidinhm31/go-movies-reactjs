package dto

type AnalysisDto struct {
	Year   string   `json:"year,omitempty"`
	Months []string `json:"months,omitempty"`
}

type ResultDto struct {
	Data []*DataDto `json:"data"`
}

type DataDto struct {
	Year       string `json:"year,omitempty"`
	Month      string `json:"month,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeCode   string `json:"type_code,omitempty"`
	Count      int    `json:"count,omitempty"`
	Cumulative int    `json:"cumulative,omitempty"`
}

type RequestData struct {
	Analysis     []*AnalysisDto `json:"analysis"`
	Name         string         `json:"name,omitempty"`
	TypeCode     string         `json:"type_code,omitempty"`
	IsCumulative bool           `json:"isCumulative,omitempty"`
}
