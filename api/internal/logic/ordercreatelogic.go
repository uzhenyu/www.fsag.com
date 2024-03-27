package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"zg6/z326/api/internal/svc"
	"zg6/z326/api/internal/types"
	"zg6/z326/api/pkg"
	"zg6/z326/goods/goods"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//func (l *OrderCreateLogic) getGoodsMap(req *types.OrderCreateRequest) (map[int64]*goods.GoodsInfo, error) {
//	var goodsIDs []int64
//	for _, v := range req.Goods {
//		goodsIDs = append(goodsIDs, v.GoodsID)
//	}
//	ds, err := l.svcCtx.GoodsRpc.GetGoodsIDs(l.ctx, &goods.GetGoodsIDsRequest{
//		IDs: goodsIDs,
//	})
//	if err != nil {
//		return nil, err
//	}
//	goodMap := make(map[int64]*goods.GoodsInfo)
//	for _, v := range ds.Infos {
//		goodMap[v.ID] = v
//	}
//	return goodMap, err
//}
//
//func (l *OrderCreateLogic) getGoodAmount(req *types.OrderCreateRequest) (decimal.Decimal, error) {
//	goodsMap, err := l.getGoodsMap(req)
//	if err != nil {
//		return decimal.Decimal{}, err
//	}
//	amount := decimal.NewFromInt(0)
//	for _, v := range req.Goods {
//		goodsInfo, ok := goodsMap[v.GoodsID]
//		if !ok {
//			return decimal.Decimal{}, fmt.Errorf("商品不存在")
//		}
//		unitPrice, err := decimal.NewFromString(goodsInfo.Price)
//		if err != nil {
//			return decimal.Decimal{}, err
//		}
//		amount = amount.Add(unitPrice.Mul(decimal.NewFromInt(v.Num)))
//	}
//	return amount, nil
//}
//
//func (l *OrderCreateLogic) createOrder(orderNO, amount string, in *types.OrderCreateRequest) error {
//	createOrder, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderRequest{
//		Info: &order.OrderInfo{
//			UserID:  in.UserID,
//			OrderNO: orderNO,
//			Amount:  amount,
//			State:   1,
//		},
//	})
//	if err != nil {
//		return err
//	}
//	goodsMap, err := l.getGoodsMap(in)
//	if err != nil {
//		return err
//	}
//	var req []*order.OrderGoodsInfo
//	for _, v := range in.Goods {
//		goodsInfo, ok := goodsMap[v.GoodsID]
//		if !ok {
//			return fmt.Errorf("商品不存在")
//		}
//		req = append(req, &order.OrderGoodsInfo{
//			OrderID:   createOrder.Info.ID,
//			GoodsID:   v.GoodsID,
//			UnitPrice: goodsInfo.Price,
//			GoodsName: goodsInfo.Name,
//			Num:       v.Num,
//		})
//	}
//	_, err = l.svcCtx.OrderRpc.CreateOrderGoods(l.ctx, &order.CreateOrderGoodsRequest{
//		Info: req,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (l *OrderCreateLogic) updateStock(in *types.OrderCreateRequest) error {
//	goodsMap, err := l.getGoodsMap(in)
//	if err != nil {
//		return err
//	}
//	var req []*goods.UpdateStock
//	for _, v := range in.Goods {
//		goodInfo, ok := goodsMap[v.GoodsID]
//		if !ok {
//			return fmt.Errorf("商品不存在")
//		}
//		if goodInfo.Stock-v.Num < 0 {
//			return fmt.Errorf("库存不足")
//		}
//		req = append(req, &goods.UpdateStock{
//			ID:  goodInfo.ID,
//			Num: -v.Num,
//		})
//	}
//	_, err = l.svcCtx.GoodsRpc.UpdateStocks(l.ctx, &goods.UpdateStocksRequest{Infos: req})
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (l *OrderCreateLogic) getGoodsMap(req *types.OrderCreateRequest) (map[int64]*goods.GoodsInfo, error) {
	var goodIDs []int64
	for _, v := range req.Goods {
		goodIDs = append(goodIDs, v.GoodsID)
	}
	ds, err := l.svcCtx.GoodsRpc.GetGoodsIDs(l.ctx, &goods.GetGoodsIDsRequest{IDs: goodIDs})
	if err != nil {
		return nil, err
	}
	goodsMap := make(map[int64]*goods.GoodsInfo)
	for _, v := range ds.Infos {
		goodsMap[v.ID] = v
	}
	return goodsMap, nil
}

