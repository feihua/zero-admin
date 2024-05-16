package memberproductcollectionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCollectionDeleteLogic 会员收藏的商品
/*
Author: LiuFeiHua
Date: 2023/11/30 12:02
*/
type MemberProductCollectionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionDeleteLogic {
	return &MemberProductCollectionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCollectionDelete 删除商品收藏/清空当前用户商品收藏列表
func (l *MemberProductCollectionDeleteLogic) MemberProductCollectionDelete(in *umsclient.MemberProductCollectionDeleteReq) (*umsclient.MemberProductCollectionDeleteResp, error) {
	q := query.UmsMemberProductCollection
	collectionDo := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		collectionDo = collectionDo.Where(q.ID.In(in.Ids...))
	}
	_, err := collectionDo.Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCollectionDeleteResp{}, nil
}
