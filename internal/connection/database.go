package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-rest/internal/config"
)

func GetDatabase(conf config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port, // Port sudah string dari model
		conf.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err // Jangan panic, kembalikan error
	}

	err = db.Ping()
	if err != nil {
		return nil, err // Kembalikan error jika ping gagal
	}

	// Hapus 'defer db.Close()' dari sini
	// Hapus juga fmt.Println agar tidak 'spam' terminal

	return db, nil
}
