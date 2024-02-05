package model

type PageRequest struct {
	Filters  []map[string]map[string]interface{} `json:"filter"`
	Sort     string                              `json:"sort"`
	Search   string                              `json:"search"`
	Paginate int                                 `json:"paginate"`
	PerPage  int                                 `json:"per_page"`
	Page     int                                 `json:"page"`
}
