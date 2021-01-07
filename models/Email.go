package models

import "github.com/mindstand/gogm"

type Email struct {
	gogm.BaseNode

	Address string `gogm:"name=add"`
	Primary bool
	Type    string `gogm:"name=limit"`
	Owner   *Owner `gogm:"direction=incoming;relationship=email"`
}