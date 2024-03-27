package model

import "gorm.io/gorm"

func AutoTable() error {
	err := WithMysql(func(cli *gorm.DB) error {
		err := cli.AutoMigrate(new(Goods))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
