package model

const (
	PaginationMinItems = 25
	PaginationMaxItems = 50
)

type Pagination struct {
	Page  int32 `json:"page" query:"page"`
	Items int32 `json:"items" query:"items"`
}

func (p *Pagination) NormalizeInput() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Items != PaginationMinItems && p.Items != PaginationMaxItems {
		p.Items = PaginationMinItems
	}

	page := p.Page
	p.Page = (p.Page - 1) * p.Items
	p.Items = page * p.Items
}
