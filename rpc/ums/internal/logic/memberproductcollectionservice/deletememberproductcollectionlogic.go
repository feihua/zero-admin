package memberproductcollectionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberProductCollectionLogic 删除商品收藏/清空当前用户商品收藏列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:08
*/
type DeleteMemberProductCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberProductCollectionLogic {
	return &DeleteMemberProductCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberProductCollection 删除商品收藏/清空当前用户商品收藏列表
func (l *DeleteMemberProductCollectionLogic) DeleteMemberProductCollection(in *umsclient.DeleteMemberProductCollectionReq) (*umsclient.DeleteMemberProductCollectionResp, error) {
	q := query.UmsMemberProductCollection
	collectionDo := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		collectionDo = collectionDo.Where(q.ID.In(in.Ids...))
	}
	_, err := collectionDo.Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberProductCollectionResp{}, nil
}
