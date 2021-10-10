package migrate

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	MigrateUser(db)
	MigrateItem(db)
	MigrateCart(db)
}
