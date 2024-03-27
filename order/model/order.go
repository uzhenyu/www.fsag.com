package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  int64
	OrderNO string
	Amount  string
	State   int64
}

func NewOrder() *Order {
	return new(Order)
}

func (o *Order) CreateOrder(order *Order) (info *Order, err error) {
	err = WithMysql(func(cli *gorm.DB) error {
		err = cli.Create(order).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) GetOrder(ID int64) (info *Order, err error) {
	err = WithMysql(func(cli *gorm.DB) error {
		err = cli.Model(o).Where("id = ?", ID).First(&info).Error
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

func (o *Order) UpdateByOrderNO(orderNO string, state int64) error {
	err := WithMysql(func(cli *gorm.DB) error {
		err := cli.Model(o).Where("order_no = ?", orderNO).Update("state", state).Error
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
