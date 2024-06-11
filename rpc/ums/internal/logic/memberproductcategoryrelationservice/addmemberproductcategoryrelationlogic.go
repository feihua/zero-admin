package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberProductCategoryRelationLogic 添加会员与产品分类关系表（用户喜欢的分类）
/*
Author: LiuFeiHua
Date: 2024/6/11 14:09
*/
type AddMemberProductCategoryRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberProductCategoryRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberProductCategoryRelationLogic {
	return &AddMemberProductCategoryRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberProductCategoryRelation 添加会员与产品分类关系表（用户喜欢的分类）
func (l *AddMemberProductCategoryRelationLogic) AddMemberProductCategoryRelation(in *umsclient.AddMemberProductCategoryRelationReq) (*umsclient.AddMemberProductCategoryRelationResp, error) {
	err := query.UmsMemberProductCategoryRelation.WithContext(l.ctx).Create(&model.UmsMemberProductCategoryRelation{
		MemberID:          in.MemberId,
		ProductCategoryID: in.ProductCategoryId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberProductCategoryRelationResp{}, nil
}
