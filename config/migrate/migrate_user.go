package migrate

import "gorm.io/gorm"

func MigrateUser(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id int NOT NULL AUTO_INCREMENT,
		username varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
		password varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci`)
}
