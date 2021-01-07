package models

import "github.com/mindstand/gogm"

type Address struct {
	gogm.BaseNode

	City       string `gogm:"name=city"`
	Region     string `gogm:"name=region"`
	Street     string `gogm:"name=street"`
	PostalCode string `gogm:"name=zip"`
	Primary    bool
	Owner      *Owner `gogm:"direction=incoming;relationship=address"`
}
