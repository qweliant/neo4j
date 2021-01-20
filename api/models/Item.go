package models

import "github.com/mindstand/gogm"

type Item struct {
	gogm.BaseNode

	User        *User        `gogm:"direction=incoming;relationship=item"`
	Institution *Institution `gogm:"direction=outgoing;relationship=institution"`
	Accounts     []*Account   `gogm:"direction=outgoing;relationship=account"`
}