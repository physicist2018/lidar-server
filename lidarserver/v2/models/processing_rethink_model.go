package models

// ProcessingKletModel struct настройки алгоритма
// расчета отношения рассеяния по методу Клета
type ProcessingKletModel struct {
	Id             string    `rethinkdb:"id" json:"id"`
	ReferenceAlt   float64   `rethinkdb:"reference_alt" json:"reference_alt"`
	ReferenceValue float64   `rethinkdb:"reference_value" json:"reference_value"`
	BackgroundAlt  float64   `rethinkdb:"background_alt" json:"background_alt"`
	Profile        []float64 `rethinkdb:"profile" json:"profile"`
}
