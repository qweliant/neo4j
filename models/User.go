package models

import "github.com/mindstand/gogm"


type User struct {
	gogm.BaseNode

	UserID int64  `gogm:"name=user_id"`
	Items  []*Item `gogm:"direction=outgoing;relationship=item"`
}