package database

import (
	"database/sql"
	"fmt"
	"restaurant-management/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBInstance() (*sql.DB, error) {
	// Yapılandırma dosyasını yükle (db için gerekli parametreleri al)
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %v", err)
	}

	// MySQL DSN oluştur
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.MySQLUsername, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDBName)

	// MySQL veritabanına bağlantı aç
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Bağlantı havuz parametrelerini ayarla
	db.SetConnMaxLifetime(10 * time.Minute) // Bağlantı ömrünü sınırla
	db.SetMaxOpenConns(10)                  // Max açık bağlantı sayısını sınırla
	db.SetMaxIdleConns(5)                   // Veritabanı havuzundaki boş bağlantı sayısını sınırla

	// Veritabanına ping atarak bağlantıyı test et
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Printf("Connected to MySQL database: %s\n", cfg.MySQLDBName)

	return db, nil
}
