package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zg6/z326/api/internal/config"
	"zg6/z326/goods/goodsclient"
	"zg6/z326/order/orderclient"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goodsclient.Goods
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
