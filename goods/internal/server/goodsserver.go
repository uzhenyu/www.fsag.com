// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package server

import (
	"context"

	"zg6/z326/goods/goods"
	"zg6/z326/goods/internal/logic"
	"zg6/z326/goods/internal/svc"
)

type GoodsServer struct {
	svcCtx *svc.ServiceContext
	goods.UnimplementedGoodsServer
}

func NewGoodsServer(svcCtx *svc.ServiceContext) *GoodsServer {
	return &GoodsServer{
		svcCtx: svcCtx,
	}
}

func (s *GoodsServer) GetGoodsIDs(ctx context.Context, in *goods.GetGoodsIDsRequest) (*goods.GetGoodsIDsResponse, error) {
	l := logic.NewGetGoodsIDsLogic(ctx, s.svcCtx)
	return l.GetGoodsIDs(in)
}

func (s *GoodsServer) UpdateStocks(ctx context.Context, in *goods.UpdateStocksRequest) (*goods.UpdateStocksResponse, error) {
	l := logic.NewUpdateStocksLogic(ctx, s.svcCtx)
	return l.UpdateStocks(in)
}
