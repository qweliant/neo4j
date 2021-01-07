package models

import "github.com/mindstand/gogm"

type Balance struct {
	gogm.BaseNode

	Available float64  `gogm:"name=avail"`
	Current   float64  `gogm:"name=current"`
	Limit     float64  `gogm:"name=limit"`
	Currency  string   `gogm:"name=currency"`
	Account   *Account `gogm:"direction=incoming;relationship=balance"`
}