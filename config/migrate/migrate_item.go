package migrate

import "gorm.io/gorm"

func MigrateItem(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS item (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
		quantity double NOT NULL DEFAULT '0',
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci`)
}
