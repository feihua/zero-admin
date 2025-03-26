package memberlevelservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberLevelDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLevelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelDetailLogic {
	return &QueryMemberLevelDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询会员等级详情
func (l *QueryMemberLevelDetailLogic) QueryMemberLevelDetail(in *umsclient.QueryMemberLevelDetailReq) (*umsclient.QueryMemberLevelDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberLevelDetailResp{}, nil
}
