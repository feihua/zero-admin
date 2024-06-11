package memberbrandattentionservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberBrandAttentionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberBrandAttentionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberBrandAttentionDetailLogic {
	return &QueryMemberBrandAttentionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询会员关注品牌详情
func (l *QueryMemberBrandAttentionDetailLogic) QueryMemberBrandAttentionDetail(in *umsclient.QueryMemberBrandAttentionDetailReq) (*umsclient.QueryMemberBrandAttentionDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberBrandAttentionDetailResp{}, nil
}
