package main

import (
	"fmt"

	"github.com/maxexq/isekei-shop-api/config"
	"github.com/maxexq/isekei-shop-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPosgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
