package main

import (
	"github.com/maxexq/isekei-shop-api/config"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPosgresDatabase(conf.Database)

	playerMigration(db)
	adminMigration(db)
	itemMigration(db)
	playerCoinMigration(db)
	inventoryMigration(db)
	purchaseHistoryMigration(db)

}

func playerMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func adminMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
