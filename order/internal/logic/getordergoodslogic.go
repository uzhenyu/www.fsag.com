package logic

import (
	"context"
	"zg6/z326/order/service"

	"zg6/z326/order/internal/svc"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderGoodsLogic {
	return &GetOrderGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderGoodsLogic) GetOrderGoods(in *order.GetOrderGoodsRequest) (*order.GetOrderGoodsResponse, error) {
	goods, err := service.GetOrderGoods(in.OrderID)
	if err != nil {
		return nil, err
	}
	return &order.GetOrderGoodsResponse{Infos: goods}, nil
}
