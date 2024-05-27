package preferredareaproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// PreferredAreaProductRelationAddLogic 优选商品关联
/*
Author: LiuFeiHua
Date: 2024/5/11 10:59
*/
type PreferredAreaProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaProductRelationAddLogic {
	return &PreferredAreaProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PreferredAreaProductRelationAdd 添加优选商品关联
func (l *PreferredAreaProductRelationAddLogic) PreferredAreaProductRelationAdd(in *cmsclient.PreferredAreaProductRelationAddReq) (*cmsclient.PreferredAreaProductRelationAddResp, error) {
	//1.先删除优选商品的关联
	q := query.CmsPreferredAreaProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		return nil, err
	}

	//2.重新添加优选商品的关联
	productRelations := make([]*model.CmsPreferredAreaProductRelation, 0)
	for _, id := range in.PreferredAreaId {
		productRelations = append(productRelations, &model.CmsPreferredAreaProductRelation{
			PreferredAreaID: id,
			ProductID:       in.ProductId,
		})
	}

	err = q.WithContext(l.ctx).CreateInBatches(productRelations, len(productRelations))
	if err != nil {
		return nil, err
	}

	return &cmsclient.PreferredAreaProductRelationAddResp{}, nil
}
