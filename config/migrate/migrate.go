package migrate

import "gorm.io/gorm"

//do check n migrate if table is not exist
func AutoMigrate(db *gorm.DB) {
	MigrateUser(db)
	MigrateItem(db)
	MigrateCart(db)
}
