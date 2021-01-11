package models

import (
	
	"github.com/mindstand/gogm"

)


type Transaction struct {
	gogm.BaseNode

	Name	       string    `gogm:"name=name"`
	MerchantName   string    `gogm:"name=merchant"`
	Ammount        float64   `gogm:"name=amount"`
	Currency       string    `gogm:"name=currency"`
	Category       []string 
	Account        *Account  `gogm:"direction=incoming;relationship=transaction"`
	Location       *Location `gogm:"direction=outgoing;relationship=location"`
	PaymentChannel string    `gogm:"name=type"`
	Pending        bool 	 `gogm:"name=pending"`
	
}