package model

const (
	PaginationMinItems = 25
	PaginationMaxItems = 50
)

func Pagination(page, items int32) (int32, int32) {
	if page < 1 {
		page = 1
	}

	if items != PaginationMinItems && items != PaginationMaxItems {
		items = PaginationMinItems
	}

	p := page
	page = (page - 1) * items
	items = p * items
	return page, items
}
