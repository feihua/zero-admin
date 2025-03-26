package subjectcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteSubjectCategoryLogic 删除专题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeleteSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectCategoryLogic {
	return &DeleteSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSubjectCategory 删除专题分类
func (l *DeleteSubjectCategoryLogic) DeleteSubjectCategory(in *cmsclient.DeleteSubjectCategoryReq) (*cmsclient.DeleteSubjectCategoryResp, error) {
	q := query.CmsSubjectCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除专题分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除专题分类失败")
	}

	return &cmsclient.DeleteSubjectCategoryResp{}, nil
}
