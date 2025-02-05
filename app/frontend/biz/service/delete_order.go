package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpcorder "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteOrderService(Context context.Context, RequestContext *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteOrderService) Run(OrderId string, req *order.DeleteOrderReq) (resp *common.Empty, err error) {
	// 获取当前用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	rpcReq := &rpcorder.DeleteOrderReq{
		OrderId: OrderId,
		UserId:  userId,
	}
	//打印
	_, err = rpc.OrderClient.DeleteOrder(h.Context, rpcReq)
	return
}
