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

// MemberProductCollectionDelete 删除会员收藏的商品
func (l *MemberProductCollectionDeleteLogic) MemberProductCollectionDelete(in *umsclient.MemberProductCollectionDeleteReq) (*umsclient.MemberProductCollectionDeleteResp, error) {
	q := query.UmsMemberProductCollection
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCollectionDeleteResp{}, nil
}
