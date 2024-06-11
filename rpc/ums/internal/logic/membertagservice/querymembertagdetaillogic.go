package membertagservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberTagDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagDetailLogic {
	return &QueryMemberTagDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户标签表详情
func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(in *umsclient.QueryMemberTagDetailReq) (*umsclient.QueryMemberTagDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberTagDetailResp{}, nil
}
