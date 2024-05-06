package subjectproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SubjectProductRelationAddLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 10:53
*/
type SubjectProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationAddLogic {
	return &SubjectProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SubjectProductRelationAdd 专题关联
func (l *SubjectProductRelationAddLogic) SubjectProductRelationAdd(in *cmsclient.SubjectProductRelationAddReq) (*cmsclient.SubjectProductRelationAddResp, error) {
	q := query.CmsSubjectProductRelation
	//1.先删除专题关联
	_, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		return nil, err
	}
	//2.添加
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

	return &cmsclient.SubjectProductRelationAddResp{}, nil
}
