package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL sürücüsü için import
	"log"
	"restaurant-management/pkg/config"
	"time"
)

type DB struct {
	DB *sql.DB
}

func DBInstance() (*DB, error) {
	// Load configuration (config'den db için gerekli parametreleri alacağız)
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %v", err)
	}

	// MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MySQLUsername, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Set connection pool parameters
	db.SetConnMaxLifetime(10 * time.Minute) // adı üstünde bağlantıların uzun süreli olarak açık kalmasını önler, hafızayı/kaynakları tüketmesini önler.
	db.SetMaxOpenConns(10)                  // max 10 bağlantı açabilir
	db.SetMaxIdleConns(5)                   // veritabanı havuzunda aynı anda açık tutulacak maksimum boş (kullanılmayan) bağlantı sayısını belirler.

	// Ping the database to ensure a successful connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Printf("Connected to MySQL database: %s\n", cfg.MySQLDBName)

	return &DB{DB: db}, nil
}

// TODO : böyle bir yapıya gerçekten ihtiyaç var mı? TEKRAR GÖZDEN GEÇİR!
func (db *DB) OpenTable(tableName string) *sql.DB {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatalf("Error querying table %s: %v", tableName, err)
	}
	defer rows.Close()

	return db.DB

}
