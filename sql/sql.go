package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hotSearch/config"
	"log"
)

func main() {
	TestConnection()
}

func TestConnection() {

	address := config.Conf.Get("sql.address").(string)
	user := config.Conf.Get("sql.user").(string)
	passwd := config.Conf.Get("sql.passwd").(string)
	driver := config.Conf.Get("sql.driver").(string)
	_SQLConnection(driver, user, passwd, address)
}

func _SQLConnection(driver, user, passwd, address string) {
	db, err := sql.Open(driver, user+":"+passwd+"@tcp("+address+")/mysql")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}
