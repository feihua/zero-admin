package subjectproductrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSubjectProductRelationLogic 添加专题商品关系
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

// AddSubjectProductRelation 添加专题商品关系
func (l *AddSubjectProductRelationLogic) AddSubjectProductRelation(in *cmsclient.AddSubjectProductRelationReq) (*cmsclient.AddSubjectProductRelationResp, error) {
	q := query.CmsSubjectProductRelation
	// 1.先删除专题关联
	_, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "先删除专题关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("先删除专题关联失败")
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
		logc.Errorf(l.ctx, "添加专题商品关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加专题商品关系失败")
	}

	return &cmsclient.AddSubjectProductRelationResp{}, nil
}
