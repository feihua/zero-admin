package memberproductcategoryrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberProductCategoryRelationLogic 添加会员与产品分类关系（用户喜欢的分类）
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

// AddMemberProductCategoryRelation 添加会员与产品分类关系（用户喜欢的分类）
func (l *AddMemberProductCategoryRelationLogic) AddMemberProductCategoryRelation(in *umsclient.AddMemberProductCategoryRelationReq) (*umsclient.AddMemberProductCategoryRelationResp, error) {
	err := query.UmsMemberProductCategoryRelation.WithContext(l.ctx).Create(&model.UmsMemberProductCategoryRelation{
		MemberID:          in.MemberId,          // 会员id
		ProductCategoryID: in.ProductCategoryId, // 商品分类id
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员与产品分类关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加会员与产品分类关系失败")
	}

	return &umsclient.AddMemberProductCategoryRelationResp{}, nil
}
