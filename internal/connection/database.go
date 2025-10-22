package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-rest/internal/config"
	// "strconv"
)
func GetDatabase(conf config.Database)  (*sql.DB, error){

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port, // Port (int) harus diubah ke string
		conf.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Berhasil terhubung ke database MySQL!")
	return db, nil
}
