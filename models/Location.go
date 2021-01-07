package models

import "github.com/mindstand/gogm"

type Location struct {
	Address			string `gogm:"name=address"`
	City       		string `gogm:"name=city"`
	Region     		string `gogm:"name=region"`
	Street     		string `gogm:"name=street"`
	PostalCode 		string `gogm:"name=zip"`
	Country 		string `gogm:"name=country"`
	StoreNumber 	string `gogm:"name=store_num"`
	Primary    		bool
	Lat				float64
	Lon				float64
	Transaction     *Owner `gogm:"direction=incoming;relationship=address"`
}