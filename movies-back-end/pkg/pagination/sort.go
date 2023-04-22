package pagination

import (
	"fmt"
	"strings"
)

type Direction string

const (
	Undefined Direction = ""
	ASC                 = "asc"
	DESC                = "desc"
)

type Sort struct {
	Orders []*Order `json:"orders,omitempty"`
}

type Order struct {
	Property  string    `json:"property,omitempty"`
	Direction Direction `json:"direction,omitempty"`
}

func (pr *PageRequest) SortByProperties(properties []string) *PageRequest {
	if len(properties) == 0 {
		pr.Sort.Orders = []*Order{
			{
				Property: "id",
			},
		}
	} else {
		for _, p := range properties {
			pr.Sort.Orders = append(pr.Sort.Orders, &Order{
				Property: p,
			})
		}
	}
	return pr
}

func (pr *PageRequest) SortByOrder(orders []*Order) *PageRequest {
	if len(orders) == 0 {
		pr.Sort.Orders = []*Order{
			{
				Property:  "id",
				Direction: ASC,
			},
		}
	} else {
		pr.Sort.Orders = unique(orders)
	}
	return pr
}

func (pr *PageRequest) Ascending() *PageRequest {
	if len(pr.Sort.Orders) > 0 {
		pr.Sort.Orders = unique(pr.Sort.Orders)
		for _, p := range pr.Sort.Orders {
			p.Direction = ASC
		}
	}
	return pr
}

func (pr *PageRequest) Descending() *PageRequest {
	if len(pr.Sort.Orders) > 0 {
		pr.Sort.Orders = unique(pr.Sort.Orders)
		for _, p := range pr.Sort.Orders {
			p.Direction = DESC
		}
	}
	return pr
}

func unique(arr []*Order) []*Order {
	occurred := map[Order]bool{}
	var result []*Order
	for e := range arr {

		// Check if already the mapped
		// variable is set to true or not
		if occurred[*arr[e]] != true {
			occurred[*arr[e]] = true

			// Append to result slice.
			result = append(result, arr[e])
		}
	}
	return result
}

func (pr *PageRequest) GetSort() string {
	var sortStr strings.Builder
	for i, o := range pr.Sort.Orders {
		var s string
		if i != len(pr.Sort.Orders)-1 {
			s = fmt.Sprintf("%s %s, ", o.Property, o.Direction)
		} else {
			s = fmt.Sprintf("%s %s", o.Property, o.Direction)
		}
		sortStr.WriteString(s)
	}
	return sortStr.String()
}
