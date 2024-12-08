package models

type Value interface {
	int | string
}

type ResponseSimpleMap[T Value] map[string]T

type ResponseModel struct {
	Deleted       int      `rethinkdb:"deleted" json:"deleted"`
	Errors        int      `rethinkdb:"errors" json:"errors"`
	GeneratedKeys []string `rethinkdb:"generated_keys,omitempty" json:"generated_keys,omitempty"`
	Inserted      int      `rethinkdb:"inserted" json:"inserted"`
	Replaced      int      `rethinkdb:"replaced" json:"replaced"`
	Skipped       int      `rethinkdb:"skipped" json:"skipped"`
	Unchanged     int      `rethinkdb:"unchanged" json:"unchanged"`
}
