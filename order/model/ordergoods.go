package model

import "gorm.io/gorm"

type OrderGoods struct {
	gorm.Model
	OrderID   int64
	GoodsID   int64
	UnitPrice string
	GoodsName string
	Num       int64
}

func NewOrderGoods() *OrderGoods {
	return new(OrderGoods)
}

func (o *OrderGoods) CreateOrderGoods(goods *OrderGoods) (info *OrderGoods, err error) {
	err = WithMysql(func(cli *gorm.DB) error {
		err = cli.Create(goods).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (o *OrderGoods) GetOrderGoods(orderID int64) (info []*OrderGoods, err error) {
	err = WithMysql(func(cli *gorm.DB) error {
		err = cli.Model(o).Where("order_id = ?", orderID).Find(&info).Error
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
