package logic

import (
	"context"
	"zg6/z326/order/service"

	"zg6/z326/order/internal/svc"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderGoodsLogic {
	return &CreateOrderGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderGoodsLogic) CreateOrderGoods(in *order.CreateOrderGoodsRequest) (*order.CreateOrderGoodsResponse, error) {
	goods, err := service.CreateOrderGoods(in.Info)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderGoodsResponse{Info: goods}, nil
}
