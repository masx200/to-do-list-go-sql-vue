package models

type QueryParameters struct {
	ID        uint   ` form:"id"`
	Limit     int    ` form:"limit"`
	Page      int    ` form:"page"`
	Order     string ` form:"order"`
	Direction string ` form:"direction"`
}
