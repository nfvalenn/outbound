package model

import "gorm.io/gorm"

type Outbound struct {
	gorm.Model
	ID int `json:"id"`
	Products string `json:"products"`
	Date int `json:"date"`
	Qty int `json:"qty"`
	Location string `json:"location"`
	Building string `json:"building"`
	Room string `json:"room"`
	Floor string `json:"floor"`
	Area string `json:"area"`
	Rack int `json:"rack"`
	RackLevel int `json:"rack_level"`
}