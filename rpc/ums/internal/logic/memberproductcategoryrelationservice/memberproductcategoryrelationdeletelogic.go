package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCategoryRelationDeleteLogic 商品与分类的关联
/*
Author: LiuFeiHua
Date: 2024/5/7 9:32
*/
type MemberProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationDeleteLogic {
	return &MemberProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCategoryRelationDelete 删除商品与分类的关联
func (l *MemberProductCategoryRelationDeleteLogic) MemberProductCategoryRelationDelete(in *umsclient.MemberProductCategoryRelationDeleteReq) (*umsclient.MemberProductCategoryRelationDeleteResp, error) {
	q := query.UmsMemberProductCategoryRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationDeleteResp{}, nil
}
