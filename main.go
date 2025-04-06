package main

import (
	"github.com/maxexq/isekei-shop-api/config"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()

	db := databases.NewPosgresDatabase(conf.Database)

	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}
