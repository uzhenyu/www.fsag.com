package service

import (
	"gorm.io/gorm"
	"zg6/z326/order/model"
	"zg6/z326/order/order"
)

func CreateOrder(info *order.OrderInfo) (*order.OrderInfo, error) {
	newOrder := model.NewOrder()
	createOrder, err := newOrder.CreateOrder(pbToMy(info))
	if err != nil {
		return nil, err
	}
	return myToPb(createOrder), nil
}

func GetOrder(ID int64) (*order.OrderInfo, error) {
	newOrder := model.NewOrder()
	getOrder, err := newOrder.GetOrder(ID)
	if err != nil {
		return nil, err
	}
	return myToPb(getOrder), nil
}

func UpdateByOrderNO(orderNO string, state int64) error {
	newOrder := model.NewOrder()
	err := newOrder.UpdateByOrderNO(orderNO, state)
	if err != nil {
		return err
	}
	return nil
}

func myToPb(orders *model.Order) *order.OrderInfo {
	return &order.OrderInfo{
		ID:      int64(orders.ID),
		UserID:  orders.UserID,
		OrderNO: orders.OrderNO,
		Amount:  orders.Amount,
		State:   orders.State,
	}
}

func pbToMy(info *order.OrderInfo) *model.Order {
	return &model.Order{
		Model: gorm.Model{
			ID: uint(info.ID),
		},
		UserID:  info.UserID,
		OrderNO: info.OrderNO,
		Amount:  info.Amount,
		State:   info.State,
	}
}
