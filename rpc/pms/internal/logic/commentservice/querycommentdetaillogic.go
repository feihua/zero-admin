package commentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentDetailLogic {
	return &QueryCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询商品评价表详情
func (l *QueryCommentDetailLogic) QueryCommentDetail(in *pmsclient.QueryCommentDetailReq) (*pmsclient.QueryCommentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryCommentDetailResp{}, nil
}
