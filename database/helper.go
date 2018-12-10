package database

// QueryHelper - Will help you to define the sort, limit, asc/desc things
type QueryHelper struct {
	Skip      int    `json:"skip" default:"0"`
	Limit     int    `json:"limit" default:"500"`
	SortBy    string `json:"sort_by" default:"_id"`
	Ascending bool   `json:"ascending" default:"true"`
}
