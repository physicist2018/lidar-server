package models

type MolecularModel struct {
	Latitude  float64   `rethinkdb:"latitude" json:"latitude"`
	Longitude float64   `rethinkdb:"longitude" json:"longitude"`
	Altitude  []float64 `rethinkdb:"altitude" json:"altitude"`
	Density   []float64 `rethinkdb:"density" json:"density"`
}
