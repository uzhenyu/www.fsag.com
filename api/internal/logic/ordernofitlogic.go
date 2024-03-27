package logic

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"net/url"
	"zg6/z326/api/pkg"
	"zg6/z326/order/order"

	"github.com/zeromicro/go-zero/core/logx"
	"zg6/z326/api/internal/svc"
)

type OrderNofitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderNofitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderNofitLogic {
	return &OrderNofitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderNofitLogic) OrderNofit(req url.Values) error {
	fmt.Println("***********************收到回调")
	sign, err := pkg.VerifySign(req)
	if err != nil {
		logs.Info(err)
		return err
	}
	fmt.Println("*************回调信息**********")
	fmt.Println(sign)
	fmt.Println(err)
	_, err = l.svcCtx.OrderRpc.UpdateOrderByOrderNO(l.ctx, &order.UpdateOrderByOrderNORequest{
		OrderNo: sign.OutTradeNo,
		State:   2,
	})
	if err != nil {
		return err
	}
	return nil
}