func (l *OrderCreateLogic) getGoodsAmount(req *types.OrderCreateRequest) (decimal.Decimal, error) {
	goodsMap, err := l.getGoodsMap(req)
	if err != nil {
		return decimal.Decimal{}, err
	}
	amount := decimal.NewFromInt(0)
	for _, v := range req.Goods {
		goodsInfo, ok := goodsMap[v.GoodsID]
		if !ok {
			return decimal.Decimal{}, fmt.Errorf("商品不存在")
		}
		unitPrice, err := decimal.NewFromString(goodsInfo.Price)
		if err != nil {
			return decimal.Decimal{}, err
		}
		amount = amount.Add(unitPrice.Mul(decimal.NewFromInt(v.Num)))
	}
	return amount, nil
}

func (l *OrderCreateLogic) createOrder(orderNo, amount string, in *types.OrderCreateRequest) error {
	createOrder, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderRequest{
		Info: &order.OrderInfo{
			UserID:  in.UserID,
			OrderNO: orderNo,
			Amount:  amount,
			State:   1,
		},
	})
	if err != nil {
		return err
	}
	goodsMap, err := l.getGoodsMap(in)
	if err != nil {
		return err
	}
	var req []*order.OrderGoodsInfo
	for _, v := range in.Goods {
		goodsInfo, ok := goodsMap[v.GoodsID]
		if !ok {
			return fmt.Errorf("商品不存在")
		}
		req = append(req, &order.OrderGoodsInfo{
			OrderID:   createOrder.Info.ID,
			GoodsID:   v.GoodsID,
			UnitPrice: goodsInfo.Price,
			GoodsName: goodsInfo.Name,
			Num:       v.Num,
		})
	}
	_, err = l.svcCtx.OrderRpc.CreateOrderGoods(l.ctx, &order.CreateOrderGoodsRequest{Info: req})
	if err != nil {
		return err
	}
	return nil
}

func (l *OrderCreateLogic) updateStock(in *types.OrderCreateRequest) error {
	goodsMap, err := l.getGoodsMap(in)
	if err != nil {
		return err
	}
	var req []*goods.UpdateStock
	for _, v := range in.Goods {
		goodsInfo, ok := goodsMap[v.GoodsID]
		if !ok {
			return fmt.Errorf("商品不存在")
		}
		if goodsInfo.Stock-v.Num < 0 {
			return fmt.Errorf("库存不足")
		}
		req = append(req, &goods.UpdateStock{
			ID:  goodsInfo.ID,
			Num: -v.Num,
		})
	}
	_, err = l.svcCtx.GoodsRpc.UpdateStocks(l.ctx, &goods.UpdateStocksRequest{Infos: req})
	if err != nil {
		return err
	}
	return nil
}

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateRequest) (resp *types.OrderCreateResponse, err error) {
	err = l.updateStock(req)
	if err != nil {
		return nil, err
	}
	amount, err := l.getGoodsAmount(req)
	if err != nil {
		return nil, err
	}
	orderNo := uuid.NewString()
	err = l.createOrder(orderNo, amount.String(), req)
	if err != nil {
		return nil, err
	}
	url, err := pkg.GetWebPayUrl("商品", orderNo, amount.String())
	if err != nil {
		return nil, err
	}
	return &types.OrderCreateResponse{Url: url}, nil
}
