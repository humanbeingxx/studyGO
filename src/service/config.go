package service

import (
	"database/sql"
	"fmt"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Name  string
	Value string
}

func open() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("DB INITIALIZE FAILED !!!")
	}
	return db
}

func ReadConfig() []Config {
	db := open()
	defer db.Close()
	rows, err := db.Query("select name, value from tb_config")
	if err != nil {
		return nil
	}
	rows.Columns()
	configs := make([]Config, 0)

	for rows.Next() {
		var conf Config
		var name string
		var value string
		rows.Scan(&name, &value)
		conf.Name = name
		conf.Value = value
		configs = append(configs, conf)
	}

	return configs
}

func GetConfig(name string) string {
	db := open()
	defer db.Close()

	var value string
	err := db.QueryRow("select value from tb_config where name = ?", name).Scan(&value)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

func AddConfig(name string, value string) int {
	db := open()
	defer db.Close()
	stmt, _ := db.Prepare(`insert into tb_config (name,value) values (?,?) on duplicate key update value = ?`)
	ret, err := stmt.Exec(name, value, value)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if affected, err := ret.RowsAffected(); err == nil {
		return int(affected)
		// result := *((*int)(unsafe.Pointer(&affected)))
		// return result
	}
	fmt.Println(ret.RowsAffected())
	return -1
}
