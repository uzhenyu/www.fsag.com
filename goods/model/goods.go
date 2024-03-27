package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name  string
	Price string
	Stock int64
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) Gets(ID int64) (info *Goods, err error) {
	err = WithMysql(func(cli *gorm.DB) error {
		err = cli.Model(g).Where("id = ?", ID).First(&info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (g *Goods) Update(goodsID, stock int64) error {
	err := WithMysql(func(cli *gorm.DB) error {
		err := cli.Model(g).Where("id = ?", goodsID).Update("stock", gorm.Expr("stock + ?", stock)).Error
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
