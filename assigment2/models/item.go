package models

type Item struct {
	ItemID      int    `gorm:"column:item_id;type:int;primary_key;default:nextval('items_id_seq')" json:"lineItemId"`
	ItemCode    string `gorm:"column:item_code;type:varchar(30)" json:"itemCode"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Quantity    int    `gorm:"column:quantity;type:int" json:"quantity"`
	OrderID     int    `gorm:"column:order_id;type:int" json:"-"`
}
