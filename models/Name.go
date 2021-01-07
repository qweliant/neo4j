package models

import "github.com/mindstand/gogm"

type Name struct {
	gogm.BaseNode

	FullName string `gogm:"name=limit"`
	Owner    *Owner `gogm:"direction=incoming;relationship=name"`
}