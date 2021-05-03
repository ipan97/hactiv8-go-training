package models

type Item struct {
	ItemID      int    `gorm:"column:item_id;type:int;primary_key;default:nextval('items_id_seq')"`
	ItemCode    string `gorm:"column:item_code;type:varchar(30)"`
	Description string `gorm:"column:description;type:text"`
	Quantity    int    `gorm:"column:quantity;type:int"`
	OrderID     int    `gorm:"column:order_id;type:int"`
}
