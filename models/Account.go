package models

import "github.com/mindstand/gogm"

type Account struct {
	gogm.BaseNode

	AccountID          string `gogm:"name=accnt_id"`
	Name               string `gogm:"name=name"`
	OfficialName       string `gogm:"name=offic_name"`
	Type               string `gogm:"name=type"`
	Subtype            string `gogm:"name=subtype"`
	IntstitutionID     string `gogm:"name=institution_id"`
	VerificationStatus bool
	Owner              *Owner         `gogm:"direction=outgoing;relationship=owner"`
	Balance            *Balance       `gogm:"direction=outgoing;relationship=balance"`
	Item               *Item          `gogm:"direction=incoming;relationship=account"`
	Transactions       []*Transaction `gogm:"direction=outgoing;relationship=transaction"`
}
