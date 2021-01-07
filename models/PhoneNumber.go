package models

import "github.com/mindstand/gogm"


type PhoneNumber struct {
	gogm.BaseNode

	Number  string `gogm:"name=number"`
	Primary bool
	Type    string `gogm:"name=type"`
	Owner   *Owner `gogm:"direction=incoming;relationship=number"`
}