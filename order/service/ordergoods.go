package service

import (
	"gorm.io/gorm"
	"zg6/z326/order/model"
	"zg6/z326/order/order"
)

func CreateOrderGoods(info []*order.OrderGoodsInfo) ([]*order.OrderGoodsInfo, error) {
	newOrderGoods := model.NewOrderGoods()
	var newOrderGoodsList []*order.OrderGoodsInfo
	for _, v := range info {
		goods, err := newOrderGoods.CreateOrderGoods(pbToMys(v))
		if err != nil {
			return nil, err
		}
		newOrderGoodsList = append(newOrderGoodsList, myToPbs(goods))
	}
	return newOrderGoodsList, nil
}

func GetOrderGoods(orderID int64) ([]*order.OrderGoodsInfo, error) {
	newOrderGoods := model.NewOrderGoods()
	goods, err := newOrderGoods.GetOrderGoods(orderID)
	if err != nil {
		return nil, err
	}
	var newOrderGoodsList []*order.OrderGoodsInfo
	for _, v := range goods {
		newOrderGoodsList = append(newOrderGoodsList, myToPbs(v))
	}
	return newOrderGoodsList, nil
}

func myToPbs(goods *model.OrderGoods) *order.OrderGoodsInfo {
	return &order.OrderGoodsInfo{
		ID:        int64(goods.ID),
		OrderID:   goods.OrderID,
		GoodsID:   goods.GoodsID,
		UnitPrice: goods.UnitPrice,
		GoodsName: goods.GoodsName,
		Num:       goods.Num,
	}
}

func pbToMys(info *order.OrderGoodsInfo) *model.OrderGoods {
	return &model.OrderGoods{
		Model: gorm.Model{
			ID: uint(info.ID),
		},
		OrderID:   info.OrderID,
		GoodsID:   info.GoodsID,
		UnitPrice: info.UnitPrice,
		GoodsName: info.GoodsName,
		Num:       info.Num,
	}
}
