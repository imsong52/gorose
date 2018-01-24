package main

import (
	"./config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

func main() {
	db, err := gorose.Open(config.DbConfig, "mysql_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	defer db.Close()

	where := map[string]interface{}{
		"id": 17,
	}
	res := db.Table("users").Where(where).Delete()
	fmt.Println(db.LastSql())
	fmt.Println(res)

}
