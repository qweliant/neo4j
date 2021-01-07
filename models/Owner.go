package models

import "github.com/mindstand/gogm"

type Owner struct {
	gogm.BaseNode

	Account      *Account       `gogm:"direction=incoming;relationship=owner"`
	Names        []*Name        `gogm:"direction=outgoing;relationship=name"`
	PhoneNumbers []*PhoneNumber `gogm:"direction=outgoing;relationship=number"`
	Emails       []*Email       `gogm:"direction=outgoing;relationship=email"`
	Addresses    []*Address     `gogm:"direction=outgoing;relationship=address"`
}