package adapters

type PaginationRequest struct {
	Limit  int
	Page   int
	Offset int
}
type Pagination struct {
	CurrentPage int  `json:"current_page"`
	TotalPages  int  `json:"total_pages"`
	TotalCount  int  `json:"total_count"`
	HasMore     bool `json:"has_more"`
}
