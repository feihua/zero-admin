package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCategoryRelationUpdateLogic 商品与分类的关联
/*
Author: LiuFeiHua
Date: 2024/5/7 10:15
*/
type MemberProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationUpdateLogic {
	return &MemberProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCategoryRelationUpdate 更新商品与分类的关联
func (l *MemberProductCategoryRelationUpdateLogic) MemberProductCategoryRelationUpdate(in *umsclient.MemberProductCategoryRelationUpdateReq) (*umsclient.MemberProductCategoryRelationUpdateResp, error) {
	_, err := query.UmsMemberProductCategoryRelation.WithContext(l.ctx).Updates(&model.UmsMemberProductCategoryRelation{
		ID:                in.Id,
		MemberID:          in.MemberId,
		ProductCategoryID: in.ProductCategoryId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationUpdateResp{}, nil
}
