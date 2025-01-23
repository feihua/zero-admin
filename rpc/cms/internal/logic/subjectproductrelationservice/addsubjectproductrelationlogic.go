package subjectproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSubjectProductRelationLogic 添加专题商品关系表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:41
*/
type AddSubjectProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectProductRelationLogic {
	return &AddSubjectProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSubjectProductRelation 添加专题商品关系表
func (l *AddSubjectProductRelationLogic) AddSubjectProductRelation(in *cmsclient.AddSubjectProductRelationReq) (*cmsclient.AddSubjectProductRelationResp, error) {
	q := query.CmsSubjectProductRelation
	// 1.先删除专题关联
	_, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		return nil, err
	}
	// 2.添加
	productRelations := make([]*model.CmsSubjectProductRelation, 0)
	for _, id := range in.SubjectId {
		productRelations = append(productRelations, &model.CmsSubjectProductRelation{
			SubjectID: id,
			ProductID: in.ProductId,
		})
	}

	err = q.WithContext(l.ctx).CreateInBatches(productRelations, len(productRelations))
	if err != nil {
		return nil, err
	}

	return &cmsclient.AddSubjectProductRelationResp{}, nil
}
