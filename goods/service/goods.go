package service

import (
	"zg6/z326/goods/goods"
	"zg6/z326/goods/model"
)

func Gets(ID []int64) ([]*goods.GoodsInfo, error) {
	newGoods := model.NewGoods()
	var goodsInfoList []*goods.GoodsInfo
	for _, v := range ID {
		gets, err := newGoods.Gets(v)
		if err != nil {
			return nil, err
		}
		goodsInfoList = append(goodsInfoList, myToPb(gets))
	}
	return goodsInfoList, nil
}

func Update(req []*goods.UpdateStock) error {
	newGoods := model.NewGoods()
	for _, v := range req {
		err := newGoods.Update(v.ID, v.Num)
		if err != nil {
			return err
		}
	}
	return nil
}

func myToPb(good *model.Goods) *goods.GoodsInfo {
	return &goods.GoodsInfo{
		ID:    int64(good.ID),
		Name:  good.Name,
		Price: good.Price,
		Stock: good.Stock,
	}
}
