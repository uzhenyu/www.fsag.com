package logic

import (
	"context"
	"zg6/z326/goods/service"

	"zg6/z326/goods/goods"
	"zg6/z326/goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStocksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStocksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStocksLogic {
	return &UpdateStocksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStocksLogic) UpdateStocks(in *goods.UpdateStocksRequest) (*goods.UpdateStocksResponse, error) {
	err := service.Update(in.Infos)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateStocksResponse{}, nil
}
