package model

type PageRequest struct {
	Paginate int                                 `json:"paginate"`
	PerPage  int                                 `json:"per_page"`
	Page     int                                 `json:"page"`
	Sort     string                              `json:"sort"`
	Search   string                              `json:"search"`
	Filter   []map[string]map[string]interface{} `json:"filter"`
}
