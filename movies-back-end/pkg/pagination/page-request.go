package pagination

type PageRequest struct {
	PageSize   int  `json:"size,omitempty"`
	PageNumber int  `json:"page,omitempty"`
	Sort       Sort `json:"sort,omitempty"`
}

func PageRequestOf() *PageRequest {
	pageable := &PageRequest{}
	return pageable
}

func (pr *PageRequest) Set(pageNumber, pageSize int) *PageRequest {
	pr.PageNumber = pageNumber
	pr.PageSize = pageSize
	return pr
}

func (pr *PageRequest) GetOffset() int {
	return pr.GetPage() * pr.GetLimit()
}

func (pr *PageRequest) GetLimit() int {
	return pr.PageSize
}

func (pr *PageRequest) GetPage() int {
	return pr.PageNumber
}
