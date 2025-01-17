// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

type Battery struct {
	BatteryID        string
	BatteryName      string
	Brand            string
	Capacity         int64
	PurchaseDate     string
	PurchasePrice    string
	PurchaseCurrency string
	CompetitionClass string
	Status           string
	Created          string
	Updated          string
}

type BatteryHistory struct {
	ID        string
	BatteryID string
	Timestamp string
	Comment   string
	Created   string
	Updated   string
}
