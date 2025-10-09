package psql

import (
	"fmt"

	"github.com/godev-lib/golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionPsql(config *config.Config) *gorm.DB {
	var err error

	dns := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=disable ",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("error connection database: ", err.Error())
		return nil
	}

	fmt.Println("conection database successfully!")
	return db
}
