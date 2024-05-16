package subjectproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SubjectProductRelationListLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 10:53
*/
type SubjectProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationListLogic {
	return &SubjectProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SubjectProductRelationList 查询关联专题
func (l *SubjectProductRelationListLogic) SubjectProductRelationList(in *cmsclient.SubjectProductRelationListReq) (*cmsclient.SubjectProductRelationListResp, error) {
	var ids []int64
	q := query.CmsSubjectProductRelation
	err := q.WithContext(l.ctx).Select(q.SubjectID).Where(q.ProductID.Eq(in.ProductId)).Scan(&ids)

	if err != nil {
		logc.Errorf(l.ctx, "查询关联专题列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &cmsclient.SubjectProductRelationListResp{
		SubjectIds: ids,
	}, nil

}
