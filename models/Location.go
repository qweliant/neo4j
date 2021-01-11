package models

import "github.com/mindstand/gogm"

type Location struct {
	gogm.BaseNode

	Address     string `gogm:"name=address"`
	City        string `gogm:"name=city"`
	Region      string `gogm:"name=region"`
	PostalCode  string `gogm:"name=zip"`
	Country     string `gogm:"name=country"`
	StoreNumber string `gogm:"name=store_num"`
	Lat         float64
	Lon         float64
	Transaction *Transaction `gogm:"direction=incoming;relationship=location"`
}
