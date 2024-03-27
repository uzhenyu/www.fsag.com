package logic

import (
	"context"
	"zg6/z326/order/service"

	"zg6/z326/order/internal/svc"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderByOrderNOLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderByOrderNOLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderByOrderNOLogic {
	return &UpdateOrderByOrderNOLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderByOrderNOLogic) UpdateOrderByOrderNO(in *order.UpdateOrderByOrderNORequest) (*order.UpdateOrderByOrderNOResponse, error) {
	err := service.UpdateByOrderNO(in.OrderNo, in.State)
	if err != nil {
		return nil, err
	}
	return &order.UpdateOrderByOrderNOResponse{}, nil
}
