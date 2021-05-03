package migrations

import (
	"github.com/ipan97/hactiv8-assigment2/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func Up(db *gorm.DB) {
	errCreateOrderSeq := db.Exec("CREATE SEQUENCE IF NOT EXISTS orders_id_seq").Error
	if errCreateOrderSeq != nil {
		log.Error(errCreateOrderSeq)
	}

	errCreateItemSeq := db.Exec("CREATE SEQUENCE IF NOT EXISTS items_id_seq").Error
	if errCreateItemSeq != nil {
		log.Error(errCreateItemSeq)
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})
	errAddFKOrderID := db.Model(&models.Item{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE").Error
	if errAddFKOrderID != nil {
		log.Errorf("Error alter table orders : %v", errAddFKOrderID)
	}
}
