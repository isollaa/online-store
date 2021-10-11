package config

import (
	"online-store/config/migrate"

	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Orm *gorm.DB

func init() {
	loadEnvVars()
	connectDB()
	migrate.AutoMigrate(Orm)
}

//load env.json
func loadEnvVars() {
	viper.SetEnvPrefix("evermos")
	errBind := viper.BindEnv("env")

	if errBind != nil {
		panic(fmt.Errorf(errBind.Error()))
	}

	currentDirectory, _ := os.Getwd()

	viper.AddConfigPath(fmt.Sprintf("%s/config/", currentDirectory))
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")

	errRead := viper.ReadInConfig()

	if errRead != nil {
		panic(fmt.Errorf("Fatal error config file: %s", errRead.Error()))
	}
}

func connectDB() {
	conf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.schema"),
	)

	db, err := gorm.Open(
		mysql.Open(conf),
	)
	if err != nil {
		panic("failed to connect to database")
	}

	//show query if debug_mode = true
	if viper.GetBool("database.debug_mode") {
		Orm = db.Debug()
		return
	}

	Orm = db
}
