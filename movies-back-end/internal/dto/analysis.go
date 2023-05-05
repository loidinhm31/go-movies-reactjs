package dto

type AnalysisDto struct {
	Year   string   `json:"year,omitempty"`
	Months []string `json:"months,omitempty"`
}

type ResultDto struct {
	Data []*DataDto `json:"data"`
}

type DataDto struct {
	Year  string `json:"year,omitempty"`
	Month string `json:"month,omitempty"`
	Genre string `json:"genre,omitempty"`
	Count int    `json:"count,omitempty"`
}

type RequestData struct {
	Analysis []*AnalysisDto `json:"analysis"`
	Genre    string         `json:"genre,omitempty"`
}
