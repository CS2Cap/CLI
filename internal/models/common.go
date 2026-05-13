package models

type PaginationMeta struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Total      int    `json:"total"`
	HasNext    bool   `json:"has_next"`
	HasPrev    bool   `json:"has_prev"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type ErrorResponse struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func (e ErrorResponse) Error() string {
	return e.Code + ": " + e.Detail
}

type ValidationErrorItem struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

type ValidationErrorResponse struct {
	Code   string                `json:"code,omitempty"`
	Detail []ValidationErrorItem `json:"detail"`
}
