package logic

import (
	"context"
	"zg6/z326/goods/service"

	"zg6/z326/goods/goods"
	"zg6/z326/goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsIDsLogic {
	return &GetGoodsIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsIDsLogic) GetGoodsIDs(in *goods.GetGoodsIDsRequest) (*goods.GetGoodsIDsResponse, error) {
	gets, err := service.Gets(in.IDs)
	if err != nil {
		return nil, err
	}
	return &goods.GetGoodsIDsResponse{Infos: gets}, nil
}
