package logic

import (
	"context"
	"zg6/z326/order/service"

	"zg6/z326/order/internal/svc"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	createOrder, err := service.CreateOrder(in.Info)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{Info: createOrder}, nil
}
