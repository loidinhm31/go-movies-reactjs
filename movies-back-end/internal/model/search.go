package model

import "movies-service/pkg/pagination"

const (
	AND = "and"
	OR  = "or"
)

const (
	DATE   = "date"
	STRING = "string"
	NUMBER = "number"
)

type SearchParams struct {
	Filters []*FieldData            `json:"filters"`
	Page    *pagination.PageRequest `json:"page_request"`
}

type FieldData struct {
	Operator  string    `json:"operator"`
	Field     string    `json:"field"`
	TypeValue TypeValue `json:"def"`
}

type TypeValue struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}
