package models

import "time"

type Order struct {
	OrderID      int       `gorm:"column:order_id;type:int;primary_key;default:nextval('orders_id_seq')"`
	CustomerName string    `gorm:"column:customer_name;type:varchar(50)"`
	OrderedAt    time.Time `gorm:"column:ordered_at"`
	Items        []Item    `gorm:"foreignkey:OrderID"`
}
