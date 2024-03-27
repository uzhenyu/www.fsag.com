package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func WithMysql(re func(cli *gorm.DB) error) error {
	viper.SetConfigFile("./etc/order.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	username := viper.GetString("Mysql.Username")
	password := viper.GetString("Mysql.Password")
	host := viper.GetString("Mysql.Host")
	post := viper.GetString("Mysql.Port")
	database := viper.GetString("Mysql.Database")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		post,
		database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, _ := db.DB()
	defer d.Close()
	err = re(db)
	if err != nil {
		return err
	}
	return nil
}
