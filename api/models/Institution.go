package models

import "github.com/mindstand/gogm"

type Institution struct {
	// provides required node fields
	gogm.BaseNode

	IntstitutionID string `gogm:"name=inst_id"`
	Name           string `gogm:"name=name"`
	Products       []string
	PrimaryColor   string `gogm:"name=prim_color"`
	Logo           string `gogm:"name=logo"`
	Item           *Item  `gogm:"direction=incoming;relationship=institution"`
}
