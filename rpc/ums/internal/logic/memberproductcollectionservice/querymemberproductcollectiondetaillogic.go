package memberproductcollectionservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberProductCollectionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberProductCollectionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberProductCollectionDetailLogic {
	return &QueryMemberProductCollectionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户收藏的商品详情
func (l *QueryMemberProductCollectionDetailLogic) QueryMemberProductCollectionDetail(in *umsclient.QueryMemberProductCollectionDetailReq) (*umsclient.QueryMemberProductCollectionDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberProductCollectionDetailResp{}, nil
}
