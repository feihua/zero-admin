package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCategoryRelationAddLogic 商品与分类的关联
/*
Author: LiuFeiHua
Date: 2024/5/7 10:14
*/
type MemberProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationAddLogic {
	return &MemberProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCategoryRelationAdd 添加商品与分类的关联
func (l *MemberProductCategoryRelationAddLogic) MemberProductCategoryRelationAdd(in *umsclient.MemberProductCategoryRelationAddReq) (*umsclient.MemberProductCategoryRelationAddResp, error) {
	err := query.UmsMemberProductCategoryRelation.WithContext(l.ctx).Create(&model.UmsMemberProductCategoryRelation{
		MemberID:          in.MemberId,
		ProductCategoryID: in.ProductCategoryId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationAddResp{}, nil
}
