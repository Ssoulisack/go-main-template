package middleware

type PageQuery struct {
	Page       int         `json:"page,omitempty" query:"page"`
	Limit      int         `json:"limit,omitempty" query:"limit"`
	TotalPages int         `json:"totalPages"`
	TotalRows  int64       `json:"totalRows"`
	Rows       interface{} `json:"rows"`
	Search     string      `json:"search,omitempty" query:"search"`
}
