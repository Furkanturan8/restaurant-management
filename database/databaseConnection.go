package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restaurant-management/models"
	"restaurant-management/pkg/config"
)

func DBInstance() (*gorm.DB, error) {

	// Yapılandırma dosyasını yükle (db için gerekli parametreleri alıyoruz)
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %v", err)
	}

	// MySQL DSN oluştur
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQLUsername, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDBName)

	// MySQL veritabanına bağlantı aç
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Automigrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Table{},
		&models.Menu{},
		&models.Invoice{},
		&models.Order{},
		&models.OrderItem{},
		&models.Food{},
	)
	if err != nil {
		return nil, fmt.Errorf("AutoMigrate failed: %v", err)
	}

	fmt.Printf("Connected to MySQL database: %s\n", cfg.MySQLDBName)

	return db, nil
}
