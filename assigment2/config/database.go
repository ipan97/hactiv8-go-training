package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Database struct {
	Driver            string
	Host              string
	Port              int
	User              string
	Password          string
	DatabaseName      string
	MaxIdleConnection int
	MaxOpenConnection int
}

func PGConnect(d *Database) (*gorm.DB, error) {
	args := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.DatabaseName,
	)
	db, err := gorm.Open(d.Driver, args)
	if err != nil {
		return db, err
	}
	db.SetNowFuncOverride(func() time.Time {
		var location, err = time.LoadLocation("Asia/Jakarta")
		if err != nil {
			return time.Now().UTC()
		}
		return time.Now().In(location)
	})
	db.DB().SetMaxIdleConns(d.MaxIdleConnection)
	db.DB().SetMaxOpenConns(d.MaxOpenConnection)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.LogMode(true)
	return db, nil
}
