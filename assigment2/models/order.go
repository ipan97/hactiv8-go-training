package models

import "time"

type Order struct {
	OrderID      int       `gorm:"column:order_id;type:int;primary_key;default:nextval('orders_id_seq')" json:"orderId"`
	CustomerName string    `gorm:"column:customer_name;type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"column:ordered_at" json:"orderedAt"`
	Items        []Item    `gorm:"foreignkey:OrderID" json:"items"`
}
