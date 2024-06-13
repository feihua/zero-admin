package homeadvertiseservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHomeAdvertiseDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeAdvertiseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeAdvertiseDetailLogic {
	return &QueryHomeAdvertiseDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询首页轮播广告表详情
func (l *QueryHomeAdvertiseDetailLogic) QueryHomeAdvertiseDetail(in *smsclient.QueryHomeAdvertiseDetailReq) (*smsclient.QueryHomeAdvertiseDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryHomeAdvertiseDetailResp{}, nil
}
