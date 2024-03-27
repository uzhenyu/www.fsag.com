package logic

import (
	"context"
	"zg6/z326/order/service"

	"zg6/z326/order/internal/svc"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	getOrder, err := service.GetOrder(in.ID)
	if err != nil {
		return nil, err
	}
	return &order.GetOrderResponse{Info: getOrder}, nil
}
